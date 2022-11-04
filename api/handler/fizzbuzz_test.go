package handler

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_saveQuery(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := saveQuery(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("saveQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("saveQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		a     int
		b     int
		limit int
		str1  string
		str2  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.a, tt.args.b, tt.args.limit, tt.args.str1, tt.args.str2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFizzBuzzHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FizzBuzzHandler(tt.args.c)
		})
	}
}
