package service

import (
	"context"
	"errors"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"testing"
)

// newDraftAnalysis sets up a lead, project, corpus and a draft analysis so that
// publish flows have something to operate on. It returns the analysis id.
func newDraftAnalysis(t *testing.T, s *Service) (domain.User, int64) {
	t.Helper()
	u := lead(t, s)
	p, e := s.CreateProject(context.Background(), u, "P", "", 1000)
	if e != nil {
		t.Fatal(e)
	}
	c, e := s.AddCorpus(context.Background(), u, p.ID, "C", "en", "CC", domain.Public, 10)
	if e != nil {
		t.Fatal(e)
	}
	a, e := s.RecordAnalysis(context.Background(), u, AnalysisInput{ProjectID: p.ID, CorpusID: c.ID, Kind: "topic", Source: "manual", AIUse: "none"})
	if e != nil {
		t.Fatal(e)
	}
	return u, a.ID
}

func TestPublishAnalysisWithAuditAtomic(t *testing.T) {
	s := setup(t)
	_, id := newDraftAnalysis(t, s)

	// Force foreign keys on for the connection pool so the audit INSERT below
	// can fail on its actor_id foreign key, simulating an ethics-committee
	// rejection of the audit write. Because the status change and audit write
	// now share one transaction, the publication must roll back and leave the
	// analysis in draft.
	if _, e := s.DB.ExecContext(context.Background(), "PRAGMA foreign_keys=ON"); e != nil {
		t.Fatal(e)
	}

	// Publish under a non-existent actor so the audit_events.actor_id FK insert
	// fails, proving the publication rolls back atomically.
	err := s.PublishAnalysisWithAudit(context.Background(), domain.User{ID: 999999, Role: domain.RoleLead}, id, "result", "req")
	if err == nil {
		t.Fatal("expected audit write to fail and roll back publication")
	}

	got, e := s.Analyses.Get(context.Background(), id)
	if e != nil {
		t.Fatal(e)
	}
	if got.Status != "draft" {
		t.Fatalf("expected status to remain draft after rolled-back publish, got %q", got.Status)
	}
	if got.Result != "" {
		t.Fatalf("expected result to be cleared after rollback, got %q", got.Result)
	}
	// No audit event should have been committed.
	evts, e := s.Audit.List(context.Background(), "analysis", id)
	if e != nil {
		t.Fatal(e)
	}
	if len(evts) != 0 {
		t.Fatalf("expected no committed audit events, got %d", len(evts))
	}
}

func TestPublishAnalysisWithAuditSuccess(t *testing.T) {
	s := setup(t)
	u, id := newDraftAnalysis(t, s)

	if e := s.PublishAnalysisWithAudit(context.Background(), u, id, "result", "req"); e != nil {
		t.Fatalf("publish: %v", e)
	}
	got, e := s.Analyses.Get(context.Background(), id)
	if e != nil {
		t.Fatal(e)
	}
	if got.Status != "published" || got.Result != "result" {
		t.Fatalf("unexpected analysis state: %+v", got)
	}
	evts, e := s.Audit.List(context.Background(), "analysis", id)
	if e != nil {
		t.Fatal(e)
	}
	if len(evts) != 1 || evts[0].Action != "analysis_published" {
		t.Fatalf("expected one analysis_published audit event, got %+v", evts)
	}
}

func TestPublishAnalysisWithAuditConflict(t *testing.T) {
	s := setup(t)
	u, id := newDraftAnalysis(t, s)
	// Publish once to leave the draft state.
	if e := s.PublishAnalysisWithAudit(context.Background(), u, id, "result", "req"); e != nil {
		t.Fatal(e)
	}
	// Second publish from the now-published analysis must conflict.
	if e := s.PublishAnalysisWithAudit(context.Background(), u, id, "result2", "req2"); !errors.Is(e, domain.ErrConflict) {
		t.Fatalf("expected conflict, got %v", e)
	}
}
