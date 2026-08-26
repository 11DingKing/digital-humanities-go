package repository

import (
	"context"
	"database/sql"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"time"
)

type Users struct{ DB *sql.DB }

func (r Users) Create(ctx context.Context, u domain.User) (int64, error) {
	res, e := r.DB.ExecContext(ctx, "INSERT INTO users(email,name,role,password_hash,created_at) VALUES(?,?,?,?,?)", u.Email, u.Name, u.Role, u.PasswordHash, u.CreatedAt.Format(time.RFC3339Nano))
	if e != nil {
		return 0, e
	}
	return res.LastInsertId()
}
func (r Users) ByEmail(ctx context.Context, email string) (domain.User, error) {
	var u domain.User
	var t string
	e := r.DB.QueryRowContext(ctx, "SELECT id,email,name,role,password_hash,created_at FROM users WHERE email=?", email).Scan(&u.ID, &u.Email, &u.Name, &u.Role, &u.PasswordHash, &t)
	if e == nil {
		u.CreatedAt, _ = time.Parse(time.RFC3339Nano, t)
	}
	return u, e
}
func (r Users) ByID(ctx context.Context, id int64) (domain.User, error) {
	var u domain.User
	var t string
	e := r.DB.QueryRowContext(ctx, "SELECT id,email,name,role,password_hash,created_at FROM users WHERE id=?", id).Scan(&u.ID, &u.Email, &u.Name, &u.Role, &u.PasswordHash, &t)
	if e == nil {
		u.CreatedAt, _ = time.Parse(time.RFC3339Nano, t)
	}
	return u, e
}
