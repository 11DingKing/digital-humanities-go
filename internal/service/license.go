package service

import (
	"context"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/repository"
)

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
	if err = s.Corpora.UpdateLicense(ctx, corpusID, license, c.Version); err != nil {
		return err
	}
	return repository.RecordLicenseChange(ctx, s.DB, corpusID, u.ID, c.License, license, reason, s.Clock.Now())
}
