package copier

import (
	"reflect"
	"sync"
	"unsafe"

	"github.com/jinzhu/copier"
)

var caches = make(map[string]typ)
var mux sync.Mutex

type typ struct {
	typ    reflect.Type
	offset int          // 偏移量
	name   string       // 标识
	kind   reflect.Kind // 类型
	fields []typ        // 包含的字段
}

// 能够相互间赋值的类型映射
var types = map[reflect.Kind][]reflect.Kind{
	reflect.Int8:    {reflect.Int8, reflect.Uint8},
	reflect.Int16:   {reflect.Int8, reflect.Int16, reflect.Uint8, reflect.Uint16},
	reflect.Int32:   {reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint8, reflect.Uint16, reflect.Uint32},
	reflect.Int64:   {reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint},
	reflect.Int:     {reflect.Int, reflect.Uint},
	reflect.Uint8:   {reflect.Int8, reflect.Uint8},
	reflect.Uint16:  {reflect.Int8, reflect.Int16, reflect.Uint8, reflect.Uint16},
	reflect.Uint32:  {reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint8, reflect.Uint16, reflect.Uint32},
	reflect.Uint64:  {reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint},
	reflect.Uint:    {reflect.Int, reflect.Uint},
	reflect.Bool:    {reflect.Bool},
	reflect.Slice:   {reflect.Slice},
	reflect.Map:     {reflect.Map},
	reflect.Struct:  {reflect.Struct},
	reflect.Float32: {reflect.Float32},
	reflect.Float64: {reflect.Float32, reflect.Float64},
}

// 对应 runtime.iface 结构体
type iface struct {
	tab  uintptr // 对应运行时 iface.tab
	data unsafe.Pointer
}

func Copy(rsc, dst interface{}) error {
	return copier.Copy(dst, rsc)
}

func val(v interface{}) value {
	_t := parse(v)
	return value{
		addr: uintptr((*iface)(unsafe.Pointer(&v)).data),
		typ:  &_t,
	}
}

// parse 解析 struct、map类型
func parse(v interface{}) typ {
	t := reflect.TypeOf(v)
	var ty typ
	key := t.PkgPath() + "." + t.Name()
	mux.Lock()
	ty, ok := caches[key]
	if ok {
		mux.Unlock()
		return ty
	}
	mux.Unlock()

	switch t.Kind() {
	case reflect.Struct:
		ty = parseStruct(t)
	case reflect.Pointer:
		ty = parseStruct(t.Elem())
	default:

	}

	mux.Lock()
	caches[key] = ty
	mux.Unlock()

	return ty
}

func parseStruct(v reflect.Type) typ {
	//res := typ{
	//	rtyp:   v,
	//	fields: make([]filed, v.NumField()),
	//}
	//for i := 0; i < v.NumField(); i++ {
	//	f := v.Field(i)
	//	res.fields[i].name = f.Name
	//	res.fields[i].kind = f.Type.Kind()
	//	res.fields[i].offset = int(f.Offset)
	//}
	//return res
	return typ{}
}

func validType(rsc, dst reflect.Kind) bool {
	for _, kind := range types[rsc] {
		if kind == dst {
			return true
		}
	}
	return false
}

//// name 返回变量的偏移量
//func (t typ) find(s string) filed {
//	//for _, f := range t.fields {
//	//	if f.name != s {
//	//		continue
//	//	}
//	//	return f
//	//
//	//}
//	//return zerofiled
//
//	return field{}
//}

//func copier(rsc, dst value) error {
//	for _, r := range rsc.typ.fields {
//		f := dst.typ.find(r.name)
//		if f == zerofiled {
//			continue
//		}
//
//		for _, d := range dst.typ.fields {
//			if r.name != d.name {
//				continue
//			}
//			if !validType(r.kind, d.kind) {
//				continue
//			}
//			switch r.kind {
//			case reflect.Int64:
//				*(*int64)(unsafe.Pointer(dst.addr + uintptr(d.offset))) = *(*int64)(unsafe.Pointer(rsc.addr + uintptr(d.offset)))
//			}
//
//		}
//	}
//
//	return nil
//}

func assist(rsc, dst value) {
}
