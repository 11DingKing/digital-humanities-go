package service

import (
 "errors"
 "github.com/11DingKing/digital-humanities-go/internal/repository"
)

func LeaseHTTPStatus(taskID int64) int {
 err:=repository.LeaseFailure(taskID)
 if errors.Is(err, repository.ErrLeaseConflict) { return 409 }
 return 500
}
