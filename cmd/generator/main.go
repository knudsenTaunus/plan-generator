package main

import (
	"github.com/knudsenTaunus/plan-generator/internal/handler"
	"github.com/knudsenTaunus/plan-generator/internal/router"
	"github.com/knudsenTaunus/plan-generator/internal/server"
	"github.com/knudsenTaunus/plan-generator/internal/service"
)

func main() {

	calculationService := service.NewCalculate()

	planHandler := handler.New(calculationService)

	r := router.New(planHandler)

	svr := server.NewHTTPServer(
		server.WithRouter(r),
	)
	svr.Start()
}
