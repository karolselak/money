package utiltest

import (
	"testing"

	"github.com/mohfunk/money/pkg/util"
)

func TestGetPrice(t *testing.T) {
	type args struct {
		sym string
	}
	tests := []struct {
		name string
		args args
	}{
		{"btc-test", args{"bitcoin"}},
		{"dcr-test", args{"decred"}},
		{"xrp-test", args{"ripple"}},
		{"eos-test", args{"eos"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.GetPrice(tt.args.sym); got == 0 {
				t.Errorf("GetPrice() = %v, want %v", got, 0)
			}
		})
	}
}

func TestGetPriceV2(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{

		{"btc-test", args{"bitcoin"}},
		{"dcr-test", args{"decred"}},
		{"xrp-test", args{"ripple"}},
		{"eos-test", args{"eos"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.GetPriceV2(tt.args.name); got == 0 {
				t.Errorf("GetPriceV2() = %v, want %v", got, 0)
			}
		})
	}
}
