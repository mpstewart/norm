// sources.go contains utilities for registering sources for norm
package norm

import (
	"reflect"
	"sync"
)

var (
	mu      = new(sync.Mutex)
	sources = make(map[reflect.Type]string)
)

// Register registers a given type X with norm's internal table mapping cache.
// X must embed norm.Table, and that field must have a norm tag specifying the
// table name that backs X in the database. If either of those are missing,
// Register will panic.
func Register[X any]() {
	var x X
	xtype := reflect.TypeOf(x)

	modelField := getTableField(x)
	if tableName, ok := modelField.Tag.Lookup("norm"); !ok {
		panicf("no norm struct tag model embed for %s", xtype.Name())
	} else {
		mu.Lock()
		defer mu.Unlock()

		sources[xtype] = tableName
	}
}

func getSource[X any](x X) string {
	xtype := reflect.TypeOf(x)
	mu.Lock()
	defer mu.Unlock()

	source, known := sources[xtype]
	if !known {
		panicf("do not know source of %s", xtype.Name())
	}

	return source
}
