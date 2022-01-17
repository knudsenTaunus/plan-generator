package router

import "net/http"

func (r *Router) registerRoutes() {
	r.Router.Handle("/rest/api/v1/plan/generate", r.PlanHandler).Methods(http.MethodPost)
}
