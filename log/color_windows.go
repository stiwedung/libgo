package log

import (
	"syscall"

	"golang.org/x/sys/windows"
)

func init() {
	var origMode uint32
	stdout := windows.Handle(syscall.Stdout)
	windows.GetConsoleMode(stdout, &origMode)
	windows.SetConsoleMode(stdout, origMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
