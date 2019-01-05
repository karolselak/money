package main

import (
	"reflect"
	"testing"
)

func TestApplication_info(t *testing.T) {
	tests := []struct {
		name string
		a    *Application
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.info()
		})
	}
}

func TestApplication_setLog(t *testing.T) {
	tests := []struct {
		name string
		a    *Application
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.setLog()
		})
	}
}

func TestApplication_init(t *testing.T) {
	tests := []struct {
		name string
		a    *Application
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.init()
		})
	}
}

func TestApplication_run(t *testing.T) {
	tests := []struct {
		name string
		a    *Application
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.run()
		})
	}
}

func TestApplication_cmdbasic(t *testing.T) {
	tests := []struct {
		name string
		a    *Application
		want *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.cmdbasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Application.cmdbasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApplication_register(t *testing.T) {
	tests := []struct {
		name string
		a    *Application
		want *[]Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.register(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Application.register() = %v, want %v", got, tt.want)
			}
		})
	}
}
