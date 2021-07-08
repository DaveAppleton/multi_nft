package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"art.minty.monitor/etherdb"
	"github.com/spf13/viper"
)

func chkErrW(w http.ResponseWriter, msg string, err error) bool {
	if err == nil {
		return false
	}
	fmt.Fprintln(w, msg, err)
	return true
}

func projectIndex(w http.ResponseWriter, r *http.Request) {
	var data struct {
		ERC721s  []etherdb.Lookup
		ERC1155s []etherdb.Lookup
		Network  string
	}
	var err error
	data.ERC721s, err = etherdb.GetAllUniqueTokenIds()
	if chkErrW(w, "GetUniqueTokenIds", err) {
		return
	}
	data.ERC1155s, err = etherdb.GetAllMultiTokenIds()
	if chkErrW(w, "GetMultiTokenIds", err) {
		return
	}
	data.Network = viper.GetString("NETWORK")

	index, err := template.ParseFiles("files/main.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	if err = index.Execute(w, data); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

func timeX(tm uint64) string {
	return time.Unix(int64(tm), 0).Format("02 Jan 2006 15:04")
}

func projectUnique(w http.ResponseWriter, r *http.Request) {
	var err error
	var tt etherdb.UniqueTokenTransfer
	var data struct {
		Transfers []etherdb.UniqueTokenTransfer
		Lookup    etherdb.Lookup
		Name      string
		PrevPage  string
		NextPage  string
		Count     int
		Search    string
		Searching bool
		TokenID   int
		Tokening  bool
		TxHash    string
		TxIng     bool
	}

	data.TxHash = r.FormValue("tx")
	data.TxIng = data.TxHash != ""

	var t = template.FuncMap{"timeX": timeX}

	data.Search = r.FormValue("search")

	tokenId, err := strconv.Atoi(r.FormValue("token_id"))
	if err == nil {
		data.TokenID = tokenId
		data.Tokening = true
	}

	start, err := strconv.Atoi(r.FormValue("start"))
	if err != nil {
		start = 0
	}
	if start > 0 {
		nStart := start - 50
		if nStart < 0 {
			nStart = 0
		}
		data.PrevPage = fmt.Sprintf(r.URL.Path+"?start=%d", nStart)
	}
	data.Name = viper.GetString("NAME")
	expected := "/monitor/unique/"
	idStr := r.URL.Path[len(expected):]
	fmt.Println(idStr)
	tt.LookupID, err = strconv.Atoi(idStr)
	if chkErrW(w, "Convert LookupID", err) {
		return
	}
	data.Lookup, err = etherdb.LookupFromID(tt.LookupID)
	if chkErrW(w, "Get LookupFromID", err) {
		return
	}
	if start < data.Count-50 {
		data.NextPage = fmt.Sprintf(r.URL.Path+"?start=%d", start+50)
	}

	if data.TxIng {
		tt.TxHash = data.TxHash
		data.Count, err = tt.CountOfThisTxHash()
		if chkErrW(w, "Get Count of this TxHash", err) {
			return
		}
		data.Transfers, err = tt.FindByTxHash()
	} else if data.Tokening {
		tt.TokenID = uint64(data.TokenID)
		data.Count, err = tt.CountOfThisTokenID()
		if chkErrW(w, "Get Count of this TokenID", err) {
			return
		}
		data.Transfers, err = tt.FindByTokenId()
	} else if len(data.Search) == 0 {
		data.Count, err = tt.CountOfThisToken()
		if chkErrW(w, "Get Count of this Token", err) {
			return
		}
		data.Transfers, err = tt.ListAll(start)
	} else {
		data.Count, err = tt.SearchCountOfThisToken(data.Search)
		if chkErrW(w, "Get Count of this TokenD", err) {
			return
		}
		data.Searching = true
		data.Transfers, err = tt.FindAllByAddress(data.Search)
	}

	if chkErrW(w, "Get  Transfers", err) {
		return
	}
	f, err := ioutil.ReadFile("files/unique.html")
	if chkErrW(w, "Get Unique Transfers", err) {
		return
	}
	index, err := template.New("xx").Funcs(t).Parse(string(f))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err = index.Execute(w, data); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func projectMulti(w http.ResponseWriter, r *http.Request) {
	var err error
	var tt etherdb.MultiTokenTransfer
	var data struct {
		Transfers []etherdb.MultiTokenTransfer
		Lookup    etherdb.Lookup
		Name      string
		PrevPage  string
		NextPage  string
		Count     int
		Search    string
		Searching bool
		TokenID   int
		Tokening  bool
		TxHash    string
		TxIng     bool
	}

	data.TxHash = r.FormValue("tx")
	data.TxIng = data.TxHash != ""

	var t = template.FuncMap{"timeX": timeX}

	data.Search = r.FormValue("search")

	tokenId, err := strconv.Atoi(r.FormValue("token_id"))
	if err == nil {
		data.TokenID = tokenId
		data.Tokening = true
	}

	start, err := strconv.Atoi(r.FormValue("start"))
	if err != nil {
		start = 0
	}
	if start > 0 {
		nStart := start - 50
		if nStart < 0 {
			nStart = 0
		}
		data.PrevPage = fmt.Sprintf(r.URL.Path+"?start=%d", nStart)
	}
	data.Name = viper.GetString("NAME")
	expected := "/monitor/unique/"
	idStr := r.URL.Path[len(expected):]
	fmt.Println(idStr)
	tt.LookupID, err = strconv.Atoi(idStr)
	if chkErrW(w, "Convert LookupID", err) {
		return
	}
	data.Lookup, err = etherdb.LookupFromID(tt.LookupID)
	if chkErrW(w, "Get LookupFromID", err) {
		return
	}
	if start < data.Count-50 {
		data.NextPage = fmt.Sprintf(r.URL.Path+"?start=%d", start+50)
	}

	if data.TxIng {
		tt.TxHash = data.TxHash
		data.Count, err = tt.CountOfThisTxHash()
		if chkErrW(w, "Get Count of this TxHash", err) {
			return
		}
		data.Transfers, err = tt.FindByTxHash()
	} else if data.Tokening {
		tt.TokenID = uint64(data.TokenID)
		data.Count, err = tt.CountOfThisTokenID()
		if chkErrW(w, "Get Count of this TokenID", err) {
			return
		}
		data.Transfers, err = tt.FindByTokenId()
	} else if len(data.Search) == 0 {
		data.Count, err = tt.CountOfThisToken()
		if chkErrW(w, "Get Count of this Token", err) {
			return
		}
		data.Transfers, err = tt.ListAll(start)
	} else {
		data.Count, err = tt.SearchCountOfThisToken(data.Search)
		if chkErrW(w, "Get Count of this TokenD", err) {
			return
		}
		data.Searching = true
		data.Transfers, err = tt.FindAllByAddress(data.Search)
	}

	if chkErrW(w, "Get  Transfers", err) {
		return
	}
	f, err := ioutil.ReadFile("files/multi.html")
	if chkErrW(w, "Get Unique Transfers", err) {
		return
	}
	index, err := template.New("xx").Funcs(t).Parse(string(f))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err = index.Execute(w, data); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}
