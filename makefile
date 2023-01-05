create-migration-files:
	migrate create -ext sql -dir ./database/migrations -seq $(seq)

run-up-migrations:
	migrate -path database/migrations \
		-database "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable" \
		-verbose up

run-down-migrations:
	migrate -path database/migrations \
		-database "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable" \
		-verbose down

protoc-generate:
	protoc --go_out=./generated/proto --go-grpc_out=./generated/proto \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative \
		--proto_path=./proto ./proto/*.proto