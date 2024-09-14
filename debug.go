// debug.go has facilities for debugging.
// You can set NORM_DEBUG=1 to see generated SQL every time something calls
// Query().
package norm

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

var slogLevel = new(slog.LevelVar)

func init() {
	viper.MustBindEnv("DEBUG")

	h := slog.NewJSONHandler(
		os.Stderr,
		&slog.HandlerOptions{
			Level: slogLevel,
		},
	)

	slog.SetDefault(slog.New(h))

	if viper.GetBool("DEBUG") {
		slogLevel.Set(slog.LevelDebug)
	}
}
