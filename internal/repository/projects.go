package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

type Projects struct{ DB *sql.DB }

func (r Projects) Create(ctx context.Context, p domain.Project) (int64, error) {
	res, e := r.DB.ExecContext(ctx, "INSERT INTO projects(name,description,quota_bytes,status,version,created_at,updated_at) VALUES(?,?,?,?,1,?,?)", p.Name, p.Description, p.QuotaBytes, p.Status, p.CreatedAt.Format(time.RFC3339Nano), p.UpdatedAt.Format(time.RFC3339Nano))
	if e != nil {
		return 0, e
	}
	id, err := res.LastInsertId()
	return id, err
}
func (r Projects) Get(ctx context.Context, id int64) (domain.Project, error) {
	var p domain.Project
	var c, u string
	e := r.DB.QueryRowContext(ctx, "SELECT id,name,description,quota_bytes,status,version,created_at,updated_at FROM projects WHERE id=?", id).Scan(&p.ID, &p.Name, &p.Description, &p.QuotaBytes, &p.Status, &p.Version, &c, &u)
	p.CreatedAt, _ = time.Parse(time.RFC3339Nano, c)
	p.UpdatedAt, _ = time.Parse(time.RFC3339Nano, u)
	return p, e
}
func (r Projects) BumpVersion(ctx context.Context, id, version int64) error {
	res, e := r.DB.ExecContext(ctx, "UPDATE projects SET version=version+1,updated_at=? WHERE id=? AND version=?", time.Now().UTC().Format(time.RFC3339Nano), id, version)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}
