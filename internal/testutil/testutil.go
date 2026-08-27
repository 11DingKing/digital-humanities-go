package testutil

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/storage"
	"os"
	"testing"
)

func DB(t *testing.T) *sql.DB {
	t.Helper()
	f, e := os.CreateTemp("", "dh-*.db")
	if e != nil {
		t.Fatal(e)
	}
	f.Close()
	db, e := storage.Open(context.Background(), f.Name())
	if e != nil {
		t.Fatal(e)
	}
	if e = storage.Migrate(context.Background(), db); e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { db.Close(); os.Remove(f.Name()) })
	return db
}
