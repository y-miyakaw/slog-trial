package logger

import (
	"log/slog"
	"os"
)

var GlobalLogger *slog.Logger

var programLevel = new(slog.LevelVar)

func init() {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		// remove the time key if there are no groups
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		// if a.Key == slog.SourceKey {
		// 	source := a.Value.Any().(*slog.Source)
		// 	source.File = fmt.Sprintf("%s:%d", source.File, source.Line)
		// 	source.Function = ""
		// 	source.Line = 0
		// }
		return a
	}

	level := new(slog.Level)

	if levelErr := level.UnmarshalText([]byte(os.Getenv("LOG_LEVEL"))); levelErr != nil {
		d := slog.LevelDebug
		level = &d
	}
	programLevel.Set(*level)

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       programLevel,
		ReplaceAttr: replace,
	})

	GlobalLogger = slog.New(handler)
}
