package api

import (
	"database/sql"
	db "github.com/AveryKing/tasks/db/sqlc"
	"github.com/AveryKing/tasks/util"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"testing"
)

func newTestServer(t *testing.T) (server *Server) {
	config, err := util.LoadConfig("../.")
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

	server, err = NewServer(testQueries)
	if err != nil {
		log.Fatal("error creating gin server: ", err)
	}
	return
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
