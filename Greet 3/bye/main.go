package main

import (
	"awesomeProject/internal/http/bye"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false
	byeHTTP := bye.New()
	app.REST("bye", byeHTTP)

	app.Start()

}
