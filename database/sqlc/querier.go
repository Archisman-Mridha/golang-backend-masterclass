// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package database

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccountDetails(ctx context.Context, id int64) (Account, error)
	GetAccountDetails_WithLock(ctx context.Context, id int64) (Account, error)
	GetEntryDetails(ctx context.Context, arg GetEntryDetailsParams) (Entry, error)
	GetMoneyTransferDetails(ctx context.Context, id int64) (Transfer, error)
	GetUserDetails(ctx context.Context, username string) (User, error)
	//TODO: understand why "no key" was added
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListMoneyTransfers(ctx context.Context, arg ListMoneyTransfersParams) ([]Transfer, error)
	TransferMoney(ctx context.Context, arg TransferMoneyParams) (Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
	UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (Account, error)
	UpdateUserDetails(ctx context.Context, arg UpdateUserDetailsParams) (User, error)
}

var _ Querier = (*Queries)(nil)
