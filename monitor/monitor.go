package main

/*
CREATE DATABASE blockchain;
CREATE USER token with password 'erc20';
GRANT ALL PRIVILEGES ON DATABASE blockchain TO token;

./art.minty.monitor -database minty-node -new_db true

Then set up tokens in lookup table
*/

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"art.minty.monitor/etherdb"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	client          *ethclient.Client
	database        *string = flag.String("database", "", "which DB to load")
	newDatabase     *bool   = flag.Bool("new_db", false, "new database")
	initUniqueToken *string = flag.String("init_unique_token", "", "load unique tokens to Lookup")
	initMultiToken  *string = flag.String("init_multi_token", "", "load unique tokens to Lookup")
	initUniqueSale  *string = flag.String("init_unique_sale", "", "load unique tokens to Lookup")
	initMultiSale   *string = flag.String("init_multi_sale", "", "load unique tokens to Lookup")
	web             *bool   = flag.Bool("web", false, "load webserver")
	monitor         *bool   = flag.Bool("monitor", false, "token monitor")
)

func readCSV(input string) (data [][]string, err error) {
	sf, err := os.Open(input)
	if err != nil {
		fmt.Println(err, "[", input, "]")
		log.Fatal(err, input)
	}
	data, err = csv.NewReader(sf).ReadAll()
	return
}

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
	etherdb.InitLockingDB(viper.GetString("DB_HOST"))

	if *newDatabase {
		fmt.Println("create Lookup : ")
		err = etherdb.CreateLookupTable()
		if err != nil {
			fmt.Println("Lookup : ", err)
		}
		fmt.Println("create Multi Token : ")
		err := etherdb.CreateMultiTokenTable()
		if err != nil {
			fmt.Println("Multi Token : ", err)
		}
		fmt.Println("create Unique Token : ")
		err = etherdb.CreateUniqueTokenTable()
		if err != nil {
			fmt.Println("Unique Token : ", err)
		}
		fmt.Println("create Multi Sale : ")
		err = etherdb.CreateMultiSaleTable()
		if err != nil {
			fmt.Println("Multi Sale : ", err)
		}
		fmt.Println("create Unique Sale : ")
		err = etherdb.CreateUniqueSaleTable()
		if err != nil {
			fmt.Println("Unique Sale : ", err)
		}
		fmt.Println("create Unique Owners : ")
		err = etherdb.CreateUniqueOwnersTable()
		if err != nil {
			fmt.Println("Unique Owners : ", err)
		}

		os.Exit(0)
	}

	type adder struct {
		ParamPointer *string
		FuncPointer  func(uint64, string, string) (*etherdb.Lookup, error)
	}

	adders := []adder{
		{initUniqueToken, etherdb.AddUniqueTokenToLookup},
		{initMultiToken, etherdb.AddMultiTokenToLookup},
		{initUniqueSale, etherdb.AddUniqueSaleToLookup},
		{initMultiSale, etherdb.AddMultiSaleToLookup},
	}

	loaded := false
	for _, adx := range adders {
		pp := *adx.ParamPointer
		if len(pp) > 0 {
			data, err := readCSV(pp)
			if err != nil {
				log.Fatal(err)
			}
			for _, line := range data {
				startAt, err := strconv.ParseUint(line[0], 10, 64)
				if err != nil {
					log.Println("Adding ", pp, "to", *database, line[0], err)
					continue
				}
				_, err = adx.FuncPointer(startAt, line[1], line[2])
				if err != nil {
					log.Println("Adding ", pp, "to", *database, line[0], line[1], line[2], err)
					continue
				}
			}
			loaded = true
		}
	}
	if loaded {
		os.Exit(0)
	}
	if *monitor {
		fmt.Println("Starting Monitor")
		wg.Add(1)
		go uniqueTransfers(client, &wg)
		wg.Add(1)
		go uniqueSales(client, &wg)
		wg.Add(1)
		go multiTransfers(client, &wg)
		wg.Add(1)
		go multiSales(client, &wg)
		wg.Add(1)
		go lockingMonitor(client, &wg)
	}
	if *web {
		fmt.Println("Starting Web")
		// web section
		http.HandleFunc("/monitor/", projectIndex)
		http.HandleFunc("/monitor/unique/", projectUnique)
		http.HandleFunc("/monitor/multi/", projectMulti)
		http.HandleFunc("/monitor/usales/", uniqueSalePage)
		http.HandleFunc("/monitor/msales/", multiSalePage)
		http.Handle("/inc/", http.StripPrefix("/inc", http.FileServer(http.Dir("./public"))))
		err = http.ListenAndServe(viper.GetString("PORT"), nil)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
	wg.Wait()
}

// daily tx
// select id, name, image from project_arts where project_id = 1 order by id;
//
// select tokenid,operation,price,timestamp from unique_sale order by timestamp;
//
