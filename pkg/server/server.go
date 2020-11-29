package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server armazena o servidor da API
type Server struct {
	srv *http.Server
}

// GetServer retorna o servidor da API
func GetServer() *Server {
	return &Server{
		srv: &http.Server{},
	}
}

// WithAddr adiciona o endereço ao servidor
func (s *Server) WithAddr(addr string) *Server {
	s.srv.Addr = addr
	return s
}

// WithLogger adiciona o logger ao servidor
func (s *Server) WithLogger(l *log.Logger) *Server {
	s.srv.ErrorLog = l
	return s
}

// WithRouter adiciona o roteador ao servidor
func (s *Server) WithRouter(router *mux.Router) *Server {
	s.srv.Handler = router
	return s
}

// StartServer abre a conexão do servidor
func (s *Server) StartServer() error {
	if len(s.srv.Addr) == 0 {
		return errors.New("Server missing address")
	}

	if s.srv.Handler == nil {
		return errors.New("Server missing handler")
	}

	return s.srv.ListenAndServe()
}

// CloseServer fecha a conexão do servidor
func (s *Server) CloseServer() error {
	return s.srv.Close()
}
