package storage

import (
	"context"
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

func Open(ctx context.Context, path string) (*sql.DB, error) {
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}
	db, e := sql.Open("sqlite", path)
	if e != nil {
		return nil, e
	}
	db.SetMaxOpenConns(1)
	if e = db.PingContext(ctx); e != nil {
		db.Close()
		return nil, e
	}
	return db, nil
}
func Migrate(ctx context.Context, db *sql.DB) error {
	_, e := db.ExecContext(ctx, `PRAGMA foreign_keys=ON; CREATE TABLE IF NOT EXISTS schema_migrations(version INTEGER PRIMARY KEY, applied_at TEXT NOT NULL);`)
	if e != nil {
		return e
	}
	var v int
	_ = db.QueryRowContext(ctx, "SELECT COALESCE(MAX(version),0) FROM schema_migrations").Scan(&v)
	if v < 1 {
		if e := migration1(ctx, db); e != nil {
			return e
		}
		_, e = db.ExecContext(ctx, "INSERT INTO schema_migrations(version,applied_at) VALUES(1,datetime('now'))")
	}
	return e
}
func migration1(ctx context.Context, db *sql.DB) error {
	_, e := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY,email TEXT UNIQUE NOT NULL,name TEXT NOT NULL,role TEXT NOT NULL,password_hash TEXT NOT NULL,created_at TEXT NOT NULL); CREATE TABLE IF NOT EXISTS sessions(id INTEGER PRIMARY KEY,user_id INTEGER NOT NULL REFERENCES users(id),token_hash TEXT UNIQUE NOT NULL,expires_at TEXT NOT NULL,revoked_at TEXT,created_at TEXT NOT NULL); CREATE TABLE IF NOT EXISTS projects(id INTEGER PRIMARY KEY,name TEXT NOT NULL,description TEXT,quota_bytes INTEGER NOT NULL,status TEXT NOT NULL,version INTEGER NOT NULL DEFAULT 1,created_at TEXT NOT NULL,updated_at TEXT NOT NULL); CREATE TABLE IF NOT EXISTS corpora(id INTEGER PRIMARY KEY,project_id INTEGER NOT NULL REFERENCES projects(id),title TEXT NOT NULL,language TEXT NOT NULL,license TEXT NOT NULL,sensitivity TEXT NOT NULL,status TEXT NOT NULL,bytes INTEGER NOT NULL,version INTEGER NOT NULL DEFAULT 1,created_at TEXT NOT NULL,updated_at TEXT NOT NULL); CREATE TABLE IF NOT EXISTS batches(id INTEGER PRIMARY KEY,corpus_id INTEGER NOT NULL REFERENCES corpora(id),name TEXT NOT NULL,status TEXT NOT NULL,priority INTEGER NOT NULL,concurrency INTEGER NOT NULL,version INTEGER NOT NULL DEFAULT 1,started_at TEXT,finished_at TEXT,created_at TEXT NOT NULL); CREATE TABLE IF NOT EXISTS annotation_tasks(id INTEGER PRIMARY KEY,batch_id INTEGER NOT NULL REFERENCES batches(id),assignee_id INTEGER NOT NULL REFERENCES users(id),segment TEXT NOT NULL,status TEXT NOT NULL,attempts INTEGER NOT NULL,lease_until TEXT,version INTEGER NOT NULL DEFAULT 1,created_at TEXT NOT NULL,updated_at TEXT NOT NULL); CREATE TABLE IF NOT EXISTS analyses(id INTEGER PRIMARY KEY,project_id INTEGER NOT NULL REFERENCES projects(id),corpus_id INTEGER NOT NULL REFERENCES corpora(id),author_id INTEGER NOT NULL REFERENCES users(id),kind TEXT NOT NULL,source TEXT NOT NULL,ai_use TEXT NOT NULL,status TEXT NOT NULL,result TEXT,created_at TEXT NOT NULL); CREATE TABLE IF NOT EXISTS audit_events(id INTEGER PRIMARY KEY,actor_id INTEGER REFERENCES users(id),entity_type TEXT NOT NULL,entity_id INTEGER NOT NULL,action TEXT NOT NULL,outcome TEXT NOT NULL,request_id TEXT NOT NULL,details TEXT,created_at TEXT NOT NULL); CREATE TABLE IF NOT EXISTS license_changes(id INTEGER PRIMARY KEY,corpus_id INTEGER NOT NULL REFERENCES corpora(id),actor_id INTEGER NOT NULL REFERENCES users(id),previous TEXT NOT NULL,current TEXT NOT NULL,reason TEXT NOT NULL,created_at TEXT NOT NULL); CREATE UNIQUE INDEX IF NOT EXISTS idx_tasks_batch_assignee ON annotation_tasks(batch_id,assignee_id,segment); CREATE INDEX IF NOT EXISTS idx_tasks_status ON annotation_tasks(status,lease_until); CREATE INDEX IF NOT EXISTS idx_audit_entity ON audit_events(entity_type,entity_id,created_at);`)
	return e
}
func Tx(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	tx, e := db.BeginTx(ctx, nil)
	if e != nil {
		return e
	}
	if e = fn(tx); e != nil {
		_ = tx.Rollback()
		return fmt.Errorf("transaction: %w", e)
	}
	return tx.Commit()
}
