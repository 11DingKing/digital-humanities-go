package storage

import (
	"context"
	"database/sql"
)

// TxDetached opens a transaction that is not tied to the caller's context
// for its connection lifetime, but still respects cancellation: if ctx is
// already cancelled or becomes cancelled while the transaction is in flight,
// the work is rolled back and the context error is returned, so a cancelled
// request can never commit a new record.
func TxDetached(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	if err = fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err = ctx.Err(); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
