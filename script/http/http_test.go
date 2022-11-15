package http

import (
	"net/url"
	"reflect"
	"testing"
)

func TestPathsValues(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want url.Values
	}{
		{
			name: "匿名结构体",
			args: args{
				v: struct {
					Name string `form:"name"`
				}{
					Name: "栎阳",
				},
			},
			want: url.Values{
				"name": []string{"栎阳"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathsValues(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathsValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
