package database_tests

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"

	database "simple_bank/database/sqlc"
)

var accountOperationQueries *database.Queries

//* main entrypoint of unit testing inside the package
func TestMain(m *testing.M) {

	// creating connection with the database
	dbConnection, error := sql.Open("postgres", "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable")

	if error != nil {
		log.Fatalf("❌ error connecting to database")

		log.Fatal(error.Error( )) }

	accountOperationQueries= database.New(dbConnection)

	m.Run( )
}