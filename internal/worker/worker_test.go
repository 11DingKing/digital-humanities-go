package worker

import (
	"context"
	"github.com/11DingKing/digital-humanities-go/internal/testutil"
	"testing"
	"time"
)

func TestWorkerStops(t *testing.T) {
	w := New(testutil.DB(t), time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() { w.Run(ctx); close(done) }()
	cancel()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("worker did not stop")
	}
}
func TestWorkerExplicitStop(t *testing.T) {
	w := New(testutil.DB(t), time.Hour)
	done := make(chan struct{})
	go func() { w.Run(context.Background()); close(done) }()
	w.Stop()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("stop")
	}
}
