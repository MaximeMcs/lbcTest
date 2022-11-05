package handler

import (
	"reflect"
	"testing"
)

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
		{
			name: "Test fizzBuzz",
			args: args{
				a:     2,
				b:     3,
				limit: 6,
				str1:  "test",
				str2:  "toto",
			},
			want: []string{"1", "test", "toto", "test", "5", "testtoto"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.a, tt.args.b, tt.args.limit, tt.args.str1, tt.args.str2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
