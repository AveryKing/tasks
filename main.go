package main

import (
	"database/sql"
	"github.com/AveryKing/tasks/api"
	db "github.com/AveryKing/tasks/db/sqlc"
	"github.com/AveryKing/tasks/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("error loading config: ", err)
	}

	var testDB *sql.DB

	testDB, _ = sql.Open(config.DBDriver, config.DBSource)
	err = testDB.Ping()
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}
	testQueries := db.New(testDB)

	server, err := api.NewServer(testQueries)
	if err != nil {
		log.Fatal("error creating gin server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("error starting gin server: ", err)
	}
}
