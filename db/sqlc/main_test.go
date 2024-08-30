package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/quickbank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = New(connPool)

	os.Exit(m.Run())
}
