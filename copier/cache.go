package copier

import (
	"reflect"
	"sync"
)

var typeCache sync.Map

// get 返回当前结构体类型包含字段
func get(typ reflect.Type) []*field {
	if typ.Kind() != reflect.Struct {
		return nil
	}

	val, ok := typeCache.Load(typ)
	if !ok {
		fields := parseStruct(typ)
		val = fields
		typeCache.Store(typ, fields)
	}

	return val.([]*field)
}
