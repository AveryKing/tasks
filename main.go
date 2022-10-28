package main

import (
	"github.com/AveryKing/tasks/api"
	"github.com/AveryKing/tasks/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("error loading config: ", err)
	}

	server, err := api.NewServer()
	if err != nil {
		log.Fatal("error creating gin server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("error starting gin server: ", err)
	}
}
