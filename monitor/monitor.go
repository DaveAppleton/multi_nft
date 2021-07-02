package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"art.minty.monitor/etherdb"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	client   *ethclient.Client
	database *string = flag.String("database", "", "which DB to load")
	newMulti         = flag.String("new_multi", "", "new multi token")
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

func chkErrX(msg string, err error, term bool) bool {
	if err == nil {
		return false
	}
	fmt.Println(msg, err)
	if term {
		log.Fatal(msg, err)
	}
	log.Println(msg, err)
	return true
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
	etherdb.InitDB(viper.GetString("TOKEN_HOST"))

	if len(*newMulti) > 0 {
		fmt.Println("create new Multi Table : ", *newMulti)
		err := etherdb.CreateMultiTokenTable()
		if err != nil {
			fmt.Println("Status : ", err)
		}
		os.Exit(0)
	}

	wg.Add(1)
	go uniqueTransfers(client, &wg)
	wg.Add(1)
	go uniqueSales(client, &wg)
	wg.Add(1)
	go multiTransfers(client, &wg)
	wg.Add(1)
	go multiSales(client, &wg)

	wg.Wait()
}
