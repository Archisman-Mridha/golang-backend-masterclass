// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package database

import (
	"database/sql"
)

type Account struct {
	ID        int64        `json:"id"`
	Owner     string       `json:"owner"`
	Balance   int64        `json:"balance"`
	Currency  string       `json:"currency"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// it can be positive or negative
	Amount    int64        `json:"amount"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// it should always be positive
	Amount    int64        `json:"amount"`
	CreatedAt sql.NullTime `json:"created_at"`
}
