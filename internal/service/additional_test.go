package service

import (
	"context"
	"errors"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"testing"
)

func TestInvalidInputs(t *testing.T) {
	s := setup(t)
	cases := []struct {
		name string
		fn   func() error
	}{
		{"empty email", func() error { _, e := s.Register(context.Background(), "", "n", "p", domain.RoleLead); return e }},
		{"empty password", func() error { _, e := s.Register(context.Background(), "x@x", "n", "", domain.RoleLead); return e }},
		{"bad quota", func() error {
			ss := setup(t)
			u := lead(t, ss)
			_, e := ss.CreateProject(context.Background(), u, "p", "", 0)
			return e
		}},
		{"bad priority", func() error {
			ss := setup(t)
			u := lead(t, ss)
			_, e := ss.CreateBatch(context.Background(), u, 1, "b", 11, 1)
			return e
		}},
		{"bad concurrency", func() error {
			ss := setup(t)
			u := lead(t, ss)
			_, e := ss.CreateBatch(context.Background(), u, 1, "b", 1, 0)
			return e
		}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if !errors.Is(c.fn(), domain.ErrInvalid) {
				t.Fatalf("expected invalid")
			}
		})
	}
}
func TestRoleMatrix(t *testing.T) {
	roles := []domain.Role{domain.RoleLead, domain.RoleCurator, domain.RoleAnnotator, domain.RoleReviewer}
	for _, r := range roles {
		t.Run(string(r), func(t *testing.T) {
			s := setup(t)
			u, e := s.Register(context.Background(), string(r)+"@x", "N", "p", r)
			if e != nil {
				t.Fatal(e)
			}
			if u.Role != r {
				t.Fatal("role mismatch")
			}
		})
	}
}
func TestProjectNamesPersist(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	for i := 1; i <= 20; i++ {
		p, e := s.CreateProject(context.Background(), u, "project", string(rune(i)), 100)
		if e != nil || p.ID == 0 {
			t.Fatal(e)
		}
	}
}
func TestManyCorpusLanguages(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 1000)
	langs := []string{"en", "fr", "de", "es", "zh", "ja", "ko", "ar", "ru", "it"}
	for _, lang := range langs {
		c, e := s.AddCorpus(context.Background(), u, p.ID, "title", lang, "CC", domain.Public, 10)
		if e != nil || c.Language != lang {
			t.Fatal(e)
		}
	}
}
func TestForbiddenCorpusRole(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	for _, r := range []domain.Role{domain.RoleAnnotator, domain.RoleReviewer} {
		x, _ := s.Register(context.Background(), string(r)+"@x", "N", "p", r)
		if _, e := s.AddCorpus(context.Background(), x, p.ID, "C", "en", "CC", domain.Public, 1); !errors.Is(e, domain.ErrForbidden) {
			t.Fatalf("role %s", r)
		}
	}
}
func TestAdvanceRequiresOrder(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	c, _ := s.AddCorpus(context.Background(), u, p.ID, "C", "en", "CC", domain.Public, 1)
	for _, next := range []domain.CorpusStatus{domain.Licensed, domain.Released, domain.Withdrawn} {
		if e := s.AdvanceCorpus(context.Background(), u, c.ID, next, "r"); e == nil {
			t.Fatalf("skipped to %s", next)
		}
	}
}
func TestBatchStartRole(t *testing.T) {
	s := setup(t)
	for _, r := range []domain.Role{domain.RoleAnnotator, domain.RoleReviewer} {
		u, _ := s.Register(context.Background(), string(r)+"@x", "N", "p", r)
		if e := s.StartBatch(context.Background(), u, 1, "r"); !errors.Is(e, domain.ErrForbidden) {
			t.Fatal(e)
		}
	}
}
func TestBatchStartWithdrawnCorpus(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	c, _ := s.AddCorpus(context.Background(), u, p.ID, "C", "en", "CC", domain.Public, 1)
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "r"); e != nil {
		t.Fatal(e)
	}
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "r"); e != nil {
		t.Fatal(e)
	}
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Withdrawn, "r"); e != nil {
		t.Fatal(e)
	}
	b, _ := s.CreateBatch(context.Background(), u, c.ID, "b", 1, 1)
	if e := s.StartBatch(context.Background(), u, b.ID, "r"); !errors.Is(e, domain.ErrConflict) {
		t.Fatalf("expected conflict on withdrawn corpus, got %v", e)
	}
	got, _ := s.Batches.Get(context.Background(), b.ID)
	if got.Status != domain.BatchPending {
		t.Fatalf("batch regressed to %s", got.Status)
	}
}
func TestBatchStartReleasedCorpus(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	c, _ := s.AddCorpus(context.Background(), u, p.ID, "C", "en", "CC", domain.Public, 1)
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "r"); e != nil {
		t.Fatal(e)
	}
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "r"); e != nil {
		t.Fatal(e)
	}
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Released, "r"); e != nil {
		t.Fatal(e)
	}
	b, _ := s.CreateBatch(context.Background(), u, c.ID, "b", 1, 1)
	if e := s.StartBatch(context.Background(), u, b.ID, "r"); !errors.Is(e, domain.ErrConflict) {
		t.Fatalf("expected conflict on released corpus, got %v", e)
	}
}
func TestLogoutIdempotent(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	tok, _, _ := s.Login(context.Background(), u.Email, "secret")
	if e := s.Logout(context.Background(), tok); e != nil {
		t.Fatal(e)
	}
	if e := s.Logout(context.Background(), tok); e != nil {
		t.Fatal(e)
	}
}
func TestAuthenticateUnknown(t *testing.T) {
	s := setup(t)
	if _, e := s.Authenticate(context.Background(), "unknown"); e == nil {
		t.Fatal("unknown accepted")
	}
}
func TestRegisterClock(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	if u.CreatedAt.IsZero() {
		t.Fatal("timestamp")
	}
}
func TestQuotaExactBoundary(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 10)
	if _, e := s.AddCorpus(context.Background(), u, p.ID, "C", "en", "CC", domain.Public, 10); e != nil {
		t.Fatal(e)
	}
}
func TestQuotaWithdrawnIgnored(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 10)
	c, _ := s.AddCorpus(context.Background(), u, p.ID, "C", "en", "CC", domain.Public, 10)
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "r"); e != nil {
		t.Fatal(e)
	}
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Licensed, "r"); e != nil {
		t.Fatal(e)
	}
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Withdrawn, "r"); e != nil {
		t.Fatal(e)
	}
	if _, e := s.AddCorpus(context.Background(), u, p.ID, "D", "en", "CC", domain.Public, 10); e != nil {
		t.Fatal(e)
	}
}
func TestServiceMethodsNonNil(t *testing.T) {
	s := setup(t)
	if s.Users.DB == nil || s.Sessions.DB == nil || s.Projects.DB == nil || s.Corpora.DB == nil || s.Batches.DB == nil || s.Tasks.DB == nil || s.Audit.DB == nil {
		t.Fatal("repositories")
	}
}
