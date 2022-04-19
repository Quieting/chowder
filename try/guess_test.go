package try

import (
	"reflect"
	"testing"
)

func Test_arrangement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "示例1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{3, 2, 1},
				{2, 3, 1},
				{2, 1, 3},
				{3, 1, 2},
				{1, 3, 2},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrangement() = %v, want %v", got, tt.want)
			}
		})
	}
}
