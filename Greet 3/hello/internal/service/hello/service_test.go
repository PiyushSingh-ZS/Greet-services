package service

import (
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	app := gofr.New()
	tests := []struct {
		desc   string
		input  string
		output interface{}
		expErr error
	}{
		{"Success", "chakradhar", "hello, chakradhar", nil},
		{"Success", "", "hello", nil},
	}
	for i, val := range tests {

		ctx := gofr.NewContext(nil, nil, app)

		svc := New()
		got, err := svc.Greet(ctx, val.input)

		assert.Equalf(t, val.output, got, "TEST[%d], failed.\n%s", i, val.desc)
		assert.Equalf(t, val.expErr, err, "TEST[%d], failed.\n%s", i, val.desc)
	}

}
