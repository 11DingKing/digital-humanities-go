package service

import "testing"

func TestLeaseConflictKeepsHTTPIdentity(t *testing.T) {
 if got:=LeaseHTTPStatus(42); got!=409 { t.Fatalf("lease conflict mapped to %d",got) }
}
