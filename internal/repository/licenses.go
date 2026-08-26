package repository

import (
	"context"
	"database/sql"
	"time"
)

func RecordLicenseChange(ctx context.Context, db *sql.DB, corpusID, actorID int64, previous, current, reason string, now time.Time) error {
	_, err := db.ExecContext(ctx, "INSERT INTO license_changes(corpus_id,actor_id,previous,current,reason,created_at) VALUES(?,?,?,?,?,?)", corpusID, actorID, previous, current, reason, now.Format(time.RFC3339Nano))
	return err
}
