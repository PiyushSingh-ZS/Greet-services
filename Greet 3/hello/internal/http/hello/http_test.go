package http

import (
	"awesomeProject/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr/request"

	"github.com/golang/mock/gomock"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	app := gofr.New()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceMock := service.NewMockService(ctrl)
	tests := []struct {
		desc   string
		path   string
		input  string
		output interface{}
		expErr error
		calls  []*gomock.Call
	}{

		{"Success", "/hello?name=chakradhar", "chakradhar", "hello, chakradhar", nil, []*gomock.Call{
			serviceMock.EXPECT().Greet(gomock.AssignableToTypeOf(&gofr.Context{}), "chakradhar").Return("hello, chakradhar", nil),
		}},
		{"Success", "/hello", "", "hello", nil, []*gomock.Call{
			serviceMock.EXPECT().Greet(gomock.AssignableToTypeOf(&gofr.Context{}), "").Return("hello", nil),
		}},
	}

	for i, val := range tests {
		r := httptest.NewRequest(http.MethodGet, val.path, http.NoBody)
		req := request.NewHTTPRequest(r)
		ctx := gofr.NewContext(nil, req, app)

		h := New(serviceMock)
		got, err := h.Index(ctx)

		assert.Equalf(t, val.output, got, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}
}
