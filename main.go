package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"simple_bank/configuration"
	"simple_bank/server"
)

func main( ) {

	appConfiguration, error := configuration.LoadAppConfiguration("./envs")
	if error != nil {
		log.Fatal("❌ error loading application configurations")

		log.Fatal(error.Error( )) }

	_, error= sql.Open("postgres", appConfiguration.POSTGRES_DB_URL)
	if error != nil {
		log.Fatalf("❌ error connecting to database")

		log.Fatal(error.Error( )) }

	server.StartGRPCServer(appConfiguration.GRPC_SERVER_ADDRESS)
}