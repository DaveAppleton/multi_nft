package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"art.minty.backend/pSale"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func errorReturn(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}
	fmt.Fprintln(w, err)
	return true

}

func showArtists(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Artists")
	var artistResponse struct {
		Artists []Artist
		Max     int
		Start   int
		Next    int
		Prev    int
		End     int
	}
	var err error
	artistResponse.Start, err = strconv.Atoi(r.FormValue("start"))
	if err != nil {
		artistResponse.Start = 0
	}
	artistResponse.Artists, artistResponse.Max, err = getArtists(artistResponse.Start)
	if errorReturn(w, err) {
		return
	}
	fmt.Println("max=", artistResponse.Max, "artists=", len(artistResponse.Artists))
	artistResponse.End = artistResponse.Start + len(artistResponse.Artists)
	artistResponse.Next = artistResponse.Start + 50
	for j := 0; j < 50; j++ {
		if artistResponse.Start+j >= artistResponse.Max {
			artistResponse.Next = -1
			break
		}
	}
	fmt.Println("Artists")
	index, err := template.ParseFiles("files/main.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if artistResponse.Start > 0 {
		artistResponse.Prev = artistResponse.Start - 50
		if artistResponse.Prev < 0 {
			artistResponse.Prev = 0
		} else {
			artistResponse.Prev = -1
		}
	} else {
		artistResponse.Prev = -1
	}
	artistResponse.Next = artistResponse.Start + len(artistResponse.Artists)
	if err = index.Execute(w, artistResponse); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

func approveArtist(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ApproveArtist")
	artistName := r.FormValue("name")
	pSaleA := common.HexToAddress(viper.GetString("P_SALE"))
	client, err := ethclient.Dial(viper.GetString("HOST"))
	if errorReturn(w, err) {
		return
	}
	pSaleC, err := pSale.NewPMintysale(pSaleA, client)
	if errorReturn(w, err) {
		return
	}
	fmt.Println("finding ", *userName)
	artistID, artistAddr, err := findArtistAddress(artistName)
	if errorReturn(w, err) {
		return
	}
	addr := common.HexToAddress(artistAddr)
	fmt.Println("Approving ", artistID, artistName, " at ", addr.Hex())

	tx, err := pSaleC.SetMinter(bankerTx, addr, true)
	chkErr("SetMinter", err)
	fmt.Println(tx.Hash().Hex())
	log.Println("set minter ", artistName, artistID, addr.Hex(), *database)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	rct, err := bind.WaitMined(ctx, client, tx)
	if errorReturn(w, err) {
		return
	}
	fmt.Println("status", rct.Status)
	var updated bool
	if rct.Status == 1 {
		updated, err = approveUser(artistID)
		if errorReturn(w, err) {
			return
		}
	}

	approveStatus := struct {
		Network string
		Status  int
		User    string
		TxId    string
		Updated bool
	}{
		Network: *database,
		Status:  int(rct.Status),
		User:    artistName,
		TxId:    tx.Hash().Hex(),
		Updated: updated,
	}

	index, err := template.ParseFiles("files/approveUser.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err = index.Execute(w, approveStatus); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

type Wallet struct {
	Num     int
	Address string
}

type Pool struct {
	Num  int
	Name string
}

func project(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Wallets []Wallet
		Pools   []Pool
	}
	for j := 1; j < 7; j++ {
		data.Wallets = append(data.Wallets, Wallet{j, ""})
	}
	for j := 0; j < 6; j++ {
		data.Pools = append(data.Pools, Pool{j, fmt.Sprint("Pool ", j+1)})
	}
	index, err := template.ParseFiles("files/project.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err = index.Execute(w, data); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}
