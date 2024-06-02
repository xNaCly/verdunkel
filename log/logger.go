package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

type Handler struct {
	h slog.Handler
}

func New(w io.Writer, lvl slog.Level) *Handler {
	return &Handler{
		h: slog.NewTextHandler(w, &slog.HandlerOptions{
			Level:       lvl,
			ReplaceAttr: suppressDefaults(nil),
		}),
	}
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{h: h.h.WithAttrs(attrs)}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{h: h.h.WithGroup(name)}
}

func suppressDefaults(
	next func([]string, slog.Attr) slog.Attr,
) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey ||
			a.Key == slog.LevelKey ||
			a.Key == slog.MessageKey {
			return slog.Attr{}
		}
		if next == nil {
			return a
		}
		return next(groups, a)
	}
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()

	prefix := ""
	postfix := ""
	switch r.Level {
	case slog.LevelDebug:
		level = "DBG"
		prefix = BG_CYAN
		postfix = CYAN
	case slog.LevelInfo:
		level = "INF"
		prefix = BG_BRIGHT_BLUE
		postfix = BRIGHT_BLUE
	case slog.LevelWarn:
		level = "WRN"
		prefix = BG_BRIGHT_YELLOW
		postfix = BRIGHT_YELLOW
	case slog.LevelError:
		level = "ERR"
		prefix = BG_BRIGHT_RED
		postfix = BRIGHT_RED
	}

	fmt.Print(
		GRAY,
		r.Time.Format("[15:04:05.000]"),
		RESET,
		" ",
		BLACK,
		prefix,
		" ",
		level,
		" ",
		RESET,
		" ",
		postfix,
		r.Message,
		RESET,
	)

	if r.NumAttrs() > 0 {
		fmt.Print(" ", GRAY)
		h.h.Handle(ctx, r)
		fmt.Print(RESET)
	} else {
		fmt.Print("\n")
	}

	return nil
}
