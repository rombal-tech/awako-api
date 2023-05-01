package server

import (
	"context"
	"github.com/execaus/exloggo"

	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1 mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	exloggo.Info("server started successfully")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	exloggo.Info("server shutdown process started")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		exloggo.Error(err.Error())
	} else {
		exloggo.Info("http listener shutdown successfully")
	}
	exloggo.Info("server shutdown process completed successfully")
	return s.httpServer.Shutdown(ctx)

}
