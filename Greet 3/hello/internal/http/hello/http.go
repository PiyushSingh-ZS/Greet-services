package http

import (
	"awesomeProject/internal/service"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	svc service.Service
}

func New(svc service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Index(ctx *gofr.Context) (interface{}, error) {
	name := ctx.Param("name")

	res, err := h.svc.Greet(ctx, name)

	return res, err

}
