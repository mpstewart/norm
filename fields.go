package norm

func getSelects[X any](x X) []string {
	fs := fields(x)
	selects := make([]string, 0, len(fs))

	for _, field := range fs {
		if field.Type == typ(table) {
			continue
		}
		if sel, ok := field.Tag.Lookup("norm"); ok {
			selects = append(selects, sel)
		} else {
			continue
		}
	}

	return selects
}
