package try

import (
	"reflect"
	"testing"
)

// 测试空切片是否可以使用下标赋值
// 结果：空切片不允许使用下标赋值，会抛出切片越界panic
func Test_Empty_Slice(t *testing.T) {
	var val []int
	val[0] = 10
}

func Test_openDoor(t *testing.T) {
	type args struct {
		d []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "示例1",
			args: args{
				d: []int64{
					3, 2, 1,
				},
			},
			want: []int64{
				2, 2, 2,
			},
		},
		{
			name: "示例1",
			args: args{
				d: []int64{
					2, 1, 3,
				},
			},
			want: []int64{
				2, 2, 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openDoor(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openDoor() = %v, want %v", got, tt.want)
			}
		})
	}
}
