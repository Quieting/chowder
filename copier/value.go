package copier

import (
	"fmt"
	"reflect"
	"unsafe"
)

type value struct {
	addr  uintptr
	typ   reflect.Type // addr 位置存储的类型
	field *field
}

func (v *value) size() int {
	// todo: 完善每种类型所占长度
	switch v.typ.Kind() {
	case reflect.Int8, reflect.Uint8:
		return 1
	case reflect.Int16, reflect.Uint16:
		return 2
	case reflect.Int32, reflect.Uint32, reflect.Float32:
		return 4
	case reflect.Int, reflect.Uint, reflect.Pointer:
		return int(unsafe.Sizeof(0))
	case reflect.String:
		return int(unsafe.Sizeof(""))
	case reflect.Complex128:
		return 16
	case reflect.Slice:
		return 24
	default:
		return 8
	}
}

func addr(val *value) uintptr {
	// todo：添加其他指针类型数据类型（扩展 reflect.Kind)
	if val.typ.Kind() != reflect.Pointer {
		return val.addr
	}

	val.addr = *(*uintptr)(unsafe.Pointer(val.addr)) // 指针对象取对象存储值（值是一个地址）
	val.typ = val.typ.Elem()

	return addr(val)
}

func valueOf(v interface{}) (vals []*value) {
	typ := reflect.TypeOf(v)

	// 验证参数类型(结构体或者结构体指针)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return
	}
	var start = reflect.ValueOf(v).Pointer()

	fields := get(typ)
	vals = make([]*value, 0, len(fields))
	for _, f := range fields {
		val := &value{
			addr:  start + f.offset,
			typ:   f.typ,
			field: f,
		}
		fmt.Println(f.name(), (uintptr)(unsafe.Pointer(val.addr)))
		vals = append(vals, val)
	}

	return
}
