package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

type Batches struct{ DB *sql.DB }

func (r Batches) Create(ctx context.Context, b domain.Batch) (int64, error) {
	res, e := r.DB.ExecContext(ctx, "INSERT INTO batches(corpus_id,name,status,priority,concurrency,version,created_at) VALUES(?,?,?,?,?,1,?)", b.CorpusID, b.Name, b.Status, b.Priority, b.Concurrency, b.CreatedAt.Format(time.RFC3339Nano))
	if e != nil {
		return 0, e
	}
	id, err := res.LastInsertId()
	return id, err
}
func (r Batches) Get(ctx context.Context, id int64) (domain.Batch, error) {
	var b domain.Batch
	var s, fin, cr sql.NullString
	e := r.DB.QueryRowContext(ctx, "SELECT id,corpus_id,name,status,priority,concurrency,version,started_at,finished_at,created_at FROM batches WHERE id=?", id).Scan(&b.ID, &b.CorpusID, &b.Name, &b.Status, &b.Priority, &b.Concurrency, &b.Version, &s, &fin, &cr)
	if s.Valid {
		x, _ := time.Parse(time.RFC3339Nano, s.String)
		b.StartedAt = &x
	}
	if fin.Valid {
		x, _ := time.Parse(time.RFC3339Nano, fin.String)
		b.FinishedAt = &x
	}
	b.CreatedAt, _ = time.Parse(time.RFC3339Nano, cr.String)
	return b, e
}
func (r Batches) Transition(ctx context.Context, id int64, from, to domain.BatchStatus, version int64) error {
	now := time.Now().UTC().Format(time.RFC3339Nano)
	res, e := r.DB.ExecContext(ctx, "UPDATE batches SET status=?,version=version+1,started_at=CASE WHEN ?='running' AND started_at IS NULL THEN ? ELSE started_at END,finished_at=CASE WHEN ? IN ('done','failed','cancelled') THEN ? ELSE finished_at END WHERE id=? AND status=? AND version=?", to, to, now, to, now, id, from, version)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}
