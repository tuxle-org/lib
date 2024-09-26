package internal

import (
	"log/slog"
	"os"
)

func Assert(condition bool, errMessage string) {
	if condition {
		return
	}

	slog.Error("ASSERT FAILED", "err", errMessage)
	os.Exit(2)
}
