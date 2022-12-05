package copier

import (
	"reflect"
)

type field struct {
	offset uintptr      // 相对偏移量
	typ    reflect.Type // 类型

	field reflect.StructField
}

// indirectKind 返回最底层数据类型(解指针后的数据类型)
func (f *field) indirectKind() Kind {
	if kind := f.kind(); kind != Point {
		return kind
	}

	return convert(elem(f.typ))
}

func (f *field) kind() Kind {
	return convert(f.typ)
}

func (f *field) name() string {
	return f.field.Name
}

func elem(typ reflect.Type) reflect.Type {
	if typ.Kind() == reflect.Pointer {
		return elem(typ.Elem())
	}
	return typ
}
