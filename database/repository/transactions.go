package repository

import (
	"context"

	database "simple_bank/database/sqlc"
)

//* transfer money from one account to another
func(repository *Repository) MoneyTransferTransaction(
	ctx context.Context, parameters MoneyTransferTransactionParameters) (MoneyTransferTransactionOutput, error) {

	var output MoneyTransferTransactionOutput

	error := repository.executeDatabaseTransaction(
		ctx, func(queries *database.Queries) error {

			//! making a record of the money transfer happening
			moneyTransferDetails, error := queries.TransferMoney(
				ctx, database.TransferMoneyParams{

					SenderAccountID: parameters.SenderAccountID,
					ReceiverAccountID: parameters.ReceiverAccountID,
					Amount: parameters.Amount,
				},
			)

			if error != nil { return error }
			output.MoneyTransferDetails= moneyTransferDetails

			//! create account entries

			// for sender
			entryDetailsForSender, error := queries.CreateEntry(
				ctx, database.CreateEntryParams{

					AccountID: parameters.SenderAccountID,
					Amount: -parameters.Amount,
				},
			)

			if error != nil { return error }
			output.EntryDetailsForSender= entryDetailsForSender

			// for receiver
			entryDetailsForReceiver, error := queries.CreateEntry(
				ctx, database.CreateEntryParams{

					AccountID: parameters.ReceiverAccountID,
					Amount: parameters.Amount,
				},
			)

			if error != nil { return error }
			output.EntryDetailsForReceiver= entryDetailsForReceiver

			//! update balance of both accounts

			// for sender
			error= queries.UpdateAccountBalance(
				ctx, database.UpdateAccountBalanceParams{

					ID: parameters.SenderAccountID,
					Amount: -parameters.Amount,
				},
			)

			if error != nil { return error }

			// for receiver
			error= queries.UpdateAccountBalance(
				ctx, database.UpdateAccountBalanceParams{

					ID: parameters.ReceiverAccountID,
					Amount: parameters.Amount,
				},
			)

			if error != nil { return error }

			return nil
		},
	)

	return output, error
}

type MoneyTransferTransactionParameters struct {

	SenderAccountID int64
	ReceiverAccountID int64
	Amount int64
}

type MoneyTransferTransactionOutput struct {

	MoneyTransferDetails database.Transfer
	SenderAccountDetails database.Account
	ReceiverAccountDetails database.Account
	EntryDetailsForSender database.Entry
	EntryDetailsForReceiver database.Entry
}