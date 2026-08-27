package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/digital-humanities-go/internal/audit"
	"github.com/11DingKing/digital-humanities-go/internal/auth"
	"github.com/11DingKing/digital-humanities-go/internal/clock"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/repository"
	"github.com/11DingKing/digital-humanities-go/internal/storage"
	"time"
)

type Service struct {
	DB       *sql.DB
	Users    repository.Users
	Sessions repository.Sessions
	Projects repository.Projects
	Corpora  repository.Corpora
	Batches  repository.Batches
	Tasks    repository.Tasks
	Audit    audit.Logger
	Clock    clock.Clock
}

func New(db *sql.DB) *Service {
	return &Service{DB: db, Users: repository.Users{DB: db}, Sessions: repository.Sessions{DB: db}, Projects: repository.Projects{DB: db}, Corpora: repository.Corpora{DB: db}, Batches: repository.Batches{DB: db}, Tasks: repository.Tasks{DB: db}, Audit: audit.Logger{DB: db}, Clock: clock.Real{}}
}
func (s *Service) Register(ctx context.Context, email, name, password string, role domain.Role) (domain.User, error) {
	if email == "" || password == "" {
		return domain.User{}, domain.ErrInvalid
	}
	u := domain.User{Email: email, Name: name, Role: role, PasswordHash: auth.HashPassword(password), CreatedAt: s.Clock.Now()}
	id, e := s.Users.Create(ctx, u)
	u.ID = id
	return u, e
}
func (s *Service) Login(ctx context.Context, email, password string) (string, domain.User, error) {
	u, e := s.Users.ByEmail(ctx, email)
	if e != nil {
		return "", u, domain.ErrForbidden
	}
	if !auth.VerifyPassword(u.PasswordHash, password) {
		return "", u, domain.ErrForbidden
	}
	tok, e := auth.NewToken()
	if e != nil {
		return "", u, e
	}
	ex := s.Clock.Now().Add(8 * time.Hour)
	e = s.Sessions.Create(ctx, domain.Session{UserID: u.ID, TokenHash: auth.HashToken(tok), ExpiresAt: &ex, CreatedAt: s.Clock.Now()})
	return tok, u, e
}
func (s *Service) Authenticate(ctx context.Context, tok string) (domain.User, error) {
	ss, e := s.Sessions.Find(ctx, auth.HashToken(tok), s.Clock.Now())
	if e != nil {
		return domain.User{}, e
	}
	return s.Users.ByID(ctx, ss.UserID)
}
func (s *Service) Logout(ctx context.Context, tok string) error {
	return s.Sessions.Revoke(ctx, auth.HashToken(tok), s.Clock.Now())
}
func (s *Service) CreateProject(ctx context.Context, u domain.User, name, desc string, quota int64) (domain.Project, error) {
	if u.Role != domain.RoleLead {
		return domain.Project{}, domain.ErrForbidden
	}
	if quota <= 0 {
		return domain.Project{}, domain.ErrInvalid
	}
	now := s.Clock.Now()
	p := domain.Project{Name: name, Description: desc, QuotaBytes: quota, Status: domain.Draft, CreatedAt: now, UpdatedAt: now}
	id, e := s.Projects.Create(ctx, p)
	p.ID = id
	p.Version = 1
	return p, e
}
func (s *Service) AddCorpus(ctx context.Context, u domain.User, pid int64, title, lang, license string, sens domain.Sensitivity, size int64) (domain.Corpus, error) {
	if u.Role != domain.RoleLead && u.Role != domain.RoleCurator {
		return domain.Corpus{}, domain.ErrForbidden
	}
	if size <= 0 {
		return domain.Corpus{}, domain.ErrInvalid
	}
	var c domain.Corpus
	err := storage.Tx(ctx, s.DB, func(tx *sql.Tx) error {
		var quota, used int64
		if e := tx.QueryRowContext(ctx, "SELECT quota_bytes,COALESCE((SELECT SUM(bytes) FROM corpora WHERE project_id=? AND status!='withdrawn'),0) FROM projects WHERE id=?", pid, pid).Scan(&quota, &used); e != nil {
			return e
		}
		if used+size > quota {
			return fmt.Errorf("%w: quota exceeded", domain.ErrConflict)
		}
		now := s.Clock.Now()
		res, e := tx.ExecContext(ctx, "INSERT INTO corpora(project_id,title,language,license,sensitivity,status,bytes,version,created_at,updated_at) VALUES(?,?,?,?,?,?,?,1,?,?)", pid, title, lang, license, sens, domain.Collected, size, now.Format(time.RFC3339Nano), now.Format(time.RFC3339Nano))
		if e != nil {
			return e
		}
		c = domain.Corpus{ProjectID: pid, Title: title, Language: lang, License: license, Sensitivity: sens, Status: domain.Collected, Bytes: size, Version: 1, CreatedAt: now, UpdatedAt: now}
		c.ID, _ = res.LastInsertId()
		return nil
	})
	return c, err
}
func (s *Service) AdvanceCorpus(ctx context.Context, u domain.User, id int64, next domain.CorpusStatus, req string) error {
	c, e := s.Corpora.Get(ctx, id)
	if e != nil {
		return e
	}
	if u.Role != domain.RoleLead && u.Role != domain.RoleCurator {
		return domain.ErrForbidden
	}
	if e = c.Transition(next); e != nil {
		return e
	}
	e = storage.Tx(ctx, s.DB, func(tx *sql.Tx) error {
		res, e := tx.ExecContext(ctx, "UPDATE corpora SET status=?,version=version+1,updated_at=? WHERE id=? AND status=? AND version=?", next, s.Clock.Now().Format(time.RFC3339Nano), id, c.Status, c.Version)
		if e != nil {
			return e
		}
		n, _ := res.RowsAffected()
		if n != 1 {
			return domain.ErrConflict
		}
		return nil
	})
	if e == nil {
		e = s.Audit.Record(ctx, domain.AuditEvent{ActorID: u.ID, EntityType: "corpus", EntityID: id, Action: "transition", Outcome: "ok", RequestID: req, CreatedAt: s.Clock.Now()})
	}
	return e
}
func (s *Service) CreateBatch(ctx context.Context, u domain.User, cid int64, name string, priority, conc int) (domain.Batch, error) {
	if u.Role != domain.RoleLead && u.Role != domain.RoleCurator {
		return domain.Batch{}, domain.ErrForbidden
	}
	if priority < 1 || priority > 10 || conc < 1 || conc > 32 {
		return domain.Batch{}, domain.ErrInvalid
	}
	b := domain.Batch{CorpusID: cid, Name: name, Status: domain.BatchPending, Priority: priority, Concurrency: conc, CreatedAt: s.Clock.Now()}
	id, e := s.Batches.Create(ctx, b)
	b.ID = id
	b.Version = 1
	return b, e
}
func (s *Service) StartBatch(ctx context.Context, u domain.User, id int64, req string) error {
	if u.Role != domain.RoleLead && u.Role != domain.RoleCurator {
		return domain.ErrForbidden
	}
	b, e := s.Batches.Get(ctx, id)
	if e != nil {
		return e
	}
	var corpusStatus string
	if e = s.DB.QueryRowContext(ctx, "SELECT c.status FROM corpora c JOIN batches b ON b.corpus_id=c.id WHERE b.id=?", id).Scan(&corpusStatus); e != nil {
		return e
	}
	switch domain.CorpusStatus(corpusStatus) {
	case domain.Released:
		return fmt.Errorf("%w: released corpus", domain.ErrConflict)
	case domain.Withdrawn:
		return fmt.Errorf("%w: withdrawn corpus", domain.ErrConflict)
	}
	if e = b.Transition(domain.BatchRunning); e != nil {
		return e
	}
	return storage.Tx(ctx, s.DB, func(tx *sql.Tx) error {
		res, e := tx.ExecContext(ctx, "UPDATE batches SET status='running',version=version+1,started_at=? WHERE id=? AND status=? AND version=?", s.Clock.Now().Format(time.RFC3339Nano), id, b.Status, b.Version)
		if e != nil {
			return e
		}
		n, _ := res.RowsAffected()
		if n != 1 {
			return domain.ErrConflict
		}
		return nil
	})
}
