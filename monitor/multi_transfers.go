package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"art.minty.monitor/pMintyMultiToken"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func multiTransfers(token *pMintyMultiToken.PMintyMultiTokenFilterer, addr common.Address, wg *sync.WaitGroup) {
	fmt.Println("Multi Transfers")
	defer wg.Done()
	emptyAddresses := []common.Address{}

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
		tokenOpts.End = &end

		filt, err := token.FilterTransferSingle(&tokenOpts, emptyAddresses, emptyAddresses, emptyAddresses)
		if err != nil {
			if err.Error() == "request failed or timed out" {
				time.Sleep(30 * time.Second)
				continue
			}
			chkErrX("MultiFilterNext", err, false)
			return
		}
		// Operator common.Address
		// From     common.Address
		// To       common.Address
		// Id       *big.Int
		// Value    *big.Int

		for filt.Next() {

			fmt.Println("Multi", addr.Hex(), filt.Event.Raw.TxHash.Hex(), filt.Event.Operator.Hex(), filt.Event.From.Hex(), filt.Event.To.Hex(), filt.Event.Id, filt.Event.Value)
			log.Println("Multi", addr.Hex(), filt.Event.Raw.TxHash.Hex(), filt.Event.Operator.Hex(), filt.Event.From.Hex(), filt.Event.To.Hex(), filt.Event.Id, filt.Event.Value)
			start = filt.Event.Raw.BlockNumber + 1
		}
		tokenOpts.Start = end + 1
	}
}
