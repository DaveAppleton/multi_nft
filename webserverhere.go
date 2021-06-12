package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fileServer)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
