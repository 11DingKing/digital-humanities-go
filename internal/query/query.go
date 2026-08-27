package query

import (
	"context"
	"database/sql"
)

type Summary struct {
	ProjectID                             int64
	CorpusCount, ReleasedCount, TaskCount int
}

func ProjectSummary(ctx context.Context, db *sql.DB, id int64) (Summary, error) {
	var s Summary
	s.ProjectID = id
	e := db.QueryRowContext(ctx, `SELECT COUNT(*),COALESCE(SUM(CASE WHEN status='released' THEN 1 ELSE 0 END),0) FROM corpora WHERE project_id=?`, id).Scan(&s.CorpusCount, &s.ReleasedCount)
	if e != nil {
		return s, e
	}
	e = db.QueryRowContext(ctx, `SELECT COUNT(*) FROM annotation_tasks t JOIN batches b ON b.id=t.batch_id JOIN corpora c ON c.id=b.corpus_id WHERE c.project_id=?`, id).Scan(&s.TaskCount)
	return s, e
}
