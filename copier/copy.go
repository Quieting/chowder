package copier

import (
	"reflect"
	"unsafe"
)

type Kind int

const (
	Invalid Kind = 1 << iota
	Number
	Float
	Bool
	String
	Struct
	Point
)
const (
	IndirectKind = Number | Float | Bool | String
)

func Copy(rsc, dst interface{}) (err error) {
	rscVal := valueOf(rsc)
	dstVal := valueOf(dst)

	for _, dstField := range dstVal {
		for _, rscFild := range rscVal {
			if !canCopy(dstField, rscFild) {
				continue
			}

			if dstField.typ.Kind() == reflect.Struct {
				copyForstruct(dstField, rscFild)
			}

			if dstField.field.indirectKind()&IndirectKind > 0 { // 基础数据类型值copy
				memove(addr(rscFild, false), addr(dstField, true), minOffset(dstField, rscFild))
			}
			break
		}

	}

	return
}

func canCopy(rsc, dst *value) bool {
	if rsc.field.kind() == Invalid || rsc.field.indirectKind() == Invalid ||
		dst.field.kind() == Invalid || dst.field.indirectKind() == IndirectKind {
		return false
	}

	if rsc.field.indirectKind() != dst.field.indirectKind() { // 底层数据类型一致
		return false
	}

	return rsc.field.name() == dst.field.name()
}

func minOffset(a, b *value) int {
	if a.size() > b.size() {
		return a.size()
	}
	return b.size()
}

func memove(from, to uintptr, size int) {
	if from == 0 || to == 0 {
		return
	}
	switch size {
	case 1:
		var b *[1]byte
		b = (*[1]byte)(unsafe.Pointer(from))
		*(*[1]byte)(unsafe.Pointer(to)) = *b
	case 2:
		var b *[2]byte
		b = (*[2]byte)(unsafe.Pointer(from))
		*(*[2]byte)(unsafe.Pointer(to)) = *b
	case 4:
		var b *[4]byte
		b = (*[4]byte)(unsafe.Pointer(from))
		*(*[4]byte)(unsafe.Pointer(to)) = *b
	case 8:
		var b *[8]byte
		b = (*[8]byte)(unsafe.Pointer(from))
		*(*[8]byte)(unsafe.Pointer(to)) = *b
	case 16:
		var b *[16]byte
		b = (*[16]byte)(unsafe.Pointer(from))
		*(*[16]byte)(unsafe.Pointer(to)) = *b
	case 24:
		var b *[24]byte
		b = (*[24]byte)(unsafe.Pointer(from))
		*(*[24]byte)(unsafe.Pointer(to)) = *b
	default:
		memove1(from, to, size)
	}
}

func memove1(from, to uintptr, size int) {
	writeMemory(to, readMemory(from, size))
}
func convert(t reflect.Type) Kind {
	switch t.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return Number
	case reflect.Float32, reflect.Float64:
		return Float
	case reflect.Pointer:
		return convert(t.Elem())
	case reflect.String:
		return String
	case reflect.Bool:
		return Bool
	case reflect.Struct:
		return Struct
	default:
		return Invalid
	}
}

func readMemory(addr uintptr, size int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: addr,
		Len:  size,
		Cap:  size,
	}))
}

func writeMemory(addr uintptr, d []byte) {
	to := (*[]byte)(unsafe.Pointer(
		&reflect.SliceHeader{
			Data: addr,
			Len:  len(d),
			Cap:  len(d),
		}))
	copy(*to, d)
}

func copyForstruct(dst, rsc *value) {
	dstVal := valueOf(dst)
	rscVal := valueOf(rsc)
	for _, dstField := range dstVal {
		for _, rscFild := range rscVal {
			if !canCopy(dstField, rscFild) {
				continue
			}

			if dstField.typ.Kind() == reflect.Struct {
				copyForstruct(dstField, rscFild)
			}

			if dstField.field.indirectKind()&IndirectKind > 0 { // 基础数据类型值copy
				memove(addr(rscFild, false), addr(dstField, true), minOffset(dstField, rscFild))
			}
			break
		}
	}
}
