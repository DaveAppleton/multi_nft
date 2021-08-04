package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"art.minty.monitor/etherdb"
	"art.minty.monitor/locking"
	"github.com/DaveAppleton/etherUtils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func lockingMonitor(client *ethclient.Client, wg *sync.WaitGroup) {
	delta := uint64(2000)
	var almostUptoDate bool
	fmt.Println("Locking")
	defer wg.Done()

	for {
		salez, err := etherdb.GetAllLockingIds()
		if chkErrX("get Locking Events", err, false) {
			return
		}
		if len(salez) == 0 {
			fmt.Println("NO LOCKING EVENTS FOUND")
			log.Println("NO LOCKING EVENTS FOUND")
			time.Sleep(5 * time.Minute)
			continue
		}

		for j := 0; j < len(salez); j++ {
			tokenOpts := bind.FilterOpts{Start: salez[j].LastChecked + 1}

			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()
			currentBlock, err := client.BlockNumber(ctx)
			if chkErrX("BlockNumber", err, false) {
				continue
			}
			fmt.Println("current block ", currentBlock)
			addr := common.HexToAddress(salez[j].Address)

			lockingContract, err := locking.NewLocking(addr, client) //.NewPMintySale(addr, client)
			if chkErrX("restore "+salez[j].Description, err, false) {
				continue
			}

			end := currentBlock - 60 // allow fo reorgs

			if end <= tokenOpts.Start {
				fmt.Println("skipping")
				continue
			}
			if end > tokenOpts.Start+delta {
				end = tokenOpts.Start + delta
			}
			tokenOpts.End = &end
			fmt.Println("filtering", salez[j].Description, "at", addr.Hex(), "from ", tokenOpts.Start, "to", *tokenOpts.End)
			filt1, err := lockingContract.FilterLockedBaseToken(&tokenOpts)
			fmt.Println("Locked Base Token")
			if filtErr("locked token", err) {
				for filt1.Next() {
					fmt.Println("Locked Token", filt1.Event.Raw.TxHash.Hex(), filt1.Event.Owner.Hex(), etherUtils.EtherToStr(filt1.Event.Amount))
					log.Println("Locked Token", filt1.Event.Raw.TxHash.Hex(), filt1.Event.Owner.Hex(), etherUtils.EtherToStr(filt1.Event.Amount))
					timestamp, _ := blockTimeStamp(filt1.Event.Raw.BlockHash)
					locky := etherdb.NewLockedTokenData(
						salez[j].ID,
						"Lock",
						filt1.Event.Raw.BlockNumber,
						filt1.Event.Raw.TxIndex,
						filt1.Event.Raw.TxHash.Hex(),
						filt1.Event.Owner.Hex(),
						etherUtils.EtherToStr(filt1.Event.Amount),
						timestamp)

					chkErrX("Locked Token", locky.Add(), false)
				}
			}
			filt2, err := lockingContract.FilterWithdrawBaseTokens(&tokenOpts)
			fmt.Println("Withdraw Base Token")
			if filtErr("Withdraw token", err) {
				for filt2.Next() {
					fmt.Println("Withdraw Token", filt2.Event.Raw.TxHash.Hex(), filt2.Event.Owner.Hex(), etherUtils.EtherToStr(filt2.Event.Amount))
					log.Println("Withdraw Token", filt2.Event.Raw.TxHash.Hex(), filt2.Event.Owner.Hex(), etherUtils.EtherToStr(filt2.Event.Amount))
					timestamp, _ := blockTimeStamp(filt2.Event.Raw.BlockHash)
					locky := etherdb.NewLockedTokenData(
						salez[j].ID,
						"Withdraw",
						filt2.Event.Raw.BlockNumber,
						filt2.Event.Raw.TxIndex,
						filt2.Event.Raw.TxHash.Hex(),
						filt2.Event.Owner.Hex(),
						etherUtils.EtherToStr(filt1.Event.Amount),
						timestamp)

					chkErrX("Withdraw Token", locky.Add(), false)
				}
			}

			fmt.Println("Unlock Request")
			filt3, err := lockingContract.FilterUnlockRequested(&tokenOpts)
			if filtErr("Unlock Requested", err) {
				for filt3.Next() {
					fmt.Println("Unlock Requested", filt3.Event.Raw.TxHash.Hex(), filt3.Event.Owner.Hex())
					log.Println("Unlock Requested", filt3.Event.Raw.TxHash.Hex(), filt3.Event.Owner.Hex())
					timestamp, _ := blockTimeStamp(filt2.Event.Raw.BlockHash)
					sale := etherdb.NewLockedTokenData(
						salez[j].ID,
						"UnlockRequest",
						filt3.Event.Raw.BlockNumber,
						filt3.Event.Raw.TxIndex,
						filt3.Event.Raw.TxHash.Hex(),
						filt3.Event.Owner.Hex(),
						"",
						timestamp)
					chkErrX("UnlockRequest", sale.Add(), false)
				}
			}

			fmt.Println("Unlock Cancelled")
			filt4, err := lockingContract.FilterUnlockCancelled(&tokenOpts)
			if filtErr("Unlock Cancelled", err) {
				for filt4.Next() {
					fmt.Println("Unlock Cancelled", filt3.Event.Raw.TxHash.Hex(), filt3.Event.Owner.Hex())
					log.Println("Unlock Cancelled", filt3.Event.Raw.TxHash.Hex(), filt3.Event.Owner.Hex())
					timestamp, _ := blockTimeStamp(filt2.Event.Raw.BlockHash)
					sale := etherdb.NewLockedTokenData(
						salez[j].ID,
						"UnlockCancel",
						filt4.Event.Raw.BlockNumber,
						filt4.Event.Raw.TxIndex,
						filt4.Event.Raw.TxHash.Hex(),
						filt4.Event.Owner.Hex(),
						"",
						timestamp)
					chkErrX("UnlockCancel", sale.Add(), false)
				}
			}

			fmt.Println("Update LastChecked to", *tokenOpts.End)
			chkErrX("update LastChecked", salez[j].UpdateLastChecked(*tokenOpts.End), false)
			almostUptoDate = *tokenOpts.End > currentBlock-delta
			fmt.Println(*tokenOpts.End, currentBlock-delta, almostUptoDate)

			if len(salez) < 2 && almostUptoDate {
				fmt.Println("snoozing")
				time.Sleep(5 * time.Minute)
			}
		}
	}

}
