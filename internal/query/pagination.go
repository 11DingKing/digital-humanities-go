package query

import (
	"context"
	"database/sql"
	"fmt"
)

type Page struct{ Limit, Offset int }

func NormalizePage(limit, offset int) Page {
	if limit < 1 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	return Page{limit, offset}
}
func CorpusPage(ctx context.Context, db *sql.DB, pid int64, limit, offset int) ([]map[string]any, error) {
	p := NormalizePage(limit, offset)
	rows, e := db.QueryContext(ctx, "SELECT id,title,language,status,bytes FROM corpora WHERE project_id=? ORDER BY id LIMIT ? OFFSET ?", pid, p.Limit, p.Offset)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	out := []map[string]any{}
	for rows.Next() {
		var id, bytes int64
		var title, lang, status string
		if e := rows.Scan(&id, &title, &lang, &status, &bytes); e != nil {
			return nil, e
		}
		out = append(out, map[string]any{"id": id, "title": title, "language": lang, "status": status, "bytes": bytes})
	}
	if e := rows.Err(); e != nil {
		return nil, fmt.Errorf("page: %w", e)
	}
	return out, nil
}
