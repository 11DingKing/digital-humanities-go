package service

import (
	"context"
	"testing"
)

func TestAddCorpusWithContextRespectsCancellation(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 1000)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if e := s.AddCorpusWithContext(ctx, u, p.ID, "should not persist"); e == nil {
		t.Fatalf("expected cancellation error, got nil")
	}

	var n int
	if e := s.DB.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM corpora WHERE project_id=?", p.ID).Scan(&n); e != nil {
		t.Fatal(e)
	}
	if n != 0 {
		t.Fatalf("expected 0 corpora after cancelled request, got %d", n)
	}
}
