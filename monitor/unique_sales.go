package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"art.minty.monitor/etherdb"
	"art.minty.monitor/pSale"
	"github.com/DaveAppleton/etherUtils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func uniqueSales(client *ethclient.Client, wg *sync.WaitGroup) {
	delta := uint64(2000)
	var almostUptoDate bool
	fmt.Println("Multi Sales")
	defer wg.Done()

	for {
		salez, err := etherdb.GetAllUniqueSaleIds()
		if chkErrX("get All Unique Salez", err, false) {
			return
		}
		if len(salez) == 0 {
			fmt.Println("NO UNIQUE SALES FOUND")
			log.Println("NO UNIQUE SALES FOUND")
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
			addr := common.HexToAddress(salez[j].Address)

			psale, err := pSale.NewPMintysale(addr, client) //.NewPMintySale(addr, client)
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
			filt1, err := psale.FilterNewOffer(&tokenOpts)
			fmt.Println("Offer New")
			if filtErr("unique offer new", err) {
				for filt1.Next() {
					fmt.Println("MS Offer New", filt1.Event.Raw.TxHash.Hex(), filt1.Event.Owner.Hex(), filt1.Event.TokenId, etherUtils.EtherToStr(filt1.Event.Price), filt1.Event.Hash)
					log.Println("MS Offer New", filt1.Event.Raw.TxHash.Hex(), filt1.Event.Owner.Hex(), filt1.Event.TokenId, etherUtils.EtherToStr(filt1.Event.Price), filt1.Event.Hash)
					timestamp, _ := blockTimeStamp(filt1.Event.Raw.BlockHash)
					sale := etherdb.NewUniqueSale(
						salez[j].ID,
						"OfferNew",
						filt1.Event.TokenId.Uint64(),
						filt1.Event.Raw.BlockNumber,
						filt1.Event.Raw.TxIndex,
						filt1.Event.Raw.TxHash.Hex(),
						"",
						filt1.Event.Owner.Hex(),
						etherUtils.EtherToStr(filt1.Event.Price),
						filt1.Event.Hash,
						timestamp)

					chkErrX("unique Offer New", sale.Add(), false)
				}
			}
			fmt.Println("Offer Accepted")
			filt2, err := psale.FilterOfferAccepted(&tokenOpts)
			if filtErr("unique offer accepted", err) {
				for filt2.Next() {
					fmt.Println("Multi Sale", filt2.Event.Raw.TxHash.Hex(), filt2.Event.Buyer.Hex(), etherUtils.EtherToStr(filt2.Event.Price))
					log.Println("Multi Sale", filt2.Event.Raw.TxHash.Hex(), filt2.Event.Buyer.Hex(), etherUtils.EtherToStr(filt2.Event.Price))
					timestamp, _ := blockTimeStamp(filt2.Event.Raw.BlockHash)
					sale := etherdb.NewUniqueSale(
						salez[j].ID,
						"OfferAccepted",
						filt2.Event.TokenId.Uint64(),
						filt2.Event.Raw.BlockNumber,
						filt2.Event.Raw.TxIndex,
						filt2.Event.Raw.TxHash.Hex(),
						filt2.Event.Buyer.Hex(),
						"",
						etherUtils.EtherToStr(filt2.Event.Price),
						"",
						timestamp)

					chkErrX("unique Sale Accepted", sale.Add(), false)
				}
			}
			fmt.Println("Resale")
			filt3, err := psale.FilterResaleOffer(&tokenOpts)
			if filtErr("unique resale offer", err) {
				for filt3.Next() {
					fmt.Println("Multi Sale", filt3.Event.Raw.TxHash.Hex(), filt3.Event.Owner.Hex(), filt3.Event.TokenId, etherUtils.EtherToStr(filt3.Event.Price))
					log.Println("Multi Sale", filt3.Event.Raw.TxHash.Hex(), filt3.Event.Owner.Hex(), filt3.Event.TokenId, etherUtils.EtherToStr(filt3.Event.Price))
					timestamp, _ := blockTimeStamp(filt3.Event.Raw.BlockHash)
					sale := etherdb.NewUniqueSale(
						salez[j].ID,
						"ResaleOffer",
						filt3.Event.TokenId.Uint64(),
						filt3.Event.Raw.BlockNumber,
						filt3.Event.Raw.TxIndex,
						filt3.Event.Raw.TxHash.Hex(),
						"",
						filt3.Event.Owner.Hex(),
						etherUtils.EtherToStr(filt3.Event.Price),
						"",
						timestamp)

					chkErrX("unique Resale", sale.Add(), false)

				}
			}
			fmt.Println("Update LastChecked to", *tokenOpts.End)
			chkErrX("update LastChecked", salez[j].UpdateLastChecked(*tokenOpts.End), false)
			almostUptoDate = *tokenOpts.End > currentBlock-delta
			fmt.Println(*tokenOpts.End, currentBlock-delta, almostUptoDate)
		}
		if len(salez) < 2 && almostUptoDate {
			fmt.Println("snoozing")
			time.Sleep(5 * time.Minute)
		}
	}

}
