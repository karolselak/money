package util

import "testing"

func TestStf(t *testing.T) {
	type args struct {
		s string
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
			if got := Stf(tt.args.s); got != tt.want {
				t.Errorf("Stf() = %v, want %v", got, tt.want)
			}
		})
	}
}
