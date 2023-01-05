package database_tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"simple_bank/database/repository"
)

func TestMoneyTransfer(t *testing.T) {
	createdRepository := repository.CreateRepository(DBConnection)

	senderAccountDetails := createRandomAccount(t)
	receiverAccountDetails := createRandomAccount(t)

	// creating channels to pass data between threads

	errors := make(chan error)
	transactionResults := make(chan repository.MoneyTransferTransactionOutput)

	concurrentTransactionCount := 5

	// run concurrent transactions
	for i := 0; i< concurrentTransactionCount; i++ {
		go func( ) {

			transactionResult, error := createdRepository.MoneyTransferTransaction(
				context.Background( ), repository.MoneyTransferTransactionParameters{

					SenderAccountID: senderAccountDetails.ID,
					ReceiverAccountID: receiverAccountDetails.ID,
					Amount: int64(10),
				},
			)

			errors <- error
			transactionResults <- transactionResult
		}( )
	}

	for i := 1; i <= concurrentTransactionCount; i++ {

		// receiving data into the main thread from the go routines
		error := <- errors
		transactionResult := <- transactionResults

		require.NoError(t, error)
		require.NotEmpty(t, transactionResult)

		// tests for the transaction record
		moneyTransferDetails := transactionResult.MoneyTransferDetails
		require.Equal(t, moneyTransferDetails.FromAccountID, senderAccountDetails.ID)
		require.Equal(t, moneyTransferDetails.ToAccountID, receiverAccountDetails.ID)
		require.Equal(t, moneyTransferDetails.Amount, int64(10))

		// TODO: tests for the entry record

		// test account balances

		require.Equal(t,
			senderAccountDetails.Balance - int64(10 * i), transactionResult.SenderAccountDetails.Balance)

		require.Equal(t,
			receiverAccountDetails.Balance + int64(10 * i), transactionResult.ReceiverAccountDetails.Balance)
	}
}