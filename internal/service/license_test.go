package service

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/repository"
	"github.com/11DingKing/digital-humanities-go/internal/storage"
	"testing"
)

// TestChangeCorpusLicenseRecordsHistory verifies the happy path: the
// corpus license and its history change land together so the
// authorization trail stays in sync with the current license.
func TestChangeCorpusLicenseRecordsHistory(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	c, _ := s.AddCorpus(context.Background(), u, p.ID, "A", "en", "CC", domain.Public, 10)

	if e := s.ChangeCorpusLicense(context.Background(), u, c.ID, "GPL", "license swap"); e != nil {
		t.Fatal(e)
	}
	got, _ := s.Corpora.Get(context.Background(), c.ID)
	if got.License != "GPL" {
		t.Fatalf("license = %q, want GPL", got.License)
	}
	if got.Version != c.Version+1 {
		t.Fatalf("version = %d, want %d", got.Version, c.Version+1)
	}
	var prev, cur string
	if e := s.DB.QueryRow("SELECT previous,current FROM license_changes WHERE corpus_id=?", c.ID).Scan(&prev, &cur); e != nil {
		t.Fatal(e)
	}
	if prev != "CC" || cur != "GPL" {
		t.Fatalf("history prev=%q cur=%q", prev, cur)
	}
}

// TestChangeCorpusLicenseHistoryFailureRollsBack verifies the fix for
// the reported defect: if the license history record cannot be
// written, the corpus license must revert to its previous value rather
// than drifting to the new value without a recoverable trail.
func TestChangeCorpusLicenseHistoryFailureRollsBack(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	c, _ := s.AddCorpus(context.Background(), u, p.ID, "A", "en", "CC", domain.Public, 10)

	// Drop the history table so the INSERT inside the transaction fails,
	// forcing the whole change (license update + history) to roll back.
	if _, e := s.DB.Exec("DROP TABLE license_changes"); e != nil {
		t.Fatal(e)
	}

	err := s.ChangeCorpusLicense(context.Background(), u, c.ID, "GPL", "license swap")
	if err == nil {
		t.Fatal("expected history write failure, got nil")
	}
	got, _ := s.Corpora.Get(context.Background(), c.ID)
	if got.License != "CC" {
		t.Fatalf("license drifted to %q on history failure; want CC", got.License)
	}
	if got.Version != c.Version {
		t.Fatalf("version = %d, want unchanged %d", got.Version, c.Version)
	}
}

// TestRecordLicenseChangeAcceptsTx is a compile-time guard in test form:
// it ensures the history recorder can run inside a transaction, which is
// what keeps the license update and history atomic.
func TestRecordLicenseChangeAcceptsTx(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	c, _ := s.AddCorpus(context.Background(), u, p.ID, "A", "en", "CC", domain.Public, 10)

	err := storage.Tx(context.Background(), s.DB, func(tx *sql.Tx) error {
		return repository.RecordLicenseChange(context.Background(), tx, c.ID, u.ID, "CC", "GPL", "reason", s.Clock.Now())
	})
	if err != nil {
		t.Fatal(err)
	}
}
