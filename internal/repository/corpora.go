package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

type Corpora struct{ DB *sql.DB }

func (r Corpora) Create(ctx context.Context, c domain.Corpus) (int64, error) {
	res, e := r.DB.ExecContext(ctx, "INSERT INTO corpora(project_id,title,language,license,sensitivity,status,bytes,version,created_at,updated_at) VALUES(?,?,?,?,?,?,?,1,?,?)", c.ProjectID, c.Title, c.Language, c.License, c.Sensitivity, c.Status, c.Bytes, c.CreatedAt.Format(time.RFC3339Nano), c.UpdatedAt.Format(time.RFC3339Nano))
	if e != nil {
		return 0, e
	}
	id, err := res.LastInsertId()
	return id, err
}
func (r Corpora) Get(ctx context.Context, id int64) (domain.Corpus, error) {
	var c domain.Corpus
	var a, b string
	e := r.DB.QueryRowContext(ctx, "SELECT id,project_id,title,language,license,sensitivity,status,bytes,version,created_at,updated_at FROM corpora WHERE id=?", id).Scan(&c.ID, &c.ProjectID, &c.Title, &c.Language, &c.License, &c.Sensitivity, &c.Status, &c.Bytes, &c.Version, &a, &b)
	c.CreatedAt, _ = time.Parse(time.RFC3339Nano, a)
	c.UpdatedAt, _ = time.Parse(time.RFC3339Nano, b)
	return c, e
}
func (r Corpora) Transition(ctx context.Context, id int64, from, to domain.CorpusStatus, version int64) error {
	if _, e := r.DB.ExecContext(ctx, "UPDATE corpora SET status=?,version=version+1,updated_at=? WHERE id=? AND status=? AND version=?", to, time.Now().UTC().Format(time.RFC3339Nano), id, from, version); e != nil {
		return e
	}
	var n int
	_ = r.DB.QueryRowContext(ctx, "SELECT changes()").Scan(&n)
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}
func (r Corpora) UpdateLicense(ctx context.Context, id int64, license string, version int64) error {
	res, e := r.DB.ExecContext(ctx, "UPDATE corpora SET license=?,version=version+1,updated_at=? WHERE id=? AND version=?", license, time.Now().UTC().Format(time.RFC3339Nano), id, version)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}
