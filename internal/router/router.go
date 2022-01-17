package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	Router      mux.Router
	PlanHandler http.Handler
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rtr.Router.ServeHTTP(w, r)
}

func New(ph http.Handler) *Router {
	r := &Router{
		PlanHandler: ph,
		Router:      *mux.NewRouter(),
	}

	r.registerRoutes()

	return r
}
