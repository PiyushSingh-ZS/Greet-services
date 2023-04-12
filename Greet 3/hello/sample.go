package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/service"
)

func main() {

	app := gofr.New()

	app.Server.ValidateHeaders = false

	readyURL := app.Config.Get("READY_URL")

	tempoURL := app.Config.Get("TEMPO_URL")

	readySvc := service.NewHTTPServiceWithOptions(readyURL, app.Logger, &service.Options{

		SurgeProtectorOption: &service.SurgeProtectorOption{

			Disable: true,
		},
	})

	tempoSvc := service.NewHTTPServiceWithOptions(tempoURL, app.Logger, &service.Options{

		SurgeProtectorOption: &service.SurgeProtectorOption{

			Disable: true,
		},
	})

	rSvc := ready.New(readySvc)

	rHandler := readyHandler.New(rSvc)

	app.GET("/ready", rHandler.Ready)

	app.GET("/metrics", rHandler.Metrics)

	tSvc := tempo.New(tempoSvc)

	tHandler := tempoHandler.New(tSvc)

	app.REST("{rest:[a-zA-Z0-9=\\-\\/]+}", &tHandler)

	app.Start()
}
