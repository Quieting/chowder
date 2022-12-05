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
	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64
	String string
	Bool   bool

	Int8Pointer   *int8
	Int16Pointer  *int16
	Int32Pointer  *int32
	Int64Pointer  *int64
	Uint8Pointer  *uint8
	Uint16Pointer *uint16
	Uint32Pointer *uint32
	Uint64Pointer *uint64
	StringPointer *string
	BoolPointer   *bool

	B B
}
type B struct {
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64

	String string
	Bool   bool
}
type C struct {
	Int8Pointer   *int8
	Int16Pointer  *int16
	Int32Pointer  *int32
	Int64Pointer  *int64
	Uint8Pointer  *uint8
	Uint16Pointer *uint16
	Uint32Pointer *uint32
	Uint64Pointer *uint64
	StringPointer *string
	BoolPointer   *bool
}
type D struct {
	Int8   *int8
	Int16  *int16
	Int32  *int32
	Int64  *int64
	Uint8  *uint8
	Uint16 *uint16
	Uint32 *uint32
	Uint64 *uint64
	String *string
	Bool   *bool
}

type E struct {
	B B
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
		Uint8:  8,
		Uint16: 16,
		Uint32: 32,
		Uint64: 64,
		String: "show time",
		Bool:   true,
		B: B{
			Int8:   8,
			Int16:  16,
			Int32:  32,
			Int64:  64,
			Uint8:  8,
			Uint16: 16,
			Uint32: 32,
			Uint64: 64,
			String: "show time",
			Bool:   true,
		},
	}
	from.Int8Pointer = &from.Int8
	from.Int16Pointer = &from.Int16
	from.Int32Pointer = &from.Int32
	from.Int64Pointer = &from.Int64
	from.Uint8Pointer = &from.Uint8
	from.Uint16Pointer = &from.Uint16
	from.Uint32Pointer = &from.Uint32
	from.Uint64Pointer = &from.Uint64
	from.StringPointer = &from.String
	from.BoolPointer = &from.Bool

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
					Uint8:  from.Uint8,
					Uint16: from.Uint16,
					Uint32: from.Uint32,
					Uint64: from.Uint64,
					String: from.String,
					Bool:   from.Bool,
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
					Uint8Pointer:  from.Uint8Pointer,
					Uint16Pointer: from.Uint16Pointer,
					Uint32Pointer: from.Uint32Pointer,
					Uint64Pointer: from.Uint64Pointer,
					StringPointer: from.StringPointer,
					BoolPointer:   from.BoolPointer,
				},
			},
		},
		{
			name: "指针类型和非指针类型赋值",
			args: args{
				form: &from,
				to:   &D{},
				want: &D{
					Int8:   &from.Int8,
					Int16:  &from.Int16,
					Int32:  &from.Int32,
					Int64:  &from.Int64,
					Uint8:  &from.Uint8,
					Uint16: &from.Uint16,
					Uint32: &from.Uint32,
					Uint64: &from.Uint64,
					String: &from.String,
					Bool:   &from.Bool,
				},
			},
		},
		{
			name: "结构体字段赋值",
			args: args{
				form: &from,
				to:   &E{},
				want: &E{
					B: from.B,
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
