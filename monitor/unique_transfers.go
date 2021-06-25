package main

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"art.minty.monitor/pMintyUnique"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func uniqueTransfers(token *pMintyUnique.PMintyUniqueFilterer, wg *sync.WaitGroup) {
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
