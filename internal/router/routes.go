package router

import "net/http"

func (rtr *Router) registerRoutes() {
	rtr.Router.Handle("/rest/api/v1/plan/generate", rtr.PlanHandler).Methods(http.MethodPost)
}
