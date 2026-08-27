package service

import (
	"context"
	"errors"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"testing"
	"time"
)

func TestRunningAnnotationKeepsItsAssignee(t *testing.T) {
	s := setup(t)
	leadUser := lead(t, s)
	second, err := s.Register(context.Background(), "second@x", "Second", "p", domain.RoleAnnotator)
	if err != nil {
		t.Fatal(err)
	}
	p, err := s.CreateProject(context.Background(), leadUser, "Assignments", "", 100)
	if err != nil {
		t.Fatal(err)
	}
	c, err := s.AddCorpus(context.Background(), leadUser, p.ID, "Segments", "en", "CC", domain.Public, 1)
	if err != nil {
		t.Fatal(err)
	}
	b, err := s.CreateBatch(context.Background(), leadUser, c.ID, "Batch", 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	id, err := s.Tasks.Create(context.Background(), domain.AnnotationTask{BatchID: b.ID, AssigneeID: leadUser.ID, Segment: "line", Status: domain.TaskQueued, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	if err != nil {
		t.Fatal(err)
	}
	if err = s.Tasks.Claim(context.Background(), id, 1, leadUser.ID, time.Now().Add(time.Hour)); err != nil {
		t.Fatal(err)
	}
	if err = s.ReassignTask(context.Background(), leadUser, id, second.ID); !errors.Is(err, domain.ErrConflict) {
		t.Fatalf("running task was reassigned: %v", err)
	}
}
