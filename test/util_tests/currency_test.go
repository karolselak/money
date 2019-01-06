package utiltest

import (
	"testing"

	"github.com/mohfunk/money/pkg/util"
)

func TestConvert(t *testing.T) {
	type args struct {
		from   string
		to     string
		amount float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"usd->cad", args{"USD", "CAD", 10}},
		{"usd->cad", args{"CAD", "USD", 10}},
		{"usd->cad", args{"USD", "SAR", 10}},
		{"usd->cad", args{"SAR", "USD", 10}},
		{"usd->cad", args{"CAD", "SAR", 10}},
		{"usd->cad", args{"SAR", "CAD", 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.Convert(tt.args.from, tt.args.to, tt.args.amount); got == 0 {
				t.Errorf("Convert() = %v, want %v", got, 0)
			}
		})
	}
}
