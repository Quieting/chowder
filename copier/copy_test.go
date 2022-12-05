package copier

import (
	"fmt"
	"testing"
	"unsafe"
)

type A struct {
	age   int64
	name  string
	grade int64
}
type B struct {
	num   int8
	age   int32
	grade *int64
	name  string
	names []string
}

func TestCopy(t *testing.T) {
	a := A{
		age:   20,
		name:  "Jack",
		grade: 10,
	}
	b := B{
		age:   10,
		grade: new(int64),
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
		from uintptr
		to   uintptr
		len  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "int8 赋值",
			args: args{
				from: uintptr(unsafe.Pointer(&(from.Int8))),
				to:   uintptr(unsafe.Pointer(&(to.Int16))),
				len:  1,
			},
		},
		{
			name: "slice 赋值",
			args: args{
				from: uintptr(unsafe.Pointer(&(from.Strings))),
				to:   uintptr(unsafe.Pointer(&(to.Strings))),
				len:  24,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// memove(tt.args.from, tt.args.to, tt.args.len)
			memove1(tt.args.from, tt.args.to, tt.args.len)
			t.Logf("from: %+v, to:%+v", from, to)
		})
	}
}
