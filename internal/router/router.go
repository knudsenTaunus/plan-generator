package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Router      mux.Router
	PlanHandler http.Handler
}

func New(opts ...Option) *Router {
	r := &Router{
		Router: *mux.NewRouter(),
	}

	for _, opt := range opts {
		opt(r)
	}

	r.registerRoutes()

	return r
}

type Option func(*Router)
