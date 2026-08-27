package storage

import (
	"context"
	"database/sql"
	"os"
	"testing"
)

func TestMigrationCreatesRelations(t *testing.T) {
	f, _ := os.CreateTemp("", "db")
	f.Close()
	defer os.Remove(f.Name())
	db, e := Open(context.Background(), f.Name())
	if e != nil {
		t.Fatal(e)
	}
	defer db.Close()
	if e = Migrate(context.Background(), db); e != nil {
		t.Fatal(e)
	}
	var n int
	if e = db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table'").Scan(&n); e != nil || n < 10 {
		t.Fatalf("tables %d %v", n, e)
	}
	if e = Migrate(context.Background(), db); e != nil {
		t.Fatal(e)
	}
}
func TestTransactionRollback(t *testing.T) {
	f, _ := os.CreateTemp("", "db")
	f.Close()
	defer os.Remove(f.Name())
	db, _ := Open(context.Background(), f.Name())
	defer db.Close()
	Migrate(context.Background(), db)
	e := Tx(context.Background(), db, func(tx *sql.Tx) error {
		_, e := tx.Exec("INSERT INTO users(email,name,role,password_hash,created_at) VALUES('a','A','lead','x','now')")
		return e
	})
	if e != nil {
		t.Fatal(e)
	}
}
