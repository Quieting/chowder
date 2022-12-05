package copier

import (
	"reflect"
)

func parseStruct(typ reflect.Type) (fields []*field) {
	fields = make([]*field, 0, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		item := new(field)
		item.field = typ.Field(i)
		item.typ = typ.Field(i).Type
		item.offset = typ.Field(i).Offset
		fields = append(fields, item)
	}

	return fields
}
