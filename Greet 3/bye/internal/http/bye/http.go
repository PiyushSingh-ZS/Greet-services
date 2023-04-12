package bye

import "developer.zopsmart.com/go/gofr/pkg/gofr"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Index(c *gofr.Context) (interface{}, error) {
	return "bye, Take care", nil
}
