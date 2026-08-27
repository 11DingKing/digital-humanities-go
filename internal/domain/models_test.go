package domain

import "testing"

func TestCorpusLifecycle(t *testing.T) {
	cases := []struct {
		from, to CorpusStatus
		ok       bool
	}{{Collected, Cleansed, true}, {Cleansed, Licensed, true}, {Licensed, Released, true}, {Licensed, Withdrawn, true}, {Released, Withdrawn, true}, {Collected, Released, false}, {Released, Licensed, false}}
	for _, c := range cases {
		e := (Corpus{Status: c.from}).Transition(c.to)
		if (e == nil) != c.ok {
			t.Fatalf("%s to %s", c.from, c.to)
		}
	}
}
func TestBatchLifecycle(t *testing.T) {
	if e := (Batch{Status: BatchPending}).Transition(BatchRunning); e != nil {
		t.Fatal(e)
	}
	if e := (Batch{Status: BatchDone}).Transition(BatchRunning); e == nil {
		t.Fatal("invalid transition accepted")
	}
}
func TestTaskLifecycle(t *testing.T) {
	if e := (AnnotationTask{Status: TaskQueued}).Transition(TaskRunning); e != nil {
		t.Fatal(e)
	}
	if e := (AnnotationTask{Status: TaskSucceeded}).Transition(TaskQueued); e == nil {
		t.Fatal("terminal transition accepted")
	}
}
func TestRolesDistinct(t *testing.T) {
	if RoleLead == RoleAnnotator {
		t.Fatal("roles collapsed")
	}
}
