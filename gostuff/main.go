package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"art.minty.backend/pMintyMultiToken"
	"art.minty.backend/pSale"

	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
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
)

func initViper() {
	viper.SetConfigFile("./config.json")
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

func main() {
	initViper()
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	flag.Parse()
	if *abi {
		fmt.Println("ABI")
		fmt.Println(pMintyMultiToken.PMintyMultiTokenABI)
		os.Exit(0)
	}
	client, err := ethclient.Dial(viper.GetString("HOST"))
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	banker := ethKeys.NewKey("adminKeys/banker")
	err = banker.RestoreOrCreate()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	bankerTx, err := bind.NewKeyedTransactorWithChainID(banker.GetKey(), chainID)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	if *action == "list" {
		chk_userName()

		records, err := getProjectsForUser(*userName)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
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
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
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
		if len(records[0].Address) == 0 {
			fmt.Println("Bad Address")
			return
		}
		addr := common.HexToAddress(records[0].Address)
		fmt.Println("Approving ", *search, " at ", addr.Hex())
		tx, err := pSaleC.SetMinter(bankerTx, addr, true)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex())
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		rct, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		fmt.Println("status", rct.Status)
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
		isha := strings.Split(*initialStakeHolders, ",")
		Lsha := strings.Split(*laterStakeHolders, ",")
		issa := strings.Split(*initialStakeShares, ",")
		Lssa := strings.Split(*laterStakeShares, ",")
		initialArray := makeSHA(isha, issa)
		laterArray := makeSHA(Lsha, Lssa)
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
		return
	}

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

*/
