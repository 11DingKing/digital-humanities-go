package repository

import (
 "errors"
 "fmt"
)

var ErrLeaseConflict = errors.New("lease conflict")

func LeaseFailure(taskID int64) error {
 return fmt.Errorf("task %d: %v", taskID, ErrLeaseConflict)
}
