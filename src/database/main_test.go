package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"oceanus/src/config"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	appConfig, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err = sql.Open(appConfig.DBDriver, appConfig.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
