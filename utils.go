package norm

import (
	"fmt"
	"log/slog"
)

func panicf(template string, args ...any) {
	slog.Error(fmt.Sprintf(template, args...))
	panic(fmt.Sprintf(template, args...))
}
