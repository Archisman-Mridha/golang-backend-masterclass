package database_tests

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"

	database "simple_bank/database/sqlc"
)

var (
	accountOperationQueries *database.Queries
	DBConnection *sql.DB
)

//* main entrypoint of unit testing inside the package
func TestMain(m *testing.M) {
	var error error

	// creating connection with the database
	DBConnection, error= sql.Open("postgres", "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable")

	if error != nil {
		log.Fatalf("❌ error connecting to database")

		log.Fatal(error.Error( )) }

	accountOperationQueries= database.New(DBConnection)

	m.Run( )
}