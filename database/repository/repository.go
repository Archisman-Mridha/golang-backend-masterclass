package repository

import (
	"context"
	"database/sql"
	"fmt"

	database "simple_bank/database/sqlc"
)

//* provides all the database queries and transactions
type Repository struct {

	// the database connection is required for creating transactions
	databaseConnection *sql.DB

	*database.Queries
}

//* executes a database transaction
func(repository *Repository) executeDatabaseTransaction(
	ctx context.Context, executeTransactionOperations func(*database.Queries) error) error {

	// begin the database transaction
	transaction, error := repository.databaseConnection.BeginTx(ctx, nil)
	if error != nil { return error }

	queries := database.New(transaction)

	// executing the transaction operations
	transactionError := executeTransactionOperations(queries)

	// rollback transaction in case of any error
	if transactionError != nil {
		rollbackError := transaction.Rollback( )

		if rollbackError != nil {
			return fmt.Errorf("database transaction error : %v, rollback error: %v", transactionError, rollbackError) }

		return transactionError
	}

	// commiting the transaction
	return transaction.Commit( )
}

func CreateRepository(databaseConnection *sql.DB) *Repository {
	return &Repository{

		databaseConnection: databaseConnection,
		Queries: database.New(databaseConnection),
	}
}