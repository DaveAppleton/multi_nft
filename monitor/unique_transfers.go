package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"art.minty.monitor/etherdb"
	"art.minty.monitor/pMintyUnique"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
type PMintyUniqueTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}
*/

func uniqueTransfers(client *ethclient.Client, wg *sync.WaitGroup) {
	fmt.Println("Unique Transfers")
	defer wg.Done()
	emptyAddresses := []common.Address{}

	//start := uint64(15992772)

	for {
		tokenz, err := etherdb.GetAllUniqueTokenIds() // []{address,int,max}
		if chkErrX("get All Unique Tokenz", err, false) {
			return
		}
		if len(tokenz) == 0 {
			fmt.Println("NO UNIQUE TOKENS FOUND")
			log.Println("NO UNIQUE TOKENS FOUND")
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
			token, err := pMintyUnique.NewPMintyUnique(addr, client)
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
			filt, err := token.FilterTransfer(&tokenOpts, emptyAddresses, emptyAddresses, nil)
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
				data := etherdb.NewUniqueToken(
					tokenz[j].ID,
					filt.Event.Raw.BlockNumber,
					filt.Event.Raw.TxIndex,
					filt.Event.Raw.TxHash.Hex(),
					filt.Event.From.Hex(),
					filt.Event.To.Hex(),
					filt.Event.TokenId.Uint64(),
					timestamp,
				)
				if err = data.AddIfNotFound(); err != nil {
					log.Println("add unique", err)
					fmt.Println("add unique", err)
				}
				fmt.Println("Unique", addr.Hex(), filt.Event.Raw.TxHash.Hex(), filt.Event.From.Hex(), filt.Event.To.Hex())
				log.Println("Unique", addr.Hex(), filt.Event.Raw.TxHash.Hex(), filt.Event.From.Hex(), filt.Event.To.Hex())
				//start = filt.Event.Raw.BlockNumber + 1
			}
			fmt.Println("Update LastChecked to", *tokenOpts.End)
			chkErrX("update LastChecked", tokenz[j].UpdateLastChecked(*tokenOpts.End), false)
		}
	}
}

/*	token *pMintyUnique.PMintyUniqueFilterer, wg *sync.WaitGroup) {
	fmt.Println("Unique Transfers")
	defer wg.Done()
	emptyAddresses := []common.Address{}
	emptyIds := []*big.Int{}
	start := uint64(15992772)
	tokenOpts := bind.FilterOpts{Start: start}
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		currentBlock, err := client.BlockNumber(ctx)
		if err != nil {
			chkErrX("BlockNumber", err, false)
			return
		}

		end := currentBlock - 60 // allow fo reorgs

		if end <= tokenOpts.Start {
			time.Sleep(30 * time.Second)
			continue
		}

		if end > tokenOpts.Start+50000 {
			end = tokenOpts.Start + 50000
		}

		tokenOpts := bind.FilterOpts{Start: start, End: &end}
		filt, err := token.FilterTransfer(&tokenOpts, emptyAddresses, emptyAddresses, emptyIds)
		if err != nil {
			if err.Error() == "request failed or timed out" {
				time.Sleep(30 * time.Second)
				continue
			}
			chkErrX("UniqueFilterNext", err, false)
			return
		}
		for filt.Next() {
			fmt.Println(filt.Event.From.Hex(), filt.Event.To.Hex(), filt.Event.TokenId)
			start = filt.Event.Raw.BlockNumber + 1
		}
		tokenOpts.Start = end + 1
	}
}
*/
