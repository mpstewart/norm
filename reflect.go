package norm

import "reflect"

func typ[X any](x X) reflect.Type {
	return reflect.TypeOf(x)
}

func fields[X any](x X) []reflect.StructField {
	return reflect.VisibleFields(typ(x))
}
