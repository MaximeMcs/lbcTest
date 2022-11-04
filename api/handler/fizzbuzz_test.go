package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFizzBuzzHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		name: "Test FizzBuzz Query",
		args: args{
			int1:  "3",
			int2:  "5",
			limit: "15",
			str1:  "toto",
			str2:  "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FizzBuzzHandler(tt.args.c)
		})
	}
}
