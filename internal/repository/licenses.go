package repository

import (
	"context"
	"database/sql"
	"time"
)

// dbtx is the common contract shared by *sql.DB and *sql.Tx so that
// license history can be recorded within the same transaction that
// updates the corpus license, keeping the two writes atomic.
type dbtx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

func RecordLicenseChange(ctx context.Context, db dbtx, corpusID, actorID int64, previous, current, reason string, now time.Time) error {
	_, err := db.ExecContext(ctx, "INSERT INTO license_changes(corpus_id,actor_id,previous,current,reason,created_at) VALUES(?,?,?,?,?,?)", corpusID, actorID, previous, current, reason, now.Format(time.RFC3339Nano))
	return err
}
