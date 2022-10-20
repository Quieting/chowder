package math

import (
	"reflect"
	"testing"
)

func TestMatrixMulti(t *testing.T) {
	type args struct {
		a [][]int64
		b [][]int64
	}
	tests := []struct {
		name string
		args args
		want [][]int64
	}{
		{
			name: "示例1:有效计算",
			args: args{
				a: [][]int64{
					{1, 2, 3},
					{2, 1, 3},
				},
				b: [][]int64{
					{2, 1},
					{3, 2},
					{1, 3},
				},
			},
			want: [][]int64{
				{11, 14},
				{10, 13},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixMulti(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixMulti() = %v, want %v", got, tt.want)
			}
		})
	}
}
