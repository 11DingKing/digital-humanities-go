package service

import (
	"context"
	"testing"
)

func TestCancelledCorpusWriteDoesNotCommit(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "Context", "", 100)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err = s.AddCorpusWithContext(ctx, u, p.ID, "cancelled"); err == nil {
		t.Fatal("expected cancellation")
	}
	var count int
	if err = s.DB.QueryRow("SELECT COUNT(*) FROM corpora WHERE project_id=?", p.ID).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("cancelled write committed %d corpus rows", count)
	}
}
