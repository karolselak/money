package util

import "testing"

func TestGetPrice(t *testing.T) {
	type args struct {
		sym string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPrice(tt.args.sym); got != tt.want {
				t.Errorf("GetPrice() = %v, want %v", got, tt.want)
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
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPriceV2(tt.args.name); got != tt.want {
				t.Errorf("GetPriceV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
