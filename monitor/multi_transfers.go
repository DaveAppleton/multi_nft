package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"art.minty.monitor/etherdb"
	"art.minty.monitor/pMintyMultiToken"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func multiTransfers(client *ethclient.Client, wg *sync.WaitGroup) {
	fmt.Println("Multi Transfers")
	defer wg.Done()
	emptyAddresses := []common.Address{}

	//start := uint64(15992772)

	for {
		tokenz, err := etherdb.GetAllMultiTokenIds() // []{address,int,max}
		if chkErrX("get All Multi Tokenz", err, false) {
			return
		}
		if len(tokenz) == 0 {
			fmt.Println("NO MULTI TOKENS FOUND")
			log.Println("NO MULTI TOKENS FOUND")
			return
		}
		for j := 0; j < len(tokenz); j++ {

			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()
			currentBlock, err := client.BlockNumber(ctx)
			if chkErrX("BlockNumber", err, false) {
				return
			}

			addr := common.HexToAddress(tokenz[j].Address)
			token, err := pMintyMultiToken.NewPMintyMultiToken(addr, client)
			if chkErrX("restore "+tokenz[j].Description, err, false) {
				return
			}

			tokenOpts := bind.FilterOpts{Start: tokenz[j].LastChecked + 1}
			end := currentBlock - 60 // allow fo reorgs

			if end <= tokenOpts.Start {
				fmt.Println("skipping")
				if len(tokenz) < 2 {
					time.Sleep(5 * time.Minute)
				}
				continue
			}

			if end > tokenOpts.Start+10000 {
				end = tokenOpts.Start + 10000
			}
			tokenOpts.End = &end
			fmt.Println("filtering", tokenz[j].Description, "at", addr.Hex(), "from ", tokenOpts.Start, "to", *tokenOpts.End)
			filt, err := token.FilterTransferSingle(&tokenOpts, emptyAddresses, emptyAddresses, emptyAddresses)
			if err != nil {
				if err.Error() == "request failed or timed out" {
					fmt.Println("timed out")
					time.Sleep(30 * time.Second)
					continue
				}
				chkErrX("MultiFilterNext", err, false)
				return
			}
			for filt.Next() {
				timestamp, _ := blockTimeStamp(filt.Event.Raw.BlockHash)
				data := etherdb.NewMultiToken(
					tokenz[j].ID,
					filt.Event.Id.Uint64(),
					filt.Event.Raw.BlockNumber,
					filt.Event.Raw.TxIndex,
					filt.Event.Raw.TxHash.Hex(),
					filt.Event.Operator.Hex(),
					filt.Event.From.Hex(),
					filt.Event.To.Hex(),
					filt.Event.Value.Uint64(),
					timestamp,
				)
				if err = data.AddIfNotFound(); err != nil {
					log.Println("add multi", err)
					fmt.Println("add multi", err)
				}
				fmt.Println("Multi", addr.Hex(), filt.Event.Raw.TxHash.Hex(), filt.Event.Operator.Hex(), filt.Event.From.Hex(), filt.Event.To.Hex(), filt.Event.Id, filt.Event.Value)
				log.Println("Multi", addr.Hex(), filt.Event.Raw.TxHash.Hex(), filt.Event.Operator.Hex(), filt.Event.From.Hex(), filt.Event.To.Hex(), filt.Event.Id, filt.Event.Value)
				//start = filt.Event.Raw.BlockNumber + 1
			}
			fmt.Println("Update LastChecked to", *tokenOpts.End)
			chkErrX("update LastChecked", tokenz[j].UpdateLastChecked(*tokenOpts.End), false)
		}
	}
}
