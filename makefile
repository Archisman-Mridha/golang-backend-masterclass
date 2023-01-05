create-migration-files:
	migrate create -ext sql -dir ./database/migrations -seq initialize_schema

run-migrations:
	migrate -path database/migrations \
		-database "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable" \
		-verbose up