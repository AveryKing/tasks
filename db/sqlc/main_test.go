package db

import (
	"database/sql"
	"github.com/AveryKing/tasks/util"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("error loading config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("error connecting to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
