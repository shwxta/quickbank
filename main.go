package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/shwxta/quickbank/api"
	db "github.com/shwxta/quickbank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/quickbank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start((serverAddress))
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

}
