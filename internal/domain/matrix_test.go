package domain

import "testing"

func TestCorpusInvalidTransitions(t *testing.T) {
	from := []CorpusStatus{Collected, Cleansed, Licensed, Released, Withdrawn}
	to := []CorpusStatus{Collected, Cleansed, Licensed, Released, Withdrawn}
	for _, a := range from {
		for _, b := range to {
			e := (Corpus{Status: a}).Transition(b)
			if a == b && e == nil {
				t.Fatalf("self transition %s", a)
			}
		}
	}
}
func TestBatchInvalidTransitions(t *testing.T) {
	bad := []BatchStatus{BatchPending, BatchRunning, BatchPaused, BatchDone, BatchFailed, BatchCancelled}
	for _, a := range bad {
		for _, b := range bad {
			if a == BatchDone || a == BatchCancelled {
				if e := (Batch{Status: a}).Transition(b); e == nil {
					t.Fatalf("terminal %s", a)
				}
			}
		}
	}
}
func TestTaskInvalidTransitions(t *testing.T) {
	for _, a := range []TaskStatus{TaskSucceeded, TaskCancelled} {
		for _, b := range []TaskStatus{TaskQueued, TaskRunning, TaskFailed} {
			if e := (AnnotationTask{Status: a}).Transition(b); e == nil {
				t.Fatal("terminal accepted")
			}
		}
	}
}
