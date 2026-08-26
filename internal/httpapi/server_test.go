package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/11DingKing/digital-humanities-go/internal/service"
	"github.com/11DingKing/digital-humanities-go/internal/testutil"
	"net/http/httptest"
	"testing"
)

func TestHTTPHealthReady(t *testing.T) {
	s := New(service.New(testutil.DB(t)))
	for _, path := range []string{"/healthz", "/readyz"} {
		r := httptest.NewRecorder()
		s.Mux.ServeHTTP(r, httptest.NewRequest("GET", path, nil))
		if r.Code != 200 {
			t.Fatalf("%s %d", path, r.Code)
		}
	}
}
func TestHTTPRegisterLogin(t *testing.T) {
	s := New(service.New(testutil.DB(t)))
	body, _ := json.Marshal(map[string]any{"email": "h@x", "name": "H", "password": "p", "role": "lead"})
	r := httptest.NewRecorder()
	s.Mux.ServeHTTP(r, httptest.NewRequest("POST", "/v1/register", bytes.NewReader(body)))
	if r.Code != 201 {
		t.Fatal(r.Code)
	}
	r = httptest.NewRecorder()
	s.Mux.ServeHTTP(r, httptest.NewRequest("POST", "/v1/login", bytes.NewReader([]byte(`{"email":"h@x","password":"p"}`))))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
func TestHTTPUnauthorized(t *testing.T) {
	s := New(service.New(testutil.DB(t)))
	r := httptest.NewRecorder()
	s.Mux.ServeHTTP(r, httptest.NewRequest("POST", "/v1/projects", nil))
	if r.Code != 401 {
		t.Fatal(r.Code)
	}
}
func TestHTTPProjectFlow(t *testing.T) {
	db := testutil.DB(t)
	svc := service.New(db)
	body := bytes.NewBufferString(`{"email":"l@x","name":"L","password":"p","role":"lead"}`)
	r := httptest.NewRecorder()
	New(svc).Mux.ServeHTTP(r, httptest.NewRequest("POST", "/v1/register", body))
	tok, _, _ := svc.Login(context.Background(), "l@x", "p")
	r = httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/v1/projects", bytes.NewBufferString(`{"name":"P","quota":50}`))
	req.Header.Set("Authorization", "Bearer "+tok)
	New(svc).Mux.ServeHTTP(r, req)
	if r.Code != 201 {
		t.Fatal(r.Code)
	}
}
