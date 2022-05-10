package builtins

import (
	"fmt"
	"gosh/lib/history"
)

/// The builtins package contains default functions for the Gosh shell.
/// These commands checked in the main loop.

func Help() {
	fmt.Println("TBD!")
}

func History(h *history.GoshHistory, size int) []string {
	return h.Recents[:size+1]
}
