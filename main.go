package main

import (
	"database/sql"
	"github.com/AveryKing/tasks/api"
	"github.com/AveryKing/tasks/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("error loading config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	server, err := api.NewServer(conn)
	if err != nil {
		log.Fatal("error creating gin server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("error starting gin server: ", err)
	}
}
