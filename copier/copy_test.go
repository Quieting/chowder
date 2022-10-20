package copier

import (
	"reflect"
	"testing"
)

type A struct {
	age int64
}
type B struct {
	age int64
}

func Test_parseStruct(t *testing.T) {
	type args struct {
		v reflect.Type
	}
	tests := []struct {
		name string
		args args
		want typ
	}{
		{
			name: "示例1",
			args: args{v: reflect.TypeOf(A{})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseStruct(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		rsc interface{}
		dst interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "示例：int64类型赋值",
			args: args{
				rsc: A{
					age: 20,
				},
				dst: B{
					age: 10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.args.rsc, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("A:%+v, B:%+v", tt.args.rsc, tt.args.dst)
			t.Logf("A:%v, B:%v", tt.args.rsc, tt.args.dst)
		})
	}
}
