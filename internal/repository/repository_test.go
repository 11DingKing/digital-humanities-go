package repository

import (
	"context"
	"errors"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/testutil"
	"testing"
	"time"
)

func TestRepositories(t *testing.T) {
	db := testutil.DB(t)
	ctx := context.Background()
	u := Users{db}
	id, e := u.Create(ctx, domain.User{Email: "r@x", Name: "R", Role: domain.RoleLead, PasswordHash: "h", CreatedAt: time.Now()})
	if e != nil {
		t.Fatal(e)
	}
	if _, e = u.ByID(ctx, id); e != nil {
		t.Fatal(e)
	}
	p := Projects{db}
	pid, e := p.Create(ctx, domain.Project{Name: "P", QuotaBytes: 100, Status: domain.Draft, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if e != nil {
		t.Fatal(e)
	}
	if _, e = p.Get(ctx, pid); e != nil {
		t.Fatal(e)
	}
	c := Corpora{db}
	cid, e := c.Create(ctx, domain.Corpus{ProjectID: pid, Title: "C", Language: "en", License: "CC", Sensitivity: domain.Public, Status: domain.Collected, Bytes: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if e != nil {
		t.Fatal(e)
	}
	if e = c.Transition(ctx, cid, domain.Collected, domain.Cleansed, 1); e != nil {
		t.Fatal(e)
	}
	b := Batches{db}
	bid, e := b.Create(ctx, domain.Batch{CorpusID: cid, Name: "B", Status: domain.BatchPending, Priority: 1, Concurrency: 1, CreatedAt: time.Now()})
	if e != nil {
		t.Fatal(e)
	}
	if _, e = b.Get(ctx, bid); e != nil {
		t.Fatal(e)
	}
	ts := Tasks{db}
	tid, e := ts.Create(ctx, domain.AnnotationTask{BatchID: bid, AssigneeID: id, Segment: "s", Status: domain.TaskQueued, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if e != nil {
		t.Fatal(e)
	}
	if e = ts.Claim(ctx, tid, 1, id, time.Now().Add(time.Hour)); e != nil {
		t.Fatal(e)
	}
	if e = ts.Complete(ctx, tid, 2, true); e != nil {
		t.Fatal(e)
	}
}
func TestOptimisticProjectConflict(t *testing.T) {
	db := testutil.DB(t)
	p := Projects{db}
	id, e := p.Create(context.Background(), domain.Project{Name: "P", QuotaBytes: 1, Status: domain.Draft, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if e != nil {
		t.Fatal(e)
	}
	if e = p.BumpVersion(context.Background(), id, 1); e != nil {
		t.Fatal(e)
	}
	if e = p.BumpVersion(context.Background(), id, 1); e != domain.ErrConflict {
		t.Fatal(e)
	}
}
func TestTaskRequeue(t *testing.T) {
	db := testutil.DB(t)
	u := Users{db}
	uid, _ := u.Create(context.Background(), domain.User{Email: "q@x", Name: "Q", Role: domain.RoleLead, PasswordHash: "h", CreatedAt: time.Now()})
	p := Projects{db}
	pid, _ := p.Create(context.Background(), domain.Project{Name: "P", QuotaBytes: 10, Status: domain.Draft, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	c := Corpora{db}
	cid, _ := c.Create(context.Background(), domain.Corpus{ProjectID: pid, Title: "C", Language: "en", License: "CC", Sensitivity: domain.Public, Status: domain.Collected, Bytes: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	b := Batches{db}
	bid, _ := b.Create(context.Background(), domain.Batch{CorpusID: cid, Name: "B", Status: domain.BatchPending, Priority: 1, Concurrency: 1, CreatedAt: time.Now()})
	ts := Tasks{db}
	tid, _ := ts.Create(context.Background(), domain.AnnotationTask{BatchID: bid, AssigneeID: uid, Segment: "s", Status: domain.TaskQueued, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	ts.Claim(context.Background(), tid, 1, uid, time.Now().Add(-time.Hour))
	if e := ts.RequeueExpired(context.Background(), time.Now()); e != nil {
		t.Fatal(e)
	}
}
func TestTaskReassignProtectsRunning(t *testing.T) {
	db := testutil.DB(t)
	u := Users{db}
	owner, _ := u.Create(context.Background(), domain.User{Email: "owner@x", Name: "Owner", Role: domain.RoleAnnotator, PasswordHash: "h", CreatedAt: time.Now()})
	other, _ := u.Create(context.Background(), domain.User{Email: "other@x", Name: "Other", Role: domain.RoleAnnotator, PasswordHash: "h", CreatedAt: time.Now()})
	p := Projects{db}
	pid, _ := p.Create(context.Background(), domain.Project{Name: "P", QuotaBytes: 10, Status: domain.Draft, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	c := Corpora{db}
	cid, _ := c.Create(context.Background(), domain.Corpus{ProjectID: pid, Title: "C", Language: "en", License: "CC", Sensitivity: domain.Public, Status: domain.Collected, Bytes: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	b := Batches{db}
	bid, _ := b.Create(context.Background(), domain.Batch{CorpusID: cid, Name: "B", Status: domain.BatchPending, Priority: 1, Concurrency: 1, CreatedAt: time.Now()})
	ts := Tasks{db}
	tid, _ := ts.Create(context.Background(), domain.AnnotationTask{BatchID: bid, AssigneeID: owner, Segment: "s", Status: domain.TaskQueued, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	// Reassigning a queued task to a new nominal owner is allowed.
	if e := ts.Reassign(context.Background(), tid, other); e != nil {
		t.Fatalf("reassign queued: %v", e)
	}
	// The annotator claims the shard, transitioning it to running.
	if e := ts.Claim(context.Background(), tid, 2, owner, time.Now().Add(time.Hour)); e != nil {
		t.Fatalf("claim: %v", e)
	}
	// Project lead attempts to silently reassign the in-flight task to someone
	// else. This must be rejected to protect the running task's ownership.
	if e := ts.Reassign(context.Background(), tid, other); !errors.Is(e, domain.ErrConflict) {
		t.Fatalf("expected conflict reassigning running task, got %v", e)
	}
	// The original owner must remain the recorded assignee.
	var got int64
	if e := db.QueryRowContext(context.Background(), "SELECT assignee_id FROM annotation_tasks WHERE id=?", tid).Scan(&got); e != nil {
		t.Fatal(e)
	}
	if got != owner {
		t.Fatalf("assignee rewritten to %d, expected owner %d", got, owner)
	}
}
