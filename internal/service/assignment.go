package service

import (
	"context"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
)

func (s *Service) ReassignTask(ctx context.Context, u domain.User, taskID, assigneeID int64) error {
	if u.Role != domain.RoleLead && u.Role != domain.RoleCurator {
		return domain.ErrForbidden
	}
	if taskID <= 0 || assigneeID <= 0 {
		return domain.ErrInvalid
	}
	return s.Tasks.ReassignRunning(ctx, taskID, assigneeID)
}
