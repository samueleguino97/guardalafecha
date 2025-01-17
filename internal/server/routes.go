package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *Server) Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", s.HealthHandler)
	r.Get("/", s.InvitationsHandler)
	r.Get("/stats", s.StatsHandler)

	return r
}

func (s *Server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) InvitationsHandler(w http.ResponseWriter, r *http.Request) {
	subdomain := s.ExtractSubdomain(r)
	_, err := s.Queries.GetTenantBySlug(r.Context(), subdomain)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	template, err := s.GetTemplateBySlug(subdomain)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	template.Render(r.Context(), w)

}

func (s *Server) StatsHandler(w http.ResponseWriter, r *http.Request) {
	subdomain := s.ExtractSubdomain(r)
	log.Println(subdomain)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
