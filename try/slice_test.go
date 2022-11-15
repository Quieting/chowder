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

func TestMatch(t *testing.T) {
	type args struct {
		data []RoomInfo
	}
	tests := []struct {
		name string
		args args
		want [][]RoomInfo
	}{
		{
			name: "示例1",
			args: args{
				data: []RoomInfo{
					{1, 1},
					{2, 5},
				},
			},
			want: [][]RoomInfo{
				{{2, 5}, {1, 1}},
			},
		},
		{
			name: "示例1",
			args: args{
				data: []RoomInfo{
					{1, 1},
					{3, 1},
					{2, 4},
				},
			},
			want: [][]RoomInfo{
				{{2, 4}, {1, 1}, {3, 1}},
			},
		},

		{
			name: "示例2",
			args: args{
				data: []RoomInfo{
					{0, 1},
					{0, 1},
					{0, 4},
					{0, 3},
					{0, 3},
					{0, 3},
					{0, 5},
					{0, 1},
					{0, 4},
					{0, 2},
				},
			},
			want: [][]RoomInfo{
				{{0, 3}, {0, 3}},
				{{0, 5}, {0, 1}},
				{{0, 4}, {0, 1}, {0, 1}},
			},
		},

		{
			name: "示例1",
			args: args{
				data: []RoomInfo{
					{0, 1},
					{0, 2},
					{0, 2},
					{0, 2},
					{0, 2},
					{0, 4},
				},
			},
			want: [][]RoomInfo{
				{{0, 4}, {0, 2}},
				{{0, 2}, {0, 2}, {0, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Match(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
