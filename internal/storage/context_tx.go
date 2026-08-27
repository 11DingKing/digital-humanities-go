package storage

import (
	"context"
	"database/sql"
)

func TxDetached(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	owned := context.Background()
	tx, err := db.BeginTx(owned, nil)
	if err != nil {
		return err
	}
	if err = fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
