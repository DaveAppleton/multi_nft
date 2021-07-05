package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initViper() {
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("./config.json")
		log.Fatal("config.json", err)
	}
	viper.WatchConfig()

}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
	f, _ := os.Open("images/firstcard.png")
	io.Copy(w, f)
}

func main() {
	initViper()
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	fileServer := http.FileServer(http.Dir("./files"))
	http.Handle("/", fileServer)
	http.HandleFunc("/images/", imageHandler)
	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		log.Fatal(err)
	}
}
