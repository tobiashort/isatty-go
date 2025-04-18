package isatty

import "golang.org/x/term"

func IsTerminal(fd int) bool {
	return term.IsTerminal(fd)
}
