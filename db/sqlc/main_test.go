package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	//"github.com/techschool/simplebank/util"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/quickbank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
