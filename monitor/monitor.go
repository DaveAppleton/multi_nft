package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"sync"

	"art.minty.monitor/pMintyMultiSale"
	"art.minty.monitor/pMintyMultiToken"
	"art.minty.monitor/pMintyUnique"
	"art.minty.monitor/pSale"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	client   *ethclient.Client
	database *string = flag.String("database", "", "which DB to load")
)

func initViper() {
	fmt.Println("using ./config_" + *database + ".json")
	viper.SetConfigFile("./config_" + *database + ".json")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("./config.json")
		log.Fatal("config.json", err)
	}
	viper.WatchConfig()

}

func chkErrX(msg string, err error, term bool) {
	if err == nil {
		return
	}
	fmt.Println(msg, err)
	if term {
		log.Fatal(msg, err)
	}
	log.Println(msg, err)
}

func chkErr(msg string, err error) {
	chkErrX(msg, err, true)
}

func main() {
	var err error
	var wg sync.WaitGroup
	flag.Parse()
	initViper()
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	host := viper.GetString("HOST")
	fmt.Println("Dialling ", host)
	client, err = ethclient.Dial(host)
	chkErr("Get Client", err)

	multiSaleAddress := common.HexToAddress(viper.GetString("MULTISALE"))
	multiTokenAddress := common.HexToAddress(viper.GetString("MULTITOKEN"))
	saleAddress := common.HexToAddress(viper.GetString("SALE"))
	tokenAddress := common.HexToAddress(viper.GetString("TOKEN"))

	multisale, err := pMintyMultiSale.NewPMintyMultiSale(multiSaleAddress, client)
	multitoken, err := pMintyMultiToken.NewPMintyMultiToken(multiTokenAddress, client)
	sale, err := pSale.NewPMintysale(saleAddress, client)
	token, err := pMintyUnique.NewPMintyUnique(tokenAddress, client)

	emptyAddresses := []common.Address{}
	emptyIds := []*big.Int{}

	tokenOpts := bind.FilterOpts{Start: 0}
	token.FilterTransfer(&tokenOpts, emptyAddresses, emptyAddresses, emptyIds)

	multitokenOpts := bind.FilterOpts{Start: 15999837}
	multitoken.FilterTransferSingle(&multitokenOpts, emptyAddresses, emptyAddresses, emptyAddresses)

	mutisaleOpts := bind.FilterOpts{Start: 16042696}
	multisale.FilterNewOffer(&mutisaleOpts)

	saleOpts := bind.FilterOpts{Start: 0}
	sale.FilterNewOffer(&saleOpts)
	sale.FilterOfferAccepted(&saleOpts)
	sale.FilterResaleOffer(&saleOpts)
	sale.FilterPayment(&saleOpts)

	wg.Add(1)
	go uniqueTransfers(&token.PMintyUniqueFilterer, &wg)
	wg.Add(1)
	go multiTransfers(&multitoken.PMintyMultiTokenFilterer, multiTokenAddress, &wg)

	wg.Wait()
}
