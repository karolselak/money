package base

import (
	"testing"

	money "github.com/mohfunk/money/internal"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func TestList(t *testing.T) {
	type args struct {
		w   *money.Wealth
		log *logrus.Logger
		c   *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.w, tt.args.log, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listFiat(t *testing.T) {
	type args struct {
		w *money.Wealth
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listFiat(tt.args.w)
		})
	}
}

func Test_listCrypto(t *testing.T) {
	type args struct {
		w *money.Wealth
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listCrypto(tt.args.w)
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		w   *money.Wealth
		log *logrus.Logger
		c   *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Update(tt.args.w, tt.args.log, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		w   *money.Wealth
		log *logrus.Logger
		c   *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.w, tt.args.log, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModify(t *testing.T) {
	type args struct {
		w   *money.Wealth
		log *logrus.Logger
		c   *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Modify(tt.args.w, tt.args.log, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Modify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Modify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cmcAPI(t *testing.T) {
	type args struct {
		sym string
		c   chan float64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmcAPI(tt.args.sym, tt.args.c)
		})
	}
}
