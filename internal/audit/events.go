package audit

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

func Append(ctx context.Context, tx *sql.Tx, actor int64, typ string, id int64, action, outcome, request, details string, now time.Time) error {
	_, e := tx.ExecContext(ctx, "INSERT INTO audit_events(actor_id,entity_type,entity_id,action,outcome,request_id,details,created_at) VALUES(?,?,?,?,?,?,?,?)", actor, typ, id, action, outcome, request, details, now.Format(time.RFC3339Nano))
	return e
}
func Event(typ string, id int64, action string) domain.AuditEvent {
	return domain.AuditEvent{EntityType: typ, EntityID: id, Action: action, Outcome: "ok"}
}
