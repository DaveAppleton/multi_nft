package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	action *string = flag.String("action", "", "what to do")       // list
	search *string = flag.String("search", "", "what to look for") // LaylaHifi
)

func initViper() {
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("./config.json")
		log.Fatal("config.json", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("viper config changed", e.Name)
	})
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
	flag.Parse()
	if *action == "list" {

		records, err := getProjectsForUser(*search)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		for _, rec := range records {
			fmt.Println(
				rec.Name,
				rec.Address,
				rec.Title,
				rec.ProjectID,
				rec.ProjectToken,
			)
		}
	}

}
