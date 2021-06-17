package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
