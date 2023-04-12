package name

import (
	"awesomeProject/internal/service"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Handler struct {
	svc service.Name
}

func New(svc service.Name) *Handler {
	return &Handler{svc}
}

func (h *Handler) Index(c *gofr.Context) (interface{}, error) {
	name := c.Param("name")
	if name == "" {
		return nil, errors.InvalidParam{[]string{"name"}}
	}
	return h.svc.Greet(name)
}
