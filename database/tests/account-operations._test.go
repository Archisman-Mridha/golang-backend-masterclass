package database_tests

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"

	database "simple_bank/database/sqlc"
)

//* new account should be created successfully
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

//* existing account details should be fetched successfully
func TestGetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	fetchedAccount, error := accountOperationQueries.GetAccountDetails(context.Background( ), createdAccount.ID)

	require.NoError(t, error)
	require.NotEmpty(t, fetchedAccount)

	require.Equal(t, createdAccount.ID, fetchedAccount.ID)
	require.Equal(t, createdAccount.Owner, fetchedAccount.Owner)
	require.Equal(t, createdAccount.Balance, fetchedAccount.Balance)
	require.Equal(t, createdAccount.Currency, fetchedAccount.Currency)

	require.WithinDuration(t, createdAccount.CreatedAt.Time, fetchedAccount.CreatedAt.Time, time.Second)
}

//* existing account should be updated successfully
func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	input := database.UpdateAccountParams{
		ID: createdAccount.ID,
		Balance: 150,
	}

	error := accountOperationQueries.UpdateAccount(context.Background( ), input)
	require.NoError(t, error)

	fetchedAccount, error := accountOperationQueries.GetAccountDetails(context.Background( ), createdAccount.ID)
	require.NoError(t, error)

	require.Equal(t, input.Balance, fetchedAccount.Balance)
}

//* existing account should be deleted successfully
func TestDeleteAccunt(t *testing.T) {
	createdAccount := createRandomAccount(t)

	error := accountOperationQueries.DeleteAccount(context.Background( ), createdAccount.ID)
	require.NoError(t, error)

	fetchedAccount, error := accountOperationQueries.GetAccountDetails(context.Background( ), createdAccount.ID)

	require.EqualError(t, error, sql.ErrNoRows.Error( ))
	require.Empty(t, fetchedAccount)
}

func TestListAccounts(t *testing.T) {

	for i := 0; i< 10; i++ {
		createRandomAccount(t)
	}

	input := database.ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	fetchedAccounts, error := accountOperationQueries.ListAccounts(context.Background( ), input)

	require.NoError(t, error)
	require.Len(t, fetchedAccounts, 5)
}

func createRandomAccount(t *testing.T) database.Account {

	input := database.CreateAccountParams{
		Owner: faker.Name( ),
		Balance: 100,
		Currency: faker.Currency( ),
	}

	createdAccount, error := accountOperationQueries.CreateAccount(context.Background( ), input)
	require.NoError(t, error)

	require.Equal(t, createdAccount.Owner, input.Owner)
	require.Equal(t, createdAccount.Balance, int64(input.Balance))
	require.Equal(t, createdAccount.Currency, input.Currency)

	require.NotZero(t, createdAccount.ID)
	require.NotZero(t, createdAccount.CreatedAt)

	return createdAccount
}