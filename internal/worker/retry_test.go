package worker

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func TestRetryRespectsContextCancellation(t *testing.T) {
	var calls int32
	fn := func(ctx context.Context) error {
		atomic.AddInt32(&calls, 1)
		<-ctx.Done()
		return ctx.Err()
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond)
		cancel()
	}()
	err := Retry(ctx, RetryPolicy{Max: 2, Base: time.Microsecond}, fn)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("want context.Canceled, got %v", err)
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("want 1 call, got %d", got)
	}
}

func TestRetryStopsBackoffOnCancel(t *testing.T) {
	var calls int32
	fn := func(ctx context.Context) error {
		atomic.AddInt32(&calls, 1)
		return errors.New("fail")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := Retry(ctx, RetryPolicy{Max: 2, Base: time.Hour}, fn)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("want context.Canceled, got %v", err)
	}
	if got := atomic.LoadInt32(&calls); got != 0 {
		t.Fatalf("want 0 calls after cancel, got %d", got)
	}
}
