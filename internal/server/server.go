package server

import (
	"context"
	"haha/internal/logger"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	logg       *logger.Logger
}

func (s *Server) Run(url string, handler http.Handler, logg *logger.Logger) error {
	s.httpServer = &http.Server{
		Addr:           url,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	s.logg = logg

	s.logg.Info("init http server")

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logg.Info("shutdown http server")
	return s.httpServer.Shutdown(ctx)
}
