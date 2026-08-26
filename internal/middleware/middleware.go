package middleware

import (
	"context"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/service"
	"net/http"
)

type key int

const userKey key = 1

func WithUser(u domain.User) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userKey, u)))
		})
	}
}
func User(r *http.Request) (domain.User, bool) {
	u, ok := r.Context().Value(userKey).(domain.User)
	return u, ok
}
func Require(s *service.Service, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tok := r.Header.Get("Authorization")
		if len(tok) < 8 {
			http.Error(w, "unauthorized", 401)
			return
		}
		u, e := s.Authenticate(r.Context(), tok[7:])
		if e != nil {
			http.Error(w, "unauthorized", 401)
			return
		}
		WithUser(u)(next).ServeHTTP(w, r)
	})
}
