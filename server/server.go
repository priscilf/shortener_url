package server

import (
	"context"
	"net/http"

	"github.com/priscilf/shortener_url/internal/config"
)

type Server struct {
	httpServer http.Server
}

func NewServer(address string, config config.HTTPServerConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: http.Server{
			Addr:         address,
			Handler:      handler,
			ReadTimeout:  config.Timeout,
			WriteTimeout: config.Timeout,
			IdleTimeout:  config.IdleTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
