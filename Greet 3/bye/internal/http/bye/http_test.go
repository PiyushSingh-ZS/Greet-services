package bye

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/bmizerany/assert"
	"testing"
)

func TestBye(t *testing.T) {
	testcases := []struct {
		description string
		expOut      string
	}{
		{"success case", "bye, Take care"},
	}

	for i, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		H := New()
		out, _ := H.Index(ctx)
		assert.Equalf(t, tc.expOut, out, "TEST[%v],failed. body mismatch\n%s ", i, tc.description)

	}
}
