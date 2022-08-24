package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/hariprathap-hp/backend_masterclass/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

/*const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
)*/

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal(err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln(err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
