// search.go contains utilities for performing searches
package norm

import (
	"fmt"
	"log/slog"
	"strings"
)

// Search is a type-parameterized map that facilitates searching for a model
// X in a database. See Query() for how to turn this into an actual query.
type Search[X any] map[string]any

// Query compiles a SQL query and list of bindvars from the invocant Search[X].
// The ordering of fields in X determines the ordering of the criteria in the
// final SQL query.
func (s Search[X]) Query() (string, []any) {
	var x X

	sb := new(strings.Builder)

	selects := getSelects(x)
	fmt.Fprintf(sb,
		"SELECT %s FROM %s WHERE ",
		strings.Join(selects, ", "),
		getSource(x),
	)

	criteria := make([]string, 0, len(s))
	bindvars := make([]any, 0, len(s))

	// for all of our selectable fields, check to see if we have criteria for
	// them; if we do, append them
	for _, field := range selects {
		if value, ok := s[field]; ok {
			criteria = append(criteria, fmt.Sprintf("%s = ?", field))
			bindvars = append(bindvars, value)
		} else {
			continue
		}
	}

	fmt.Fprintf(sb, strings.Join(criteria, " AND "))

	query := sb.String()

	slog.Debug(
		fmt.Sprintf("Search[%s].Query()", typ(x).Name()),
		slog.String("query", query),
		slog.Any("bindvars", bindvars),
	)

	return query, bindvars
}
