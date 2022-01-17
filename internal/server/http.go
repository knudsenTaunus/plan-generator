package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	DefaultPort           = "8080"
	DefaultAllowedHeaders = ""
	DefaultAllowedOrigins = "*"
)

type HTTPServerOption func(*HTTPServer)

type HTTPServer struct {
	router         http.Handler
	port           string
	AllowedHeaders []string
	AllowedOrigins []string
	AllowedMethods []string
}

func NewHTTPServer(opts ...HTTPServerOption) *HTTPServer {
	s := &HTTPServer{
		router:         mux.NewRouter(),
		port:           DefaultPort,
		AllowedHeaders: []string{DefaultAllowedHeaders},
		AllowedOrigins: []string{DefaultAllowedOrigins},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithRouter(router http.Handler) HTTPServerOption {
	return func(s *HTTPServer) {
		s.router = router
	}
}

func WithPort(port string) HTTPServerOption {
	return func(s *HTTPServer) {
		s.port = port
	}
}

func (s *HTTPServer) WithAllowedHeaders(allowedHeaders []string) HTTPServerOption {
	return func(s *HTTPServer) {
		s.AllowedHeaders = allowedHeaders
	}
}

func (s *HTTPServer) WithAllowedOrigins(allowedOrigins []string) HTTPServerOption {
	return func(s *HTTPServer) {
		s.AllowedOrigins = allowedOrigins
	}
}

func (s *HTTPServer) WithAllowedMethods(allowedMethods []string) {
	s.AllowedMethods = allowedMethods
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *HTTPServer) Start() error {
	corsHandler := handlers.CORS(
		handlers.AllowedHeaders(s.AllowedHeaders),
		handlers.AllowedOrigins(s.AllowedOrigins),
		handlers.AllowedMethods(s.AllowedMethods),
	)(s)

	if err := http.ListenAndServe(":"+s.port, corsHandler); err != nil {
		return err
	}

	return nil
}
