package service

import (
	service2 "developer.zopsmart.com/go/gofr/examples/universal-example/gofr-services/service"
	"developer.zopsmart.com/go/gofr/examples/using-http-service/services"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"strings"
)

type service struct {
	name services.HTTPService
	bye  services.HTTPService
}

func New(http services.HTTPService, http1 services.HTTPService) *service {
	return &service{name: http, bye: http1}
}

func (s *service) Greet(ctx *gofr.Context, name string) (string, error) {

	endPoint := ctx.Request().RequestURI
	endPoint = strings.TrimPrefix(endPoint, "/")

	if endPoint == "bye" {
		res, err := s.bye.Get(ctx, "bye", nil)
		if err != nil {
			return "", err
		}
		var r service2.Response
		err = s.bye.Bind(res.Body, &r)
		if err != nil {
			return "", err
		}
		return r.Data, err
	}

	if name == "" {
		return "hello", nil
	}

	res, err := s.name.Get(ctx, "?name="+name, nil)

	var r service2.Response
	err = s.name.Bind(res.Body, &r)
	return r.Data, err
}
