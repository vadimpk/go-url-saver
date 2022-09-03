package server

import (
	"context"
	"go-urlsaver/internal/config"
	"net/http"
	"time"
)

type Server struct {
	HTTPServer *http.Server
}

func (s *Server) Run(cfg *config.Config, handler http.Handler) error {
	s.HTTPServer = &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.HTTPServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.HTTPServer.Shutdown(ctx)
}
