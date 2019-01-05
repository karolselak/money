package money

import (
	"reflect"
	"testing"
)

func TestNewWealth(t *testing.T) {
	tests := []struct {
		name string
		want *Wealth
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWealth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
