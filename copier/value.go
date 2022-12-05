package copier

import (
	"reflect"
	"unsafe"
)

type value struct {
	addr  unsafe.Pointer
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

// isNew： 指针变为nil时是否申请新的地址
func addr(val *value, isNew bool) uintptr {
	// todo：添加其他指针类型数据类型（扩展 reflect.Kind)
	if val.typ.Kind() != reflect.Pointer {
		return uintptr(val.addr)
	}

	point := val.addr
	val.addr = unsafe.Pointer(*(*uintptr)(val.addr)) // 指针对象取对象存储值（值是一个地址）
	val.typ = val.typ.Elem()
	if val.addr == nil && isNew {
		val.addr = unsafe.Pointer(reflect.New(val.typ).Pointer())
		*(*uintptr)(point) = uintptr(val.addr)
	}

	return addr(val, isNew)
}

func valueOf(v interface{}) (vals []*value) {
	typ := reflect.TypeOf(v)

	// 验证参数类型(结构体或者结构体指针)
	// todo：支持多级指针对象
	// todo：支持 slice、array、map 类型
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return
	}

	var start = reflect.ValueOf(v).Pointer() // 底层数据存储起始位置

	fields := get(typ)
	vals = make([]*value, 0, len(fields))
	for _, f := range fields {
		val := &value{
			addr:  unsafe.Pointer(start + f.offset),
			typ:   f.typ,
			field: f,
		}

		vals = append(vals, val)
	}

	return
}
