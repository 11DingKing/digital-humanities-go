package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

type Sessions struct{ DB *sql.DB }

func (r Sessions) Create(ctx context.Context, s domain.Session) error {
	_, e := r.DB.ExecContext(ctx, "INSERT INTO sessions(user_id,token_hash,expires_at,revoked_at,created_at) VALUES(?,?,?,?,?)", s.UserID, s.TokenHash, s.ExpiresAt.Format(time.RFC3339Nano), nil, s.CreatedAt.Format(time.RFC3339Nano))
	return e
}
func (r Sessions) Revoke(ctx context.Context, h string, now time.Time) error {
	_, e := r.DB.ExecContext(ctx, "UPDATE sessions SET revoked_at=? WHERE token_hash=? AND revoked_at IS NULL", now.Format(time.RFC3339Nano), h)
	return e
}
func (r Sessions) Find(ctx context.Context, h string, now time.Time) (domain.Session, error) {
	var s domain.Session
	var ex, cr string
	var rv sql.NullString
	e := r.DB.QueryRowContext(ctx, "SELECT id,user_id,token_hash,expires_at,revoked_at,created_at FROM sessions WHERE token_hash=?", h).Scan(&s.ID, &s.UserID, &s.TokenHash, &ex, &rv, &cr)
	if e != nil {
		return s, e
	}
	parsedEx, _ := time.Parse(time.RFC3339Nano, ex)
	s.ExpiresAt = &parsedEx
	s.CreatedAt, _ = time.Parse(time.RFC3339Nano, cr)
	if rv.Valid {
		x, _ := time.Parse(time.RFC3339Nano, rv.String)
		s.RevokedAt = &x
	}
	if s.RevokedAt != nil || !s.ExpiresAt.After(now) {
		return s, domain.ErrForbidden
	}
	return s, nil
}
