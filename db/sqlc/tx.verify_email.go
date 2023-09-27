package db

import (
	"context"
	"database/sql"
)

// TransferTxParams contains the input paramters of the transfer transaction
type VerifyEmailTxParams struct {
		EmailId int64
		SecretCode string
}

// TransferTxResult is the result of the transfer transaction
type VerifyEmailTxResult struct {
	User User
	VerifyEmail VerifyEmail
}

// TransferTx performs a money transfer from one account to the other
// It creates a transfer record, add account entries, and update accounts' balance within a single database transaction
func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID: arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}

		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Username: result.VerifyEmail.Username,
			IsEmailVerified: sql.NullBool{
				Bool: true,
				Valid: true,
			},
		})
		return err
	})

	return result, err
}
