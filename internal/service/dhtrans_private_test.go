package service

import (
	"context"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"testing"
)

func TestPublishingRollsBackWhenAuditUnavailable(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "Edition", "", 100)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Texts", "en", "CC", domain.Public, 5)
	if err != nil {
		t.Fatal(err)
	}
	a, err := s.RecordAnalysis(context.Background(), u, AnalysisInput{ProjectID: p.ID, CorpusID: c.ID, Kind: "topic", Source: "human", AIUse: "none"})
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.DB.Exec("CREATE TRIGGER reject_analysis_audit BEFORE INSERT ON audit_events WHEN NEW.action='analysis_published' BEGIN SELECT RAISE(ABORT, 'audit unavailable'); END")
	if err != nil {
		t.Fatal(err)
	}
	publishErr := s.PublishAnalysisWithAudit(context.Background(), u, a.ID, "result", "req-audit")
	if publishErr == nil {
		t.Fatal("expected audit failure")
	}
	got, err := s.Analyses.Get(context.Background(), a.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != "draft" {
		t.Fatalf("analysis leaked published state: %s", got.Status)
	}
}
