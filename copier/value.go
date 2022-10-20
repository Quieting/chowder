package copier

import (
	"errors"
	"reflect"
)

var ErrMismatchType = errors.New("类型不匹配")

type Value interface {
	Type() reflect.Type
	Set(n Value)
	Addr() uintptr
}

type value struct {
	addr uintptr // 存储地址
	typ  *typ    // 类型
}

func (v *value) Type() reflect.Type {
	return v.typ.typ
}

func (v *value) Addr() uintptr {
	return v.addr
}

func (v *value) Set(n Value) {
	return
}

type number struct {
	value
	val []byte
}

var _ Value = new(number)

func (v *number) Set(val Value) {
	n, ok := val.(*number)
	if !ok {
		return
	}

	copy(n.val, v.val)

	return
}
