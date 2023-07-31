package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/vivekdubey/fiber-api/internal/config"
)

var testQueries *Queries

const dbDriver = "postgres"

var testDB *sql.DB

func TestMain(m *testing.M) {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	testDB, err = sql.Open(dbDriver, c.DBurl)
	if err != nil {
		log.Fatal("Can't connect to db")
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
