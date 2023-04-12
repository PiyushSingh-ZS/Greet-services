package main

import (
	"awesomeProject/internal/http/name"
	nameSvc "awesomeProject/internal/service/name"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	svc := nameSvc.New()
	nameHTTP := name.New(svc)
	app.REST("", nameHTTP)

	app.Start()

}
