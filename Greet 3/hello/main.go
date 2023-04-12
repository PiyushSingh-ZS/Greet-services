package main

import (
	helloHandler "awesomeProject/internal/http/hello"
	helloSvc "awesomeProject/internal/service/hello"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/service"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	nameURL := app.Config.Get("NAME_URL")
	nameSvc := service.NewHTTPServiceWithOptions(nameURL, app.Logger, &service.Options{

		SurgeProtectorOption: &service.SurgeProtectorOption{

			Disable: true,
		},
	})

	byeURL := app.Config.Get("BYE_URL")
	byeSvc := service.NewHTTPServiceWithOptions(byeURL, app.Logger, &service.Options{

		SurgeProtectorOption: &service.SurgeProtectorOption{

			Disable: true,
		},
	})

	hsvc := helloSvc.New(nameSvc, byeSvc)
	helloHTTP := helloHandler.New(hsvc)
	app.REST("hello", helloHTTP)
	app.REST("bye", helloHTTP)

	app.Start()

}
