package copier

import (
	"reflect"
	"unsafe"
)

func Copy(rsc, dst interface{}) (err error) {
	rscVal := valueOf(rsc)
	dstVal := valueOf(dst)
	for _, dstField := range dstVal.fields {
		rscField, ok := rscVal.find(dstField.Name)
		if !ok {
			continue
		}

		if !kindIsMatch(dstField.Kind, rscField.Kind) {
			continue
		}
		memove(unsafe.Pointer(rscField.Position+rscVal.addr), unsafe.Pointer(dstField.Position+dstVal.addr), minOffset(dstField, rscField))
	}

	return
}

func kindIsMatch(a, b Kind) bool {
	return a == b
}
func minOffset(a, b *field) int64 {
	if a.Len > b.Len {
		return a.Len
	}
	return b.Len
}

func valueOf(v interface{}) *structVal {
	val := reflect.ValueOf(v)
	rscPosition := getPosition(v)
	return &structVal{
		addr:   val.Pointer(),
		fields: rscPosition.Fields,
	}
}

// find 返回 name 字段的偏移量
func (p *structVal) find(name string) (*field, bool) {
	for _, val := range p.fields {
		if val.Name != name {
			continue
		}
		return val, true
	}
	return nil, false
}

func getPosition(v interface{}) *structPosition {
	rscTyp := reflect.TypeOf(v)

	// 验证参数类型(结构体或者结构体指针)
	if rscTyp.Kind() == reflect.Pointer {
		rscTyp = rscTyp.Elem()
	}
	if rscTyp.Kind() != reflect.Struct {
		return nil
	}

	// todo:并发安全
	rscPotion, ok := typeCache[rscTyp]
	if !ok {
		rscPotion = &structPosition{
			Fields: parseStruct(rscTyp),
		}
		typeCache[rscTyp] = rscPotion
	}

	return rscPotion
}

type Kind int

const (
	Invalid Kind = iota
	Number
	Float
	String
	Point
)

var typeCache = make(map[reflect.Type]*structPosition)

// 结构体相对
type structPosition struct {
	Fields []*field
}

type structVal struct {
	addr   uintptr
	fields []*field
}

type field struct {
	Name     string  // 结构体字段赋值标识
	Position uintptr // 相对地址
	Kind     Kind    // 类型
	Len      int64
}

// 结构体之间的基础类型的之间值拷贝，不考虑嵌套结构体
// 解析每个结构体字段的相对位置
func parseStruct(typ reflect.Type) (fields []*field) {
	fields = make([]*field, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		item := new(field)
		item.Name = typ.Field(i).Name
		item.Position = typ.Field(i).Offset
		item.Kind = convert(typ.Field(i).Type)

		// todo: 完善每种类型所占长度
		switch typ.Field(i).Type.Kind() {
		case reflect.Int8, reflect.Uint8:
			item.Len = 1
		case reflect.Int16, reflect.Uint16:
			item.Len = 2
		case reflect.Int32, reflect.Uint32, reflect.Float32:
			item.Len = 4
		case reflect.Int, reflect.Uint, reflect.Pointer:
			item.Len = int64(unsafe.Sizeof(0))
		case reflect.String:
			item.Len = int64(unsafe.Sizeof("s"))
		case reflect.Complex128:
			item.Len = 16
		case reflect.Slice:
			item.Len = 24
		default:
			item.Len = 8
		}

		fields = append(fields, item)
	}

	return fields
}

func memove(from, to unsafe.Pointer, len int64) {
	switch len {
	case 1:
		var b *[1]byte
		b = (*[1]byte)(from)
		*(*[1]byte)(to) = *b
	case 2:
		var b *[2]byte
		b = (*[2]byte)(from)
		*(*[2]byte)(to) = *b
	case 4:
		var b *[4]byte
		b = (*[4]byte)(from)
		*(*[4]byte)(to) = *b
	case 8:
		var b *[8]byte
		b = (*[8]byte)(from)
		*(*[8]byte)(to) = *b
	case 16:
		var b *[16]byte
		b = (*[16]byte)(from)
		*(*[16]byte)(to) = *b
	case 24:
		var b *[24]byte
		b = (*[24]byte)(from)
		*(*[24]byte)(to) = *b
	}
}

func readMemory(addr uintptr, size int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: addr,
		Len:  size,
		Cap:  size,
	}))
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
	default:
		return Invalid
	}
}
