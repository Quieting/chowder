package match

import (
	"reflect"
	"testing"
)

type RoomInfo struct {
	Id        int
	PeopleNum int
}

func (r RoomInfo) Number() int {
	return r.PeopleNum
}
func TestMatch(t *testing.T) {

	type args struct {
		data []Matcher
	}
	tests := []struct {
		name string
		args args
		want [][]RoomInfo
	}{
		{
			name: "示例1",
			args: args{
				data: []Matcher{
					RoomInfo{1, 1},
					RoomInfo{2, 5},
				},
			},
			want: [][]RoomInfo{
				{{2, 5}, {1, 1}},
			},
		},
		{
			name: "示例2",
			args: args{
				data: []Matcher{
					RoomInfo{1, 1},
					RoomInfo{3, 1},
					RoomInfo{2, 4},
				},
			},
			want: [][]RoomInfo{
				{RoomInfo{2, 4}, RoomInfo{1, 1}, RoomInfo{3, 1}},
			},
		},

		{
			name: "示例3",
			args: args{
				data: []Matcher{
					RoomInfo{0, 1},
					RoomInfo{0, 1},
					RoomInfo{0, 4},
					RoomInfo{0, 3},
					RoomInfo{0, 3},
					RoomInfo{0, 3},
					RoomInfo{0, 5},
					RoomInfo{0, 1},
					RoomInfo{0, 4},
					RoomInfo{0, 2},
				},
			},
			want: [][]RoomInfo{
				{{0, 3}, {0, 3}},
				{{0, 5}, {0, 1}},
				{{0, 4}, {0, 1}, {0, 1}},
				{{0, 4}, {0, 2}},
			},
		},

		{
			name: "示例4",
			args: args{
				data: []Matcher{
					RoomInfo{0, 1},
					RoomInfo{0, 2},
					RoomInfo{0, 2},
					RoomInfo{0, 2},
					RoomInfo{0, 2},
					RoomInfo{0, 4},
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
			if got := Match(tt.args.data, 6); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
