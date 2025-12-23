package log

import (
	"log/slog"
	"os"
)

const (
	LevelTrace = slog.Level(-8)
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelFixMe = slog.Level(6)
	LevelError = slog.LevelError
	LevelFatal = slog.Level(12)
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	Level: LevelTrace, // TODO: SHOULD BE TAKEN FROM THE CONFIGS
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {

		// Customize the name of the level key and the output string, including
		// custom level values.
		if a.Key == slog.LevelKey {
			// Handle custom level values.
			level := a.Value.Any().(slog.Level)

			// This could also look up the name from a map or other structure, but
			// this demonstrates using a switch statement to rename levels. For
			// maximum performance, the string values should be constants, but this
			// example uses the raw strings for readability.
			switch {
			case level < LevelDebug:
				a.Value = slog.StringValue("TRACE")
			case level < LevelInfo:
				a.Value = slog.StringValue("DEBUG")
			case level < LevelWarn:
				a.Value = slog.StringValue("INFO")
			case level < LevelFixMe:
				a.Value = slog.StringValue("WARN")
			case level < LevelError:
				a.Value = slog.StringValue("FIXME")
			case level < LevelFatal:
				a.Value = slog.StringValue("ERROR")
			default:
				a.Value = slog.StringValue("FATAL")
			}
		}

		return a
	},
}))
