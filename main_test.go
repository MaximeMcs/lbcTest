package main

import (
	_ "root/docs/ginsimple"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Test_fizzBuzzHandler(t *testing.T) {
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
			fizzBuzzHandler(tt.args.c)
		})
	}
}

func Test_queriesHandler(t *testing.T) {
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
			queriesHandler(tt.args.c)
		})
	}
}
