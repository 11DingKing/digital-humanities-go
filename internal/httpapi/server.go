package httpapi

import (
	"encoding/json"
	"github.com/11DingKing/digital-humanities-go/internal/domain"
	"github.com/11DingKing/digital-humanities-go/internal/middleware"
	"github.com/11DingKing/digital-humanities-go/internal/service"
	"net/http"
	"strings"
)

type Server struct {
	S   *service.Service
	Mux *http.ServeMux
}

func New(s *service.Service) *Server {
	x := &Server{S: s, Mux: http.NewServeMux()}
	x.routes()
	return x
}
func (x *Server) routes() {
	x.Mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { write(w, 200, map[string]string{"status": "ok"}) })
	x.Mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if e := x.S.DB.PingContext(r.Context()); e != nil {
			write(w, 503, map[string]string{"status": "not_ready"})
			return
		}
		write(w, 200, map[string]string{"status": "ready"})
	})
	x.Mux.HandleFunc("/v1/register", x.register)
	x.Mux.HandleFunc("/v1/login", x.login)
	x.Mux.Handle("/v1/projects", middleware.Require(x.S, http.HandlerFunc(x.project)))
	x.Mux.Handle("/v1/corpora", middleware.Require(x.S, http.HandlerFunc(x.corpus)))
}
func (x *Server) register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		write(w, 405, nil)
		return
	}
	var in struct {
		Email, Name, Password string
		Role                  domain.Role
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		write(w, 400, map[string]string{"error": "invalid_json"})
		return
	}
	u, e := x.S.Register(r.Context(), in.Email, in.Name, in.Password, in.Role)
	if e != nil {
		writeErr(w, e)
		return
	}
	write(w, 201, u)
}
func (x *Server) login(w http.ResponseWriter, r *http.Request) {
	var in struct{ Email, Password string }
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		write(w, 400, nil)
		return
	}
	t, u, e := x.S.Login(r.Context(), in.Email, in.Password)
	if e != nil {
		writeErr(w, e)
		return
	}
	write(w, 200, map[string]any{"token": t, "user": u})
}
func (x *Server) project(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.User(r)
	if r.Method == "POST" {
		var in struct {
			Name, Description string
			Quota             int64
		}
		_ = json.NewDecoder(r.Body).Decode(&in)
		p, e := x.S.CreateProject(r.Context(), u, in.Name, in.Description, in.Quota)
		if e != nil {
			writeErr(w, e)
			return
		}
		write(w, 201, p)
		return
	}
	write(w, 405, nil)
}
func (x *Server) corpus(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.User(r)
	if r.Method != "POST" {
		write(w, 405, nil)
		return
	}
	var in struct {
		ProjectID                int64
		Title, Language, License string
		Sensitivity              domain.Sensitivity
		Bytes                    int64
	}
	_ = json.NewDecoder(r.Body).Decode(&in)
	c, e := x.S.AddCorpus(r.Context(), u, in.ProjectID, in.Title, in.Language, in.License, in.Sensitivity, in.Bytes)
	if e != nil {
		writeErr(w, e)
		return
	}
	write(w, 201, c)
}
func write(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		_ = json.NewEncoder(w).Encode(v)
	}
}
func writeErr(w http.ResponseWriter, e error) {
	status := 500
	if strings.Contains(e.Error(), "forbidden") {
		status = 403
	}
	if e == domain.ErrInvalid || strings.Contains(e.Error(), "invalid") {
		status = 400
	}
	if e == domain.ErrConflict || strings.Contains(e.Error(), "conflict") {
		status = 409
	}
	write(w, status, map[string]string{"error": e.Error()})
}
