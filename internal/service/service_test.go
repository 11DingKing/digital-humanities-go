package service

import (
	"context"
	"errors"
	"github.com/11DingKing/digital-humanities-go/internal/clock"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/testutil"
	"testing"
	"time"
)

func setup(t *testing.T) *Service {
	s := New(testutil.DB(t))
	s.Clock = clock.Fixed{T: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)}
	return s
}
func lead(t *testing.T, s *Service) domain.User {
	u, e := s.Register(context.Background(), "lead@example.com", "Lead", "secret", domain.RoleLead)
	if e != nil {
		t.Fatal(e)
	}
	return u
}
func TestRegisterLoginLogout(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	tok, _, e := s.Login(context.Background(), u.Email, "secret")
	if e != nil || tok == "" {
		t.Fatal(e)
	}
	if _, e = s.Authenticate(context.Background(), tok); e != nil {
		t.Fatal(e)
	}
	if e = s.Logout(context.Background(), tok); e != nil {
		t.Fatal(e)
	}
	if _, e = s.Authenticate(context.Background(), tok); !errors.Is(e, domain.ErrForbidden) {
		t.Fatalf("expected revoked, got %v", e)
	}
}
func TestExpiredSession(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	tok, _, _ := s.Login(context.Background(), u.Email, "secret")
	s.Clock = clock.Fixed{T: time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)}
	if _, e := s.Authenticate(context.Background(), tok); !errors.Is(e, domain.ErrForbidden) {
		t.Fatal("expired token accepted")
	}
}
func TestRoleCannotCreateProject(t *testing.T) {
	s := setup(t)
	u, e := s.Register(context.Background(), "a@x", "A", "p", domain.RoleAnnotator)
	if e != nil {
		t.Fatal(e)
	}
	if _, e = s.CreateProject(context.Background(), u, "x", "", 100); !errors.Is(e, domain.ErrForbidden) {
		t.Fatal(e)
	}
}
func TestCreateProject(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, e := s.CreateProject(context.Background(), u, "Corpus study", "desc", 1000)
	if e != nil || p.ID == 0 || p.Status != domain.Draft {
		t.Fatalf("%+v %v", p, e)
	}
}
func TestCorpusQuota(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	if _, e := s.AddCorpus(context.Background(), u, p.ID, "A", "en", "CC", domain.Public, 80); e != nil {
		t.Fatal(e)
	}
	if _, e := s.AddCorpus(context.Background(), u, p.ID, "B", "fr", "CC", domain.Public, 30); !errors.Is(e, domain.ErrConflict) {
		t.Fatalf("quota %v", e)
	}
}
func TestCorpusLifecycleService(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	c, _ := s.AddCorpus(context.Background(), u, p.ID, "A", "en", "CC", domain.Public, 20)
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Cleansed, "r1"); e != nil {
		t.Fatal(e)
	}
	if e := s.AdvanceCorpus(context.Background(), u, c.ID, domain.Released, "r2"); e == nil {
		t.Fatal("skipped state")
	}
}
func TestBatchValidation(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	if _, e := s.CreateBatch(context.Background(), u, 1, "b", 0, 1); !errors.Is(e, domain.ErrInvalid) {
		t.Fatal(e)
	}
}
func TestContextCancellation(t *testing.T) {
	s := setup(t)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, e := s.Register(ctx, "x@y", "X", "p", domain.RoleLead); e == nil {
		t.Fatal("cancel ignored")
	}
}
func TestLicenseUpdateVersionConflict(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	p, _ := s.CreateProject(context.Background(), u, "P", "", 100)
	c, _ := s.AddCorpus(context.Background(), u, p.ID, "A", "en", "CC", domain.Public, 10)
	if e := s.Corpora.UpdateLicense(context.Background(), c.ID, "GPL", 99); !errors.Is(e, domain.ErrConflict) {
		t.Fatal(e)
	}
}
func TestStartBatchAudit(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	if e := s.StartBatch(context.Background(), u, 99, "r"); e == nil {
		t.Fatal("missing batch accepted")
	}
}
func TestRegisterDuplicate(t *testing.T) {
	s := setup(t)
	lead(t, s)
	if _, e := s.Register(context.Background(), "lead@example.com", "Other", "x", domain.RoleLead); e == nil {
		t.Fatal("duplicate")
	}
}
func TestWrongPassword(t *testing.T) {
	s := setup(t)
	u := lead(t, s)
	if _, _, e := s.Login(context.Background(), u.Email, "bad"); !errors.Is(e, domain.ErrForbidden) {
		t.Fatal(e)
	}
}
func TestCuratorMayAddCorpus(t *testing.T) {
	s := setup(t)
	l := lead(t, s)
	p, _ := s.CreateProject(context.Background(), l, "P", "", 100)
	c, _ := s.Register(context.Background(), "c@x", "C", "p", domain.RoleCurator)
	if _, e := s.AddCorpus(context.Background(), c, p.ID, "A", "en", "CC", domain.Public, 10); e != nil {
		t.Fatal(e)
	}
}
