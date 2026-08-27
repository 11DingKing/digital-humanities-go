package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

type AnalysisInput struct {
	ProjectID, CorpusID, AuthorID int64
	Kind, Source, AIUse           string
}

func (s *Service) RecordAnalysis(ctx context.Context, u domain.User, in AnalysisInput) (domain.Analysis, error) {
	if u.Role != domain.RoleLead && u.Role != domain.RoleCurator {
		return domain.Analysis{}, domain.ErrForbidden
	}
	if in.ProjectID <= 0 || in.CorpusID <= 0 {
		return domain.Analysis{}, domain.ErrInvalid
	}
	if e := domain.ValidateAIUse(in.Source, in.AIUse); e != nil {
		return domain.Analysis{}, e
	}
	now := s.Clock.Now()
	a := domain.Analysis{ProjectID: in.ProjectID, CorpusID: in.CorpusID, AuthorID: u.ID, Kind: in.Kind, Source: in.Source, AIUse: in.AIUse, Status: "draft", CreatedAt: now}
	e := s.DB.QueryRowContext(ctx, "INSERT INTO analyses(project_id,corpus_id,author_id,kind,source,ai_use,status,created_at) VALUES(?,?,?,?,?,?,?,?) RETURNING id", a.ProjectID, a.CorpusID, a.AuthorID, a.Kind, a.Source, a.AIUse, a.Status, now.Format(time.RFC3339Nano)).Scan(&a.ID)
	return a, e
}
func (s *Service) PublishAnalysis(ctx context.Context, u domain.User, id int64, result string) error {
	if u.Role != domain.RoleLead {
		return domain.ErrForbidden
	}
	if result == "" {
		return domain.ErrInvalid
	}
	res, e := s.DB.ExecContext(ctx, "UPDATE analyses SET status='published',result=? WHERE id=? AND status='draft'", result, id)
	if e != nil {
		return e
	}
	n, _ := res.RowsAffected()
	if n != 1 {
		return domain.ErrConflict
	}
	return nil
}
func EnsureAnalysisProject(ctx context.Context, db *sql.DB, project, corpus int64) error {
	var n int
	if e := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM corpora WHERE id=? AND project_id=?", corpus, project).Scan(&n); e != nil {
		return e
	}
	if n != 1 {
		return fmt.Errorf("%w: corpus project mismatch", domain.ErrInvalid)
	}
	return nil
}
