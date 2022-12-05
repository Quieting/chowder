package copier

import (
	"reflect"
	"sync"
)

var typeCache sync.Map

func get(typ reflect.Type) []*field {
	val, ok := typeCache.Load(typ)
	if !ok {
		fields := parseStruct(typ)
		val = fields
		typeCache.Store(typ, fields)
	}

	return val.([]*field)
}
