package copier

import (
	"reflect"
	"unsafe"
)

var intSize = int(unsafe.Sizeof(0))
var ptrSize = int(unsafe.Sizeof(new(int)))

type value struct {
	addr  unsafe.Pointer
	typ   reflect.Type // addr 位置存储的类型
	field *field
}

func (v *value) size() int {
	switch v.typ.Kind() {
	case reflect.Int8, reflect.Uint8, reflect.Bool:
		return 1
	case reflect.Int16, reflect.Uint16:
		return 2
	case reflect.Int32, reflect.Uint32, reflect.Float32:
		return 4
	case reflect.Uint64, reflect.Int64, reflect.Float64, reflect.Complex64:
		return 8
	case reflect.Complex128:
		return 16
	case reflect.Int, reflect.Uint:
		return intSize
	case reflect.Pointer, reflect.Uintptr, reflect.Array, reflect.Map:
		return ptrSize
	case reflect.String: // todo: 考虑copy底层数组？
		return ptrSize + intSize
	case reflect.Slice:
		return ptrSize + intSize + intSize
	default: // 暂不支持的数据统一返回0
		return 0
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
	var start uintptr // 底层数据存储起始位置
	typ := reflect.TypeOf(v)

	// 验证参数类型(结构体或者结构体指针)
	// todo：支持多级指针对象
	// todo：支持 slice、array、map 类型
	if typ.Kind() == reflect.Pointer {
		start = reflect.ValueOf(v).Pointer()
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return
	}

	if start == 0 {
		return
	}

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
