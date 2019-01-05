package main

import (
	"testing"

	"github.com/urfave/cli"
)

func TestCommand_info(t *testing.T) {
	type args struct {
		n   string
		usg string
		ali []string
	}
	tests := []struct {
		name string
		c    *Command
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.info(tt.args.n, tt.args.usg, tt.args.ali)
		})
	}
}

func TestCommand_action(t *testing.T) {
	tests := []struct {
		name string
		c    *Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.action()
		})
	}
}

func TestCommand_execute(t *testing.T) {
	type args struct {
		cntxt *cli.Context
	}
	tests := []struct {
		name    string
		c       *Command
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.execute(tt.args.cntxt); (err != nil) != tt.wantErr {
				t.Errorf("Command.execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
