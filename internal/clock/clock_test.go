package clock

import (
	"testing"
	"time"
)

func TestFixedClock(t *testing.T) {
	want := time.Unix(1, 0)
	if (Fixed{want}).Now() != want {
		t.Fatal("clock")
	}
}
func TestRealClockUTC(t *testing.T) {
	if (Real{}).Now().Location() != time.UTC {
		t.Fatal("utc")
	}
}
