package query

import (
	"context"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/repository"
	"github.com/11DingKing/digital-humanities-go/internal/testutil"
	"testing"
	"time"
)

func TestSummaryQueries(t *testing.T) {
	db := testutil.DB(t)
	ctx := context.Background()
	p := repository.Projects{DB: db}
	pid, _ := p.Create(ctx, domain.Project{Name: "P", QuotaBytes: 100, Status: domain.Active, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	s, e := ProjectSummary(ctx, db, pid)
	if e != nil || s.ProjectID != pid || s.CorpusCount != 0 {
		t.Fatalf("%+v %v", s, e)
	}
}
func TestSummaryMissing(t *testing.T) {
	db := testutil.DB(t)
	if _, e := ProjectSummary(context.Background(), db, 88); e != nil {
		t.Fatal(e)
	}
}
