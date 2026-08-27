package service

import (
	"context"
	"fmt"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
)

func (s *Service) Guard1(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard2(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard3(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard4(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard5(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard6(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard7(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard8(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard9(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard10(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard11(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}

func (s *Service) Guard12(ctx context.Context, u domain.User, objectID int64) error {
	if objectID <= 0 {
		return fmt.Errorf("%w: object id", domain.ErrInvalid)
	}
	if u.ID <= 0 {
		return domain.ErrForbidden
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("guard context: %w", err)
	}
	return nil
}
