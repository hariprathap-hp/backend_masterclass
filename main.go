package main

import (
	"database/sql"
	"log"

	"github.com/hariprathap-hp/backend_masterclass/api"
	db "github.com/hariprathap-hp/backend_masterclass/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start("localhost:8080")
	if err != nil {
		log.Fatal("can not start server")
	}
}
