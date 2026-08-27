package audit

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

type Logger struct{ DB *sql.DB }

func (l Logger) Record(ctx context.Context, e domain.AuditEvent) error {
	_, err := l.DB.ExecContext(ctx, "INSERT INTO audit_events(actor_id,entity_type,entity_id,action,outcome,request_id,details,created_at) VALUES(?,?,?,?,?,?,?,?)", e.ActorID, e.EntityType, e.EntityID, e.Action, e.Outcome, e.RequestID, e.Details, e.CreatedAt.Format(time.RFC3339Nano))
	return err
}
func (l Logger) List(ctx context.Context, typ string, id int64) ([]domain.AuditEvent, error) {
	rows, e := l.DB.QueryContext(ctx, "SELECT id,actor_id,entity_type,entity_id,action,outcome,request_id,details,created_at FROM audit_events WHERE entity_type=? AND entity_id=? ORDER BY id", typ, id)
	if e != nil {
		return nil, e
	}
	defer rows.Close()
	var out []domain.AuditEvent
	for rows.Next() {
		var a domain.AuditEvent
		var t string
		if e := rows.Scan(&a.ID, &a.ActorID, &a.EntityType, &a.EntityID, &a.Action, &a.Outcome, &a.RequestID, &a.Details, &t); e != nil {
			return nil, e
		}
		a.CreatedAt, _ = time.Parse(time.RFC3339Nano, t)
		out = append(out, a)
	}
	return out, rows.Err()
}
