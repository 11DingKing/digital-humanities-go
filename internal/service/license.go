package service

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/repository"
	"github.com/11DingKing/digital-humanities-go/internal/storage"
)

// ChangeCorpusLicense updates a corpus license and records the change
// history in a single transaction. If the history write fails, the
// license update rolls back, so the corpus license never advances to a
// new value without a recoverable authorization trail.
func (s *Service) ChangeCorpusLicense(ctx context.Context, u domain.User, corpusID int64, license, reason string) error {
	if u.Role != domain.RoleLead && u.Role != domain.RoleCurator {
		return domain.ErrForbidden
	}
	c, err := s.Corpora.Get(ctx, corpusID)
	if err != nil {
		return err
	}
	if license == "" || reason == "" {
		return domain.ErrInvalid
	}
	now := s.Clock.Now()
	return storage.Tx(ctx, s.DB, func(tx *sql.Tx) error {
		if err := s.Corpora.UpdateLicenseTx(ctx, tx, corpusID, license, c.Version); err != nil {
			return err
		}
		return repository.RecordLicenseChange(ctx, tx, corpusID, u.ID, c.License, license, reason, now)
	})
}
