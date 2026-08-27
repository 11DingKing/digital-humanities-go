package service

import (
	"context"
	"testing"
)

func TestWithdrawnCorpusCannotStartBatch(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "State", "", 100)
	if err != nil { t.Fatal(err) }
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Corpus", "en", "CC", "public", 1)
	if err != nil { t.Fatal(err) }
	b, err := s.CreateBatch(context.Background(), u, c.ID, "batch", 1, 1)
	if err != nil { t.Fatal(err) }
	if _, err = s.DB.Exec("UPDATE corpora SET status='withdrawn' WHERE id=?", c.ID); err != nil { t.Fatal(err) }
	if err = s.StartBatch(context.Background(), u, b.ID, "req"); err == nil { t.Fatal("expected withdrawn corpus conflict") }
	var status string
	if err = s.DB.QueryRow("SELECT status FROM batches WHERE id=?", b.ID).Scan(&status); err != nil { t.Fatal(err) }
	if status != "pending" { t.Fatalf("withdrawn corpus batch started as %s", status) }
}
