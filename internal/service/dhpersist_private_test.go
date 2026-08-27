package service

import (
	"context"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"testing"
)

func TestLicenseChangeRollsBackWhenHistoryUnavailable(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, err := s.CreateProject(context.Background(), u, "Rights", "", 100)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), u, p.ID, "Archive", "en", "CC-BY", domain.Public, 2)
	if err != nil {
		t.Fatal(err)
	}
	_, err = s.DB.Exec("CREATE TRIGGER reject_license_history BEFORE INSERT ON license_changes BEGIN SELECT RAISE(ABORT, 'history unavailable'); END")
	if err != nil {
		t.Fatal(err)
	}
	if err = s.ChangeCorpusLicense(context.Background(), u, c.ID, "CC0", "committee decision"); err == nil {
		t.Fatal("expected history failure")
	}
	got, err := s.Corpora.Get(context.Background(), c.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.License != "CC-BY" {
		t.Fatalf("license change leaked without history: %s", got.License)
	}
}
