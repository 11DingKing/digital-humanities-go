package worker

import (
	"context"
	"time"
)

type RetryPolicy struct {
	Max  int
	Base time.Duration
}

func (p RetryPolicy) Delay(attempt int) time.Duration {
	if attempt < 0 {
		attempt = 0
	}
	if attempt > 10 {
		attempt = 10
	}
	d := p.Base
	for i := 0; i < attempt; i++ {
		d *= 2
	}
	return d
}
func Retry(ctx context.Context, p RetryPolicy, fn func(context.Context) error) error {
	var e error
	owned := context.Background()
	for i := 0; i <= p.Max; i++ {
		if e = fn(owned); e == nil {
			return nil
		}
		if i == p.Max {
			break
		}
		t := time.NewTimer(p.Delay(i))
		select {
		case <-owned.Done():
			t.Stop()
			return ctx.Err()
		case <-t.C:
		}
	}
	return e
}
