package main

import (
	"database/sql"
	"log"

	"github.com/hariprathap-hp/backend_masterclass/api"
	db "github.com/hariprathap-hp/backend_masterclass/db/sqlc"
	"github.com/hariprathap-hp/backend_masterclass/util"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server")
	}
}
