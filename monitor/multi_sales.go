package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"art.minty.monitor/etherdb"
	"art.minty.monitor/pMintyMultiSale"
	"github.com/DaveAppleton/etherUtils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func filtErr(msg string, err error) bool {
	if err != nil {
		if err.Error() == "request failed or timed out" {
			fmt.Println("timed out")
			return false
		}
		chkErrX(msg, err, false)
		return false
	}
	return true
}

/*

type MultiSale struct {
	Token    common.Address
	TokenId  *big.Int
	Seller    common.Address
	Buyer     common.Address
	Quantity *big.Int
	Price    *big.Int
	Hash     string
	Position *big.Int
}



type PMintyMultiSaleNewOffer struct {
	Token    common.Address
	TokenId  *big.Int
	Owner    common.Address
	Quantity *big.Int
	Price    *big.Int
	Hash     string
	Raw      types.Log // Blockchain specific contextual infos
}

type PMintyMultiSaleOfferAccepted struct {
	Buyer    common.Address
	Token    common.Address
	TokenId  *big.Int
	Pos      *big.Int
	Quantity *big.Int
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

type PMintyMultiSaleResaleOffer struct {
	Token    common.Address
	TokenId  *big.Int
	Quantity *big.Int
	Owner    common.Address
	Price    *big.Int
	Position *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}
*/
func multiSales(client *ethclient.Client, wg *sync.WaitGroup) {
	delta := uint64(2000)
	var almostUptoDate bool
	fmt.Println("Multi Sales")
	defer wg.Done()

	for {
		salez, err := etherdb.GetAllMultiSaleIds()
		if chkErrX("get All Multi Salez", err, false) {
			return
		}
		if len(salez) == 0 {
			fmt.Println("NO MULTI SALES FOUND")
			log.Println("NO MULTI SALES FOUND")
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

			multisale, err := pMintyMultiSale.NewPMintyMultiSale(addr, client)
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
			filt1, err := multisale.FilterNewOffer(&tokenOpts)
			fmt.Println("Offer New")
			if filtErr("offer new", err) {
				for filt1.Next() {
					fmt.Println("MS Offer New", filt1.Event.Raw.TxHash.Hex(), filt1.Event.Token.Hex(), filt1.Event.Owner.Hex(), filt1.Event.TokenId, filt1.Event.Quantity, etherUtils.EtherToStr(filt1.Event.Price), filt1.Event.Hash)
					log.Println("MS Offer New", filt1.Event.Raw.TxHash.Hex(), filt1.Event.Token.Hex(), filt1.Event.Owner.Hex(), filt1.Event.TokenId, filt1.Event.Quantity, etherUtils.EtherToStr(filt1.Event.Price), filt1.Event.Hash)
					timestamp, _ := blockTimeStamp(filt1.Event.Raw.BlockHash)
					sale := etherdb.NewMultiSale(
						salez[j].ID,
						"OfferNew",
						filt1.Event.TokenId.Uint64(),
						filt1.Event.Raw.BlockNumber,
						filt1.Event.Raw.TxIndex,
						filt1.Event.Raw.TxHash.Hex(),
						filt1.Event.Token.Hex(),
						"",
						filt1.Event.Owner.Hex(),
						filt1.Event.Quantity.Uint64(),
						etherUtils.EtherToStr(filt1.Event.Price),
						0,
						filt1.Event.Hash,
						timestamp)

					chkErrX("Offer New", sale.Add(), false)
				}
			}
			fmt.Println("Offer Accepted")
			filt2, err := multisale.FilterOfferAccepted(&tokenOpts)
			if filtErr("offer accepted", err) {
				for filt2.Next() {
					fmt.Println("Multi Sale", filt2.Event.Raw.TxHash.Hex(), filt2.Event.Token.Hex(), filt2.Event.Buyer.Hex(), filt2.Event.Pos, filt2.Event.Quantity, etherUtils.EtherToStr(filt2.Event.Price))
					log.Println("Multi Sale", filt2.Event.Raw.TxHash.Hex(), filt2.Event.Token.Hex(), filt2.Event.Buyer.Hex(), filt2.Event.Pos, filt2.Event.Quantity, etherUtils.EtherToStr(filt2.Event.Price))
					timestamp, _ := blockTimeStamp(filt2.Event.Raw.BlockHash)
					sale := etherdb.NewMultiSale(
						salez[j].ID,
						"OfferAccepted",
						filt2.Event.TokenId.Uint64(),
						filt2.Event.Raw.BlockNumber,
						filt2.Event.Raw.TxIndex,
						filt2.Event.Raw.TxHash.Hex(),
						filt2.Event.Token.Hex(),
						filt2.Event.Buyer.Hex(),
						"",
						filt2.Event.Quantity.Uint64(),
						etherUtils.EtherToStr(filt2.Event.Price),
						filt2.Event.Pos.Uint64(),
						"",
						timestamp)

					chkErrX("Sale Accepted", sale.Add(), false)
				}
			}
			fmt.Println("Resale")
			filt3, err := multisale.FilterResaleOffer(&tokenOpts)
			if filtErr("resale offer", err) {
				for filt3.Next() {
					fmt.Println("Multi Sale", filt3.Event.Raw.TxHash.Hex(), filt3.Event.Token.Hex(), filt3.Event.Owner.Hex(), filt3.Event.TokenId, filt3.Event.Position, filt3.Event.Quantity, etherUtils.EtherToStr(filt3.Event.Price))
					log.Println("Multi Sale", filt3.Event.Raw.TxHash.Hex(), filt3.Event.Token.Hex(), filt3.Event.Owner.Hex(), filt3.Event.TokenId, filt3.Event.Position, filt3.Event.Quantity, etherUtils.EtherToStr(filt3.Event.Price))
					timestamp, _ := blockTimeStamp(filt3.Event.Raw.BlockHash)
					sale := etherdb.NewMultiSale(
						salez[j].ID,
						"ResaleOffer",
						filt3.Event.TokenId.Uint64(),
						filt3.Event.Raw.BlockNumber,
						filt3.Event.Raw.TxIndex,
						filt3.Event.Raw.TxHash.Hex(),
						filt3.Event.Token.Hex(),
						"",
						filt3.Event.Owner.Hex(),
						filt3.Event.Quantity.Uint64(),
						etherUtils.EtherToStr(filt3.Event.Price),
						filt3.Event.Position.Uint64(),
						"",
						timestamp)

					chkErrX("Resale", sale.Add(), false)

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
