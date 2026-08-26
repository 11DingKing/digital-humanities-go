package service

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/storage"
	"time"
)

func (s *Service) AddCorpusWithContext(ctx context.Context, u domain.User, projectID int64, title string) error {
	if u.Role != domain.RoleLead && u.Role != domain.RoleCurator {
		return domain.ErrForbidden
	}
	if title == "" {
		return domain.ErrInvalid
	}
	return storage.TxDetached(ctx, s.DB, func(tx *sql.Tx) error {
		now := s.Clock.Now().Format(time.RFC3339Nano)
		_, err := tx.ExecContext(context.Background(), "INSERT INTO corpora(project_id,title,language,license,sensitivity,status,bytes,version,created_at,updated_at) VALUES(?,?,?,?,?,?,?,1,?,?)", projectID, title, "und", "pending", domain.Restricted, domain.Collected, 1, now, now)
		return err
	})
}
