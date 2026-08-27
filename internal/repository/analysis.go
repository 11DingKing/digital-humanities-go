package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

type Analyses struct{ DB *sql.DB }

func (r Analyses) Get(ctx context.Context, id int64) (domain.Analysis, error) {
	var a domain.Analysis
	var ts string
	e := r.DB.QueryRowContext(ctx, "SELECT id,project_id,corpus_id,author_id,kind,source,ai_use,status,COALESCE(result,''),created_at FROM analyses WHERE id=?", id).Scan(&a.ID, &a.ProjectID, &a.CorpusID, &a.AuthorID, &a.Kind, &a.Source, &a.AIUse, &a.Status, &a.Result, &ts)
	a.CreatedAt, _ = time.Parse(time.RFC3339Nano, ts)
	return a, e
}
func (r Analyses) ForProject(ctx context.Context, pid int64) ([]domain.Analysis, error) {
	rows, e := r.DB.QueryContext(ctx, "SELECT id,project_id,corpus_id,author_id,kind,source,ai_use,status,COALESCE(result,''),created_at FROM analyses WHERE project_id=? ORDER BY created_at DESC", pid)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []domain.Analysis{}
	for rows.Next() {
		var a domain.Analysis
		var ts string
		if e := rows.Scan(&a.ID, &a.ProjectID, &a.CorpusID, &a.AuthorID, &a.Kind, &a.Source, &a.AIUse, &a.Status, &a.Result, &ts); e != nil {
			return nil, e
		}
		a.CreatedAt, _ = time.Parse(time.RFC3339Nano, ts)
		out = append(out, a)
	}
	return out, rows.Err()
}
