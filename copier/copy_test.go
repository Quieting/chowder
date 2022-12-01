package copier

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type A struct {
	age  int64
	name string
}
type B struct {
	num   int8
	age   int32
	name  string
	names []string
}

func Test_parse(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want []*field
	}{
		{
			name: "结构体测试",
			args: args{v: A{}},
			want: []*field{
				{Name: "age", Position: 0, Kind: Number, Len: 8},
				{Name: "name", Position: 8, Kind: String, Len: 8},
			},
		},
		{
			name: "结构体指针",
			args: args{
				v: A{},
			},
			want: []*field{
				{Name: "age", Position: 0, Kind: Number, Len: 8},
				{Name: "name", Position: 8, Kind: String, Len: 8},
			},
		},
		{
			name: "结构体内存对齐",
			args: args{
				v: B{},
			},
			want: []*field{
				{Name: "num", Position: 0, Kind: Number, Len: 1},
				{Name: "age", Position: 8, Kind: Number, Len: 8},
				{Name: "name", Position: 16, Kind: String, Len: 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseStruct(reflect.TypeOf(tt.args.v))
			for i, val := range got {
				if !reflect.DeepEqual(val, tt.want[i]) {
					t.Errorf("got = %v, want %v", val, tt.want[i])
				}
			}
		})
	}
}

func TestCopy(t *testing.T) {
	a := A{
		age:  20,
		name: "Jack",
	}
	b := B{
		age: 10,
	}
	fmt.Printf("a = %p, b=%p\n", &a, &b)
	fmt.Printf("%d\n", unsafe.Sizeof(&b.name))

	Copy(&a, &b)
	t.Logf("rsc: %+v, dst: %+v\n", a, b)
}

func TestSizeOf(t *testing.T) {
	type A struct {
		Age   int64
		Grade int64
		Names []string
	}
	a := A{}
	t.Logf("指针对象大小:%d\n", unsafe.Sizeof(&a))
	t.Logf("struct大小:%d\n", unsafe.Sizeof(a))

	t.Logf("string 大小:%d\n", unsafe.Sizeof(""))
	t.Logf("slice 大小:%d\n", unsafe.Sizeof([]int{}))
}

func Test_memove(t *testing.T) {
	type A struct {
		Int8    int8
		Int16   int16
		Strings []string
	}
	from, to := A{Int8: 8, Strings: []string{"jack", "tom"}}, A{}
	type args struct {
		from unsafe.Pointer
		to   unsafe.Pointer
		len  int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "int8 赋值",
			args: args{
				from: unsafe.Pointer(&(from.Int8)),
				to:   unsafe.Pointer(&(to.Int16)),
				len:  1,
			},
		},
		{
			name: "slice 赋值",
			args: args{
				from: unsafe.Pointer(&(from.Strings)),
				to:   unsafe.Pointer(&(to.Strings)),
				len:  24,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memove(tt.args.from, tt.args.to, tt.args.len)
			t.Logf("from: %+v, to:%+v", from, to)
		})
	}
}
