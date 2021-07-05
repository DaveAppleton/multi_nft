package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"art.minty.backend/oldMultiToken"
	"art.minty.backend/pMintyMultiSale"
	"art.minty.backend/pMintyMultiToken"
	"art.minty.backend/pSale"
	"art.minty.backend/weth9"

	"github.com/DaveAppleton/etherUtils"
	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	database            *string = flag.String("database", "", "which DB to load")
	action              *string = flag.String("action", "", "what to do")                    // list
	search              *string = flag.String("search", "", "what to look for")              // LaylaHifi
	userName            *string = flag.String("user_name", "", "what to look for")           // LaylaHifi
	projectID           *int    = flag.Int("project_id", -1, "project to process")           // LaylaHifi
	royalty             *int    = flag.Int("royalty", -1, "royalty (0-150) where 150 = 15%") // LaylaHifi
	initialStakeHolders *string = flag.String("initial_share_addrs", "", "addresses of those who share first royalties")
	laterStakeHolders   *string = flag.String("later_share_addrs", "", "addresses of those who share later royalties")
	initialStakeShares  *string = flag.String("initial_share_shares", "", "shares (/1000) of those who share first royalties")
	laterStakeShares    *string = flag.String("later_share_shares", "", "shares (/1000) of those who share later royalties")
	override            *bool   = flag.Bool("override", false, "override empty field checks")
	abi                 *bool   = flag.Bool("abi", false, "print the abi of pMintyMultiToken")
	bankerTx            *bind.TransactOpts
)

func initViper() {
	viper.SetConfigFile("./config_" + *database + ".json")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("./config.json")
		log.Fatal("config.json", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("viper config changed", e.Name)
	})
}

// -action list -user_name LaylaHifi
// -action approve_user -user_name LaylaHifi
// -action approve_project -project_id 15

func chkErr(msg string, err error) {
	if err == nil {
		return
	}
	fmt.Println(msg, err)
	log.Fatal(msg, err)
}

func main() {
	flag.Parse()
	initViper()
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})

	if *abi {
		fmt.Println("ABI")
		fmt.Println(pMintyMultiToken.PMintyMultiTokenABI)
		os.Exit(0)
	}
	client, err := ethclient.Dial(viper.GetString("HOST"))
	chkErr("Get Client", err)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	chkErr("chain ID", err)
	var banker *ethKeys.AccountKey
	if *database == "live" {
		banker = ethKeys.NewKey("adminKeys/liveBanker")
	} else {
		banker = ethKeys.NewKey("adminKeys/banker")
	}
	err = banker.RestoreOrCreate()
	chkErr("banker", err)
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	bal, err := client.BalanceAt(ctx, banker.PublicKey(), nil)
	chkErr("Bank Balance", err)

	fmt.Println("Banker", banker.PublicKeyAsHexString(), " balance ", etherUtils.EtherToStr(bal))
	if bal.Cmp(etherUtils.PointOneEther()) < 0 {
		fmt.Println("please top up banker ")
		os.Exit(1)
	}

	bankerTx, err = bind.NewKeyedTransactorWithChainID(banker.GetKey(), chainID)
	chkErr("transactor", err)

	// go run . -action late_test -database live

	if *action == "late_test" {

		samouraiArt := ethKeys.NewKey("adminKeys/samourai")
		err = samouraiArt.RestoreOrCreate()
		chkErr("samourai", err)
		ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		bal, err := client.BalanceAt(ctx, samouraiArt.PublicKey(), nil)
		chkErr("samourai Balance", err)
		if *database != "live" {
			if bal.Cmp(etherUtils.PointOneEther()) < 0 {
				fmt.Println("please top up samourai ", samouraiArt.PublicKeyAsHexString())
				os.Exit(1)
			}
		}
		samTx, err := bind.NewKeyedTransactorWithChainID(samouraiArt.Key, chainID)
		chkErr("samourai txor", err)
		owner := samouraiArt.PublicKey()

		if *database == "live" {
			owner = common.HexToAddress("0xA55a74B9a0e9F3e091CF56aC7241eC58BB2a323e")
		}

		platform := common.HexToAddress("0xa653d4C79200f4983b5eE0E61F068bDE90F42Db7")
		patron := common.HexToAddress("0x43C41e26BBB45Be8b9dFAbe757707004ea62663c")
		sam := common.HexToAddress("0x0131fD54EFE26b2dF471182bA51CcA47dD01688e")
		pboy := common.HexToAddress("0x1af70e564847be46e4ba286c0b0066da8372f902")
		kol := common.HexToAddress("0x1f1b2406987376055491Be1EeEEECC212c19Bfa8")

		weth := common.HexToAddress("0xf738b83Fa52A7Ab570918Afe61b78b8E2DC6F4EF")
		if *database == "live" {
			weth = common.HexToAddress("0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619")
		}

		wethC, err := weth9.NewWETH9(weth, client)
		chkErr("WETH", err)

		var pContract *pMintyMultiSale.PMintyMultiSale
		var pAddress common.Address
		var tx *types.Transaction

		deploy := false
		if deploy {
			divisor := big.NewInt(1025)
			pAddress, tx, pContract, err = pMintyMultiSale.DeployPMintyMultiSale(bankerTx, client, platform, weth, divisor)
			chkErr("Deploy MultiSale", err)
			_, err = bind.WaitDeployed(ctx, client, tx)
			chkErr("WaitMined", err)
			fmt.Println("Sale deployed at ", pAddress.Hex(), "in tx", tx.Hash().Hex())
			log.Println("Sale deployed at ", pAddress.Hex(), "in tx", tx.Hash().Hex())
			if *database == "live" {
				os.Exit(0)
			}
		} else {
			pAddress = common.HexToAddress("0xfC871e50ca8870caB76cACa500711dc4060433Ab")
			if *database == "live" {
				pAddress = common.HexToAddress("0xdf4837218075c62e1B631BEE0a3c6e8df66cF8C8")
			}
			pContract, err = pMintyMultiSale.NewPMintyMultiSale(pAddress, client)
			chkErr("restore sale", err)
		}
		buyer := ethKeys.NewKey("adminKeys/buyer")
		buyer2 := ethKeys.NewKey("adminKeys/buyer2")

		minty := ethKeys.NewKey("userKeys/minty")

		chkErr("buyer", buyer.RestoreOrCreate())
		chkErr("buyer2", buyer2.RestoreOrCreate())
		chkErr("minty", minty.RestoreOrCreate())

		fmt.Println("buyer", buyer.PublicKeyAsHexString())
		fmt.Println("buyer2", buyer2.PublicKeyAsHexString())
		fmt.Println("minty", minty.PublicKeyAsHexString())

		ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		bal, err = client.BalanceAt(ctx, buyer.PublicKey(), nil)
		chkErr("Buyer Balance", err)
		if bal.Cmp(big.NewInt(0)) <= 0 && *database != "live" {
			fmt.Println("top up buyer 1 ", buyer.PublicKeyAsHexString())
			os.Exit(1)
		}

		ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		bal, err = client.BalanceAt(ctx, buyer2.PublicKey(), nil)
		chkErr("Buyer2 Balance", err)
		if bal.Cmp(big.NewInt(0)) <= 0 && *database != "live" {
			fmt.Println("top up buyer 2 ", buyer2.PublicKeyAsHexString())
			os.Exit(1)
		}
		buyerTx, err := bind.NewKeyedTransactorWithChainID(buyer.GetKey(), chainID)
		chkErr("transactor", err)
		buyer2Tx, err := bind.NewKeyedTransactorWithChainID(buyer2.GetKey(), chainID)
		chkErr("transactor", err)

		eth100 := new(big.Int).Mul(etherUtils.OneEther(), big.NewInt(100))
		appWETH := false
		if appWETH {
			fmt.Println("approving WETH")
			tx, err = wethC.Approve(buyerTx, pAddress, eth100)
			chkErr("b1 approve eth", err)
			ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			bind.WaitMined(ctx, client, tx)
			tx, err = wethC.Approve(buyer2Tx, pAddress, eth100)
			chkErr("b2 approve eth", err)
			ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			bind.WaitMined(ctx, client, tx)
		}
		typeA := []common.Address{patron, sam, pboy, kol}
		blue := []int{5, 8, 2, 0}
		yellow := []int{5, 3, 0, 7}
		pink := []int{5, 6, 2, 2}
		blueArr := makeShares(typeA, blue)
		yellowArr := makeShares(typeA, yellow)
		pinkArr := makeShares(typeA, pink)

		locking := common.HexToAddress(viper.GetString("LOCKING"))
		royaltyPerMille := big.NewInt(150)
		//pMSaleA := common.HexToAddress(viper.GetString("P_M_SALE"))
		_ = yellowArr
		_ = blueArr

		var pMMTA common.Address
		var pMMTC *pMintyMultiToken.PMintyMultiToken
		restore := false
		if restore {
			pMMTA = common.HexToAddress("0x16429506d9695E3FDeF312C2421931Dd3F84d0b4")
			pMMTC, err = pMintyMultiToken.NewPMintyMultiToken(pMMTA, client)
			chkErr("Restore Token", err)
			fmt.Println("Token restored")
			// tx, err := pMMTC.Mint(bankerTx, big.NewInt(420), big.NewInt(20), "Qmey17JmcCs7oBrQNwa165GpCDThBFJNpgpkALW8vyaeYg", big.NewInt(0))
			// chkErr("Mint", err)
			// fmt.Println("mint", tx.Hash().Hex())
			// ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
			// defer cancel()
			// _, err = bind.WaitMined(ctx, client, tx)
			// chkErr("Minted for MSIG", err)
			// var data []byte
			// multisig := common.HexToAddress("0xa653d4C79200f4983b5eE0E61F068bDE90F42Db7")
			// tx, err = pMMTC.SafeTransferFrom(bankerTx, bankerTx.From, multisig, big.NewInt(420), big.NewInt(10), data)
			// chkErr("xfer msig", err)
			// ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
			// defer cancel()
			// rct, err := bind.WaitMined(ctx, client, tx)
			// chkErr("wait mined", err)
			// fmt.Println("Send to MSIG returned ", rct.Status, tx.Hash().Hex())
			// os.Exit(1)
		} else {
			fmt.Println("deploying multi")
			pMMTA, tx, pMMTC, err = pMintyMultiToken.DeployPMintyMultiToken(
				bankerTx,
				client,
				owner,
				pAddress,
				[]common.Address{locking},
				"You need to be a MINTY patron",
				royaltyPerMille,
				blueArr,
				blueArr,
				"Art",
			)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
			fmt.Println("deploy token", pMMTA.Hex(), tx.Hash().Hex())
			log.Println("deploy token", pMMTA.Hex(), tx.Hash().Hex())
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()
			depA, err := bind.WaitDeployed(ctx, client, tx)
			chkErr("Deploy MultiToken", err)
			fmt.Println("deployed at", depA.Hex(), pMMTA.Hex())
		}
		da, err := pMMTC.Deployer(nil)
		chkErr("asking deployer", err)
		fmt.Println("Deployer address", da.Hex())

		tx, err = pMMTC.AddPools(bankerTx, yellowArr, yellowArr, "Music")
		chkErr("Add Yellow", err)
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		rct, err := bind.WaitMined(ctx, client, tx)
		chkErr("Mined Yellow", err)
		fmt.Println("yellow", tx.Hash().Hex(), rct.Status)
		tx, err = pMMTC.AddPools(bankerTx, pinkArr, pinkArr, "ArtAndMusic")
		chkErr("Add Pink", err)
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		rct, err = bind.WaitMined(ctx, client, tx)
		chkErr("Mined Pink", err)
		fmt.Println("pink", tx.Hash().Hex(), rct.Status)
		if *database == "live" {
			os.Exit(0)
		}
		tx, err = pMMTC.SetApprovalForAll(samTx, pAddress, true)
		chkErr("banker approve", err)
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		rct, err = bind.WaitMined(ctx, client, tx)
		chkErr("Mine  sam approve", err)
		fmt.Println("sam approve", tx.Hash().Hex(), rct.Status)

		tx, err = pMMTC.SetApprovalForAll(buyerTx, pAddress, true)
		chkErr("buyer approve", err)
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		rct, err = bind.WaitMined(ctx, client, tx)
		chkErr("Mined buyer approve", err)
		fmt.Println("buyer approve", tx.Hash().Hex(), rct.Status)

		tx, err = pMMTC.SetApprovalForAll(buyer2Tx, pAddress, true)
		chkErr("buyer2 approve", err)
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		rct, err = bind.WaitMined(ctx, client, tx)
		chkErr("Mined buyer2 approve", err)
		fmt.Println("buyer2 approve", tx.Hash().Hex(), rct.Status)
		pOwner, err := pMMTC.Owner(nil)
		chkErr("Get owner", err)
		fmt.Println("tokend owned by ", pOwner.Hex())

		pOne25, _ := etherUtils.StrToEther("0.1025")
		one25, _ := etherUtils.StrToEther("1.025")

		for jx := int64(1); jx < 4; jx++ {
			offer := true
			if offer {
				tx, err := pContract.OfferNew(samTx, pMMTA, big.NewInt(jx), fmt.Sprintf("HASH%d", jx), big.NewInt(5), pOne25, big.NewInt((jx - 1)))
				chkErr("Offer New", err)
				fmt.Println(jx, "offer new", tx.Hash())
				ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
				defer cancel()
				rct, err := bind.WaitMined(ctx, client, tx)
				chkErr("Wait mined", err)
				fmt.Println(jx, "Status", rct.Status)
			}
		}
		for jx := int64(1); jx < 4; jx++ {
			tx, err := pContract.AcceptOffer(buyer2Tx, pMMTA, big.NewInt(jx), big.NewInt(0), big.NewInt(5))
			chkErr("Accept offer", err)
			fmt.Println(jx, "accept offer", tx.Hash())
			ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()
			rct, err := bind.WaitMined(ctx, client, tx)
			chkErr("Wait mined", err)
			fmt.Println(jx, "Status", rct.Status)
		}
		for jx := int64(1); jx < 4; jx++ {
			tx, err := pContract.OfferResale(buyer2Tx, pMMTA, big.NewInt(jx), big.NewInt(5), one25)
			chkErr("Accept offer", err)
			fmt.Println(jx, "accept offer", tx.Hash())
			ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()
			rct, err := bind.WaitMined(ctx, client, tx)
			chkErr("Wait mined", err)
			fmt.Println(jx, "Status", rct.Status)
		}
		for jx := int64(1); jx < 4; jx++ {
			tx, err := pContract.AcceptOffer(buyerTx, pMMTA, big.NewInt(jx), big.NewInt(1), big.NewInt(5))
			chkErr("Accept offer", err)
			fmt.Println(jx, "accept offer", tx.Hash())
			ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()
			rct, err := bind.WaitMined(ctx, client, tx)
			chkErr("Wait mined", err)
			fmt.Println(jx, "Status", rct.Status)
		}
		fmt.Println("DONE")
		os.Exit(0)
	}

	if *action == "list" {
		chk_userName()

		records, err := getProjectsForUser(*userName)
		chkErr("ProjectsForUser", err)
		for _, rec := range records {
			pt := ""
			if rec.ProjectToken.Valid {
				pt = rec.ProjectToken.String
			}
			fmt.Println(
				rec.UserID,
				rec.Name,
				rec.Address,
				rec.ArtistID,
				rec.ArtistApproved,
				rec.ProjectID,
				rec.Title,
				rec.ProjectApproved,
				pt,
			)
		}
		return
	}
	if *action == "approve_user" {
		pSaleA := common.HexToAddress(viper.GetString("P_SALE"))
		pSaleC, err := pSale.NewPMintysale(pSaleA, client)
		chkErr("New Minty Sale", err)
		fmt.Println("finding ", *userName)
		_, artistAddr, err := findArtistAddress(*userName)
		chkErr("Find Artist", err)

		// records, err := getProjectsForUser(*userName)
		// chkErr("ProjectsForUser", err)
		// if len(records) == 0 {
		// 	fmt.Println("no records found")
		// 	return
		// }
		// if len(records[0].Address) == 0 {
		// 	fmt.Println("Bad Address")
		// 	return
		// }
		addr := common.HexToAddress(artistAddr)
		fmt.Println("Approving ", *userName, " at ", addr.Hex())
		doit := true
		if doit {
			tx, err := pSaleC.SetMinter(bankerTx, addr, true)
			chkErr("SetMinter", err)
			fmt.Println(tx.Hash().Hex())
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()
			rct, err := bind.WaitMined(ctx, client, tx)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
			fmt.Println("status", rct.Status)
		}
		return
	}
	if *action == "approve_project" {
		fmt.Println("Aprove Project")
		if *royalty < 0 {
			fmt.Println("royalty not set")
			flag.Usage()
			return
		}
		if *royalty > 150 {
			fmt.Println("royalty > 150 / 1000 (15%)")
			flag.Usage()
			return
		}
		records, err := getProjectsForUser(*userName)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		if len(records) == 0 {
			fmt.Println("no records found")
			return
		}
		fmt.Println(len(records), " projects found")
		if *database == "demo" {
			isha := strings.Split(*initialStakeHolders, ",")
			Lsha := strings.Split(*laterStakeHolders, ",")
			issa := strings.Split(*initialStakeShares, ",")
			Lssa := strings.Split(*laterStakeShares, ",")
			initialArray := makeSHA3(isha, issa)
			laterArray := makeSHA3(Lsha, Lssa)

			done := false
			for _, rec := range records {
				if rec.ProjectID == *projectID {
					fmt.Println(rec.Title)
					if rec.ProjectToken.Valid && !*override {
						msg := fmt.Sprintln("Project ",
							rec.Title,
							"already has a token ",
							rec.ProjectToken.String)
						fmt.Println(msg)
						log.Println(msg)
						return
					}
					fmt.Println("Approving ", rec.Title, "by", rec.Name)
					owner := common.HexToAddress(rec.Address)
					pMSaleA := common.HexToAddress(viper.GetString("P_M_SALE"))
					locking := common.HexToAddress(viper.GetString("LOCKING"))
					royaltyPerMille := big.NewInt(int64(*royalty))

					pMMTA, tx, pMMTC, err := pMintyMultiToken.DeployPMintyMultiToken(
						bankerTx,
						client,
						owner,
						pMSaleA,
						[]common.Address{locking},
						"You need to be a MINTY patron",
						royaltyPerMille,
						initialArray,
						laterArray,
						"initialPool",
					)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
					fmt.Println(pMMTA.Hex(), tx.Hash().Hex())
					ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
					defer cancel()
					depA, err := bind.WaitDeployed(ctx, client, tx)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
					fmt.Println("deployed at ", depA.Hex())
					tx, err = pMMTC.SetContractURI(bankerTx, "https://minty.mypinata.cloud/ipfs/"+rec.Hash)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
					ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
					defer cancel()
					rct, err := bind.WaitMined(ctx, client, tx)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
					if rct.Status == 0 {
						fmt.Println("Setting Contract URI failed for", pMMTA.Hex())
						log.Println("Setting Contract URI failed for", pMMTA.Hex())
					}
					err = updateProjectWithContract(*projectID, pMMTA)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
					fmt.Println("database updated")
					done = true
				}
				if done {
					break
				}
				fmt.Println("Not ", rec.Title, "(", rec.ProjectID, ")")
			}
			if !done {
				fmt.Println()
				log.Println()
			}
		} else if *database == "staging" {
			isha := strings.Split(*initialStakeHolders, ",")
			Lsha := strings.Split(*laterStakeHolders, ",")
			issa := strings.Split(*initialStakeShares, ",")
			Lssa := strings.Split(*laterStakeShares, ",")
			initialArray := makeSHA3(isha, issa)
			laterArray := makeSHA3(Lsha, Lssa)
			records, err := getProjectsForUser(*userName)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
			done := false
			for _, rec := range records {
				if rec.ProjectID == *projectID {
					fmt.Println(rec.Title)
					if rec.ProjectToken.Valid && !*override {
						msg := fmt.Sprintln("Project ",
							rec.Title,
							"already has a token ",
							rec.ProjectToken.String)
						fmt.Println(msg)
						log.Println(msg)
						return
					}
					fmt.Println("Approving ", rec.Title, "by", rec.Name)
					owner := common.HexToAddress(rec.Address)
					pMSaleA := common.HexToAddress(viper.GetString("P_M_SALE"))
					locking := common.HexToAddress(viper.GetString("LOCKING"))
					royaltyPerMille := big.NewInt(int64(*royalty))

					pMMTA, tx, _, err := pMintyMultiToken.DeployPMintyMultiToken(
						bankerTx,
						client,
						owner,
						pMSaleA,
						[]common.Address{locking},
						"You need to be a MINTY patron",
						royaltyPerMille,
						initialArray,
						laterArray,
						"initialPool",
					)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
					fmt.Println(pMMTA.Hex(), tx.Hash().Hex())
					ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
					defer cancel()
					depA, err := bind.WaitDeployed(ctx, client, tx)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
					fmt.Println("deployed at ", depA.Hex())
					err = updateProjectWithContract(*projectID, pMMTA)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
					fmt.Println("database updated")
					done = true
				}
				if done {
					break
				}
				fmt.Println("Not ", rec.Title, "(", rec.ProjectID, ")")
			}
			if !done {
				fmt.Println()
				log.Println()
			}
		}
		return
	}
	if *action == "check_approved" {
		artists, max, err := getArtists(0)
		chkErr("getArtists", err)
		fmt.Println(max, "artists found")
		pSaleA := common.HexToAddress(viper.GetString("P_SALE"))
		pSaleC, err := pSale.NewPMintysale(pSaleA, client)
		chkErr("New Minty Sale", err)

		for _, artist := range artists {
			if !artist.Nickname.Valid {
				continue
			}
			addr := common.HexToAddress(artist.Address)
			minter, err := pSaleC.IsMinter(nil, addr)
			chkErr("isMinter", err)
			if minter {
				fmt.Println(artist.Nickname.String, "is a minter")
				continue
			}
			fmt.Println("Approving ", artist.Nickname, " at ", addr.Hex())
			tx, err := pSaleC.SetMinter(bankerTx, addr, true)
			chkErr("SetMinter", err)
			fmt.Println(tx.Hash().Hex())
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()
			rct, err := bind.WaitMined(ctx, client, tx)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
			fmt.Println("status", rct.Status)
		}
	}
	if *action == "web" {
		http.HandleFunc("/artists", showArtists)
		http.HandleFunc("/approve_artist", approveArtist)
		http.HandleFunc("/project", project)
		http.HandleFunc("/projectArts", projectArt)
		http.HandleFunc("/setPoolID", setPool)
		err := http.ListenAndServe(":8090", nil)
		chkErr("Web Server", err)

	}
	// if *action == "set_metadata" {
	// 	hash, err := getContractMetaData(*projectID)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		log.Fatal(err)
	// 	}
	// }

}

func makeSHA(strAddrA []string, strShareA []string) (pe []pMintyMultiToken.PoolEntry) {
	if len(strAddrA) != len(strShareA) {
		fmt.Println(len(strAddrA), "addresses but ", len(strShareA), "shares")
		log.Fatal(len(strAddrA), "addresses but ", len(strShareA), "shares")
	}
	for pos, addrStr := range strAddrA {
		addr := common.HexToAddress(addrStr)
		share, err := strconv.ParseInt(strShareA[pos], 10, 64)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		pe = append(pe, pMintyMultiToken.PoolEntry{addr, big.NewInt(share)})
	}
	return
}

func makeSHA2(strAddrA []string, strShareA []string) (pe []oldMultiToken.PoolEntry) {
	if len(strAddrA) != len(strShareA) {
		fmt.Println(len(strAddrA), "addresses but ", len(strShareA), "shares")
		log.Fatal(len(strAddrA), "addresses but ", len(strShareA), "shares")
	}
	for pos, addrStr := range strAddrA {
		addr := common.HexToAddress(addrStr)
		share, err := strconv.ParseInt(strShareA[pos], 10, 64)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		pe = append(pe, oldMultiToken.PoolEntry{addr, big.NewInt(share)})
	}
	return
}

func makeSHA3(strAddrA []string, strShareA []string) (pe []pMintyMultiToken.PoolEntry) {
	if len(strAddrA) != len(strShareA) {
		fmt.Println(len(strAddrA), "addresses but ", len(strShareA), "shares")
		log.Fatal(len(strAddrA), "addresses but ", len(strShareA), "shares")
	}
	for pos, addrStr := range strAddrA {
		addr := common.HexToAddress(addrStr)
		share, err := strconv.ParseInt(strShareA[pos], 10, 64)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		pe = append(pe, pMintyMultiToken.PoolEntry{addr, big.NewInt(share)})
	}
	return
}

func makeShares(accounts []common.Address, shares []int) (pa []pMintyMultiToken.PoolEntry) {
	var tot int64
	for pos, share := range shares {
		vShare := int64(share * 1000 / 15)
		tot += vShare
		if vShare != 0 {
			pe := pMintyMultiToken.PoolEntry{accounts[pos], big.NewInt(vShare)}
			pa = append(pa, pe)
		}
	}
	tot = 1000 - tot
	for tot > 0 {
		pa[tot].Share = new(big.Int).Add(pa[tot].Share, big.NewInt(1))
		tot--
	}
	return
}

func chk_userName() {
	if len(*userName) == 0 {
		fmt.Println("-action list -user_name xxx")
		os.Exit(1)
	}
}

// go run . -action list -user_name LaylaHifi
// go run . -action approve_project -user_name LaylaHifi -royalty 100 -initial_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -later_share_addrs 0xcad809e60e7B93fF7343E25F891A4959F0a2BBC7,0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -initial_share_shares 1000 -later_share_shares 500,500 -project_id 75

/*
go run . -action approve_project -user_name LaylaHifi -royalty 100 -initial_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -later_share_addrs 0xcad809e60e7B93fF7343E25F891A4959F0a2BBC7,0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -initial_share_shares 1000 -later_share_shares 500,500  -project_id 162 -user_name LaylaHifi -override yes
Aprove Project
5  projects found
Not  test ( 143 )
Not  Private project 2 ( 149 )
Not  Private project 2 ( 150 )
Not  test test test ( 170 )
LIFE MUST GO ON
Approving  LIFE MUST GO ON by LaylaHifi
0x4663180 0xc43b65fc79de24633886def5a8974fe3c1d970a3d146e0d07e82c7ad23b4bf65

Actually produced : 162


go run . -action approve_project -user_name Neuman -royalty 100 -initial_share_addrs 0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -later_share_addrs 0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3,0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -initial_share_shares 1000 -later_share_shares 500,500  -project_id 160 -override yes
0x82dC0832693619bA4eDE5e647f591ec9635040fd

go run . -action approve_project -user_name Virgo -royalty 100 -initial_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -later_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 145 -override yes
0x86b581452B7768A7df8fe27091C248d207126a93

go run . -action approve_project -user_name seri -royalty 100 -initial_share_addrs 0xa454515041892eB78132293ABd5763a730412F65 -later_share_addrs 0xa454515041892eB78132293ABd5763a730412F65,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 148 -override yes
0x67851c3A36A46e9836Dbad84e5585166517798DE

go run . -action approve_project -user_name NFTYDaddy -royalty 100 -initial_share_addrs 0xa1EFF0887B93e18bB0f59334a0CB57148BC2086f -later_share_addrs 0xa1EFF0887B93e18bB0f59334a0CB57148BC2086f,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 161

go run . -action approve_project -user_name Jahanzaib -royalty 100 -initial_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1 -later_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1q,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 174

go run . -action approve_project -user_name amirfive -royalty 100 -initial_share_addrs 0xc355b9A5dE199F811542D8C5D3E309789FfBEf86 -later_share_addrs 0xc355b9A5dE199F811542D8C5D3E309789FfBEf86,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 165

go run . -action approve_project -user_name Amir -royalty 100 -initial_share_addrs 0xFE029208b267daBF5077bB3E3E7B1cc9916e9943 -later_share_addrs 0xFE029208b267daBF5077bB3E3E7B1cc9916e9943,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 163

go run . -action approve_project -user_name Jahanzaib -royalty 100 -initial_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1 -later_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 196

Abasa : 259,0xD27b5dB83D5cCF7F4e796abAa75D6C67ea9f8e7F | artist : 96 | art 463 |

go run . -action approve_project -user_name Aatqa -royalty 100 -initial_share_addrs 0xB0f1A004b34457F62a29f42dF1e6D82bf9e39c74 -later_share_addrs 0xB0f1A004b34457F62a29f42dF1e6D82bf9e39c74,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 198
0xB0f1A004b34457F62a29f42dF1e6D82bf9e39c74 Aatqa

0xb0f1a004b34457f62a29f42df1e6d82bf9e39c74

go run . -action approve_project -user_name AbdullahIsa -royalty 100 -initial_share_addrs 0x1E989cD65Fec9961EC26A3492B940045244e4893 -later_share_addrs 0x1E989cD65Fec9961EC26A3492B940045244e4893,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 200
200

go run . -action approve_project -database demo -user_name seri -royalty 100 -initial_share_addrs 0xa454515041892eB78132293ABd5763a730412F65 -later_share_addrs 0xa454515041892eB78132293ABd5763a730412F65,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 10

go run . -action approve_project -database demo -user_name Hamirul -royalty 100 -initial_share_addrs 0xcad809e60e7B93fF7343E25F891A4959F0a2BBC7 -later_share_addrs 0xcad809e60e7B93fF7343E25F891A4959F0a2BBC7,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 11

go run . -action approve_project -database demo -user_name AmirHussain -royalty 100 -initial_share_addrs 0xFE029208b267daBF5077bB3E3E7B1cc9916e9943 -later_share_addrs 0xFE029208b267daBF5077bB3E3E7B1cc9916e9943,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 12

go run . -action approve_project -database demo -user_name Virgo -royalty 100 -initial_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -later_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 6

go run . -action approve_project  -database demo -user_name Neumann -royalty 100 -initial_share_addrs 0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -later_share_addrs 0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3,0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -initial_share_shares 1000 -later_share_shares 500,500  -project_id 14

go run . -action approve_project  -database staging -user_name Neuman -royalty 100 -initial_share_addrs 0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -later_share_addrs 0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3,0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -initial_share_shares 1000 -later_share_shares 500,500  -project_id 186

go run . -action approve_project  -database demo -user_name JahanzaibRa -royalty 100 -initial_share_addrs 0x77f38012590497228Cf9b82E6174ce7539F0945B -later_share_addrs 0x77f38012590497228Cf9b82E6174ce7539F0945B,0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -initial_share_shares 1000 -later_share_shares 500,500  -project_id 17

go run . -action approve_project  -database demo -user_name Neumann -royalty 100 -initial_share_addrs 0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -later_share_addrs 0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3,0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -initial_share_shares 1000 -later_share_shares 500,500  -project_id 15

go run . -action approve_project -database demo -user_name Virgo -royalty 100 -initial_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -later_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 22

go run . -action approve_project -database demo -user_name Libra -royalty 100 -initial_share_addrs 0x08ef5C04D2F93d9858F0d87697f22cBeCD990076 -later_share_addrs 0x08ef5C04D2F93d9858F0d87697f22cBeCD990076,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 23

go run . -action approve_project -database demo -user_name Libra -royalty 100 -initial_share_addrs 0x08ef5C04D2F93d9858F0d87697f22cBeCD990076 -later_share_addrs 0x08ef5C04D2F93d9858F0d87697f22cBeCD990076,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 23

go run . -action approve_project -database demo -user_name AbdullahIsa -royalty 100 -initial_share_addrs 0x1E989cD65Fec9961EC26A3492B940045244e4893 -later_share_addrs 0x1E989cD65Fec9961EC26A3492B940045244e4893,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 29

go run . -action approve_project -database demo -user_name MUD -royalty 100 -initial_share_addrs 0x0EA0281db722F5642457B5E785dda24eA75f6149 -later_share_addrs 0x0EA0281db722F5642457B5E785dda24eA75f6149,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 30

go run . -action approve_project -database demo -user_name Jahanzaib -royalty 100 -initial_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1 -later_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 16
go run . -action approve_project -database demo -user_name Jahanzaib -royalty 100 -initial_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1 -later_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 19
go run . -action approve_project -database demo -user_name Jahanzaib -royalty 100 -initial_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1 -later_share_addrs 0xA7722c27a0eAaa4690358eE3D31a4D38F20c38A1,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 20

go run . -action approve_project -database demo -user_name sinister -royalty 100 -initial_share_addrs 0xa454515041892eB78132293ABd5763a730412F65 -later_share_addrs 0xa454515041892eB78132293ABd5763a730412F65,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 34

go run . -action approve_project -database demo -user_name Virgo -royalty 100 -initial_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -later_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 39

go run . -action approve_project -database demo -user_name Cat -royalty 100 -initial_share_addrs 0xd08132BC96bBe55E33a61E5a55cA0a237F515bf0 -later_share_addrs 0xd08132BC96bBe55E33a61E5a55cA0a237F515bf0,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 41
go run . -action approve_project -database demo -user_name LaylaHifi -royalty 100 -initial_share_addrs 0xcad809e60e7B93fF7343E25F891A4959F0a2BBC7 -later_share_addrs 0xcad809e60e7B93fF7343E25F891A4959F0a2BBC7,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 37

go run . -action approve_project -database demo -user_name Libra -royalty 100 -initial_share_addrs 0x08ef5C04D2F93d9858F0d87697f22cBeCD990076 -later_share_addrs 0x08ef5C04D2F93d9858F0d87697f22cBeCD990076,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 42

go run . -action approve_project -database demo -user_name Sinister -royalty 100 -initial_share_addrs 0xa454515041892eB78132293ABd5763a730412F65 -later_share_addrs 0xa454515041892eB78132293ABd5763a730412F65,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 43

select nickname from users a, artist b, project c where a.id=b.user_id and b.id = c.artist_id and c.id=19;

go run . -action approve_project -database demo -user_name Virgo -royalty 100 -initial_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d -later_share_addrs 0x39d07f321cAF5b0668459DB5Bcf039A462A9273d,0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3 -initial_share_shares 1000 -later_share_shares 500,500  -project_id 44


Note : need warning that artist should have nickname set
*/
