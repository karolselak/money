package util

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		from   string
		to     string
		amount float64
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
			if got := Convert(tt.args.from, tt.args.to, tt.args.amount); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
