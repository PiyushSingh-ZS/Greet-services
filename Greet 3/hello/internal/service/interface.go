package service

import "developer.zopsmart.com/go/gofr/pkg/gofr"

type Name interface {
	Greet(name string) (string, error)
}

type Service interface {
	Greet(ctx *gofr.Context, name string) (string, error)
}
