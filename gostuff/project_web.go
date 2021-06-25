package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"art.minty.backend/pMintyMultiToken"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func projectCheck(w http.ResponseWriter, r *http.Request) {
	client, err := ethclient.Dial(viper.GetString("HOST"))
	if errorReturn(w, err) {
		return
	}
	//name := r.FormValue("name")
	projectID, err := strconv.Atoi(r.FormValue("project"))
	if errorReturn(w, err) {
		return
	}
	token, err := getProjectToken(projectID)
	if errorReturn(w, err) {
		return
	}
	if *database == "staging" {

	} else {
		pmtAddress := common.HexToAddress(viper.GetString(token))
		pmtContract, err := pMintyMultiToken.NewPMintyMultiToken(pmtAddress, client)
		if errorReturn(w, err) {
			return
		}
		_ = pmtContract
	}

}

func projectArt(w http.ResponseWriter, r *http.Request) {
	var data struct {
		PArts []ProjectArt
	}
	projectID, err := strconv.Atoi(r.FormValue("project"))
	if errorReturn(w, err) {
		return
	}
	data.PArts, err = getProjectArt(projectID)
	if errorReturn(w, err) {
		return
	}
	index, err := template.ParseFiles("files/project_arts.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err = index.Execute(w, data); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func setPool(w http.ResponseWriter, r *http.Request) {

	pArtID, err := strconv.Atoi(r.FormValue("project_art"))
	if errorReturn(w, err) {
		return
	}
	poolID, err := strconv.Atoi(r.FormValue("value"))
	if errorReturn(w, err) {
		return
	}
	fmt.Println("setPiilID ", projectID, poolID)
	updated, err := setPoolID(pArtID, poolID)
	if errorReturn(w, err) {
		return
	}
	approveStatus := struct {
		ProjectArt int
		Pool       int
		Updated    bool
	}{
		ProjectArt: pArtID,
		Pool:       poolID,
		Updated:    updated,
	}
	index, err := template.ParseFiles("files/updatePool.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err = index.Execute(w, approveStatus); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}
