package worker

import (
	"context"
	"errors"
	"testing"
)

func TestCancelledWorkerRetryStopsBeforeSecondAttempt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	attempts := 0
	err := (&Worker{}).ExecuteWithRetry(ctx, func(callCtx context.Context) error {
		attempts++
		if attempts == 1 {
			cancel()
			return errors.New("temporary shard failure")
		}
		if callCtx.Err() == nil {
			t.Fatalf("retry callback lost cancellation context")
		}
		return callCtx.Err()
	})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation, got %v", err)
	}
	if attempts != 1 {
		t.Fatalf("expected one attempt after cancellation, got %d", attempts)
	}
}
