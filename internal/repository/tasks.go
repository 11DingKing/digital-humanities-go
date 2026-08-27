package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

type Tasks struct{ DB *sql.DB }

func (r Tasks) Create(ctx context.Context, t domain.AnnotationTask) (int64, error) {
	res, e := r.DB.ExecContext(ctx, "INSERT INTO annotation_tasks(batch_id,assignee_id,segment,status,attempts,version,created_at,updated_at) VALUES(?,?,?,?,0,1,?,?)", t.BatchID, t.AssigneeID, t.Segment, t.Status, t.CreatedAt.Format(time.RFC3339Nano), t.UpdatedAt.Format(time.RFC3339Nano))
	if e != nil {
		return 0, e
	}
	id, err := res.LastInsertId()
	return id, err
}
func (r Tasks) Claim(ctx context.Context, id, ver, assignee int64, lease time.Time) error {
	res, e := r.DB.ExecContext(ctx, "UPDATE annotation_tasks SET status='running',assignee_id=?,lease_until=?,version=version+1,updated_at=? WHERE id=? AND version=? AND status='queued'", assignee, lease.Format(time.RFC3339Nano), time.Now().UTC().Format(time.RFC3339Nano), id, ver)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}

// Reassign changes the nominal assignee of a task. It is only permitted while
// the task is queued, i.e. no annotator has claimed it yet. A running task is
// actively owned (an annotator claimed the shard and holds the lease), so
// reassigning it would silently rewrite the owner and leave two attributions
// for one shard. Such an update is rejected with ErrConflict to protect the
// running task's ownership.
func (r Tasks) Reassign(ctx context.Context, id, assignee int64) error {
	res, err := r.DB.ExecContext(ctx, "UPDATE annotation_tasks SET assignee_id=?,version=version+1,updated_at=? WHERE id=? AND status='queued'", assignee, time.Now().UTC().Format(time.RFC3339Nano), id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}
func (r Tasks) Complete(ctx context.Context, id, ver int64, ok bool) error {
	st := "succeeded"
	if !ok {
		st = "failed"
	}
	res, e := r.DB.ExecContext(ctx, "UPDATE annotation_tasks SET status=?,attempts=attempts+1,version=version+1,updated_at=? WHERE id=? AND version=? AND status='running'", st, time.Now().UTC().Format(time.RFC3339Nano), id, ver)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return domain.ErrConflict
	}
	return nil
}
func (r Tasks) RequeueExpired(ctx context.Context, now time.Time) error {
	_, e := r.DB.ExecContext(ctx, "UPDATE annotation_tasks SET status='queued',lease_until=NULL,version=version+1,updated_at=? WHERE status='running' AND lease_until<?", now.Format(time.RFC3339Nano), now.Format(time.RFC3339Nano))
	return e
}
