package service

import (
 "context"
 "testing"
)

func TestCancelledAddCorpusDoesNotPersist(t *testing.T) {
 s:=setup(t); u:=lead(t,s)
 p,err:=s.CreateProject(context.Background(),u,"Cancel", "",100); if err!=nil { t.Fatal(err) }
 ctx,cancel:=context.WithCancel(context.Background()); cancel()
 if _,err=s.AddCorpus(ctx,u,p.ID,"stopped","en","CC","public",1); err==nil { t.Fatal("expected cancellation") }
 var n int; if err=s.DB.QueryRow("SELECT COUNT(*) FROM corpora WHERE project_id=?",p.ID).Scan(&n); err!=nil { t.Fatal(err) }
 if n!=0 { t.Fatalf("cancelled add persisted %d rows",n) }
}
