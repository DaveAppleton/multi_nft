package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/DaveAppleton/ether_go/ethKeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
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

func chkErr(msg string, err error) {
	if err == nil {
		return
	}
	fmt.Println(err)
	log.Fatal(err)
}

func main() {
	initViper()
	fmt.Println("clients")
	g_host := viper.GetString("G_HOST")
	fmt.Println(g_host)
	g_client, err := ethclient.Dial(g_host)
	chkErr("g connect", err)
	m_host := viper.GetString("M_HOST")
	fmt.Println(m_host)
	m_client, err := ethclient.Dial(m_host)
	chkErr("m connect", err)
	fmt.Println("keys")
	banker := ethKeys.NewKey("adminKeys/banker")
	err = banker.RestoreOrCreate()
	chkErr("keys", err)
	fmt.Println("Banker", banker.PublicKeyAsHexString())

	fmt.Println("G id")

	gBankerTx := bind.NewKeyedTransactor(banker.GetKey())
	chkErr("gtx", err)

	fmt.Println("M id")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	m_id, err := m_client.ChainID(ctx)
	if err != nil {
		log.Fatal("chain ID", err)
	}
	mBankerTx, err := bind.NewKeyedTransactorWithChainID(banker.GetKey(), m_id)
	chkErr("mtx", err)
	fmt.Println("deploy G")
	gAddr, gTx, _, err := DeployBridge(gBankerTx, g_client)
	chkErr("g deploy", err)
	fmt.Println("deploy M")
	mAddr, mTx, _, err := DeployBridge(mBankerTx, m_client)
	chkErr("m deploy", err)
	fmt.Println(gAddr.Hex(), gTx.Hash().Hex())
	fmt.Println(mAddr.Hex(), mTx.Hash().Hex())
}
