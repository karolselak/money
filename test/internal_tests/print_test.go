package base

import "testing"

func Test_prnt(t *testing.T) {
	type args struct {
		data [][]string
		cap  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prnt(tt.args.data, tt.args.cap)
		})
	}
}

func Test_prntTotal(t *testing.T) {
	type args struct {
		sum string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prntTotal(tt.args.sum)
		})
	}
}
