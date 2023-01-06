package database_tests

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"

	"simple_bank/configuration"
	database "simple_bank/database/sqlc"
)

var (
	accountOperationQueries *database.Queries
	DBConnection *sql.DB
)

//* main entrypoint of unit testing inside the package
func TestMain(m *testing.M) {
	var error error

	appConfiguration, error := configuration.LoadAppConfiguration("../../envs")
	if error != nil {
		log.Fatal("❌ error loading application configurations")

		log.Fatal(error.Error( )) }

	// creating connection with the database
	DBConnection, error= sql.Open("postgres", appConfiguration.POSTGRES_DB_URL)

	if error != nil {
		log.Fatalf("❌ error connecting to database")

		log.Fatal(error.Error( )) }

	accountOperationQueries= database.New(DBConnection)

	m.Run( )
}