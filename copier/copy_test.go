package copier

import (
	"reflect"
	"testing"
	"unsafe"
)

type A struct {
	Int8   int8
	Int16  int16
	Int32  int32
	Int64  int64
	String string

	Int8Pointer   *int8
	Int16Pointer  *int16
	Int32Pointer  *int32
	Int64Pointer  *int64
	StringPointer *string
}
type B struct {
	Int8   int8
	Int16  int16
	Int32  int32
	Int64  int64
	String string
}
type C struct {
	Int8Pointer   *int8
	Int16Pointer  *int16
	Int32Pointer  *int32
	Int64Pointer  *int64
	StringPointer *string
}
type D struct {
	Int8   *int8
	Int16  *int16
	Int32  *int32
	Int64  *int64
	String *string
}

func TestCopy(t *testing.T) {
	type args struct {
		form, to interface{}
		want     interface{}
	}

	from := A{
		Int8:   8,
		Int16:  16,
		Int32:  32,
		Int64:  64,
		String: "show time",
	}
	from.Int8Pointer = &from.Int8
	from.Int16Pointer = &from.Int16
	from.Int32Pointer = &from.Int32
	from.Int64Pointer = &from.Int64
	from.StringPointer = &from.String

	tests := []struct {
		name string
		args args
	}{
		{
			name: "基础数据类型赋值",
			args: args{
				form: &from,
				to:   &B{},
				want: &B{
					Int8:   from.Int8,
					Int16:  from.Int16,
					Int32:  from.Int32,
					Int64:  from.Int64,
					String: from.String,
				},
			},
		},
		{
			name: "指针类型赋值",
			args: args{
				form: &from,
				to:   &C{},
				want: &C{
					Int8Pointer:   from.Int8Pointer,
					Int16Pointer:  from.Int16Pointer,
					Int32Pointer:  from.Int32Pointer,
					Int64Pointer:  from.Int64Pointer,
					StringPointer: from.StringPointer,
				},
			},
		},
		{
			name: "指针类型和非指针类型赋值",
			args: args{
				form: &from,
				to:   &D{},
				want: &D{
					Int8:   from.Int8Pointer,
					Int16:  from.Int16Pointer,
					Int32:  from.Int32Pointer,
					Int64:  from.Int64Pointer,
					String: from.StringPointer,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(tt.args.form, tt.args.to)
			if err != nil {
				t.Errorf("not want error")
			}
			if !reflect.DeepEqual(tt.args.want, tt.args.to) {
				t.Errorf("want: %+v, got: %+v\n", tt.args.want, tt.args.to)
			}
		})
	}
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
