package norm

import "reflect"

// Table provides a hook for registration to figure out your table name. Embed
// it into a struct and set a table name with a norm tag. When you call Register(),
// norm will look for this field and parse the table name from the tag.
type Table uint8

const table Table = iota

func getTableField[X any](x X) reflect.StructField {
	for _, field := range fields(x) {
		if fieldType := field.Type; fieldType == typ(table) {
			return field
		}
	}

	panicf("could not find embedded model field for %s", typ(table).Name())

	return reflect.StructField{}
}
