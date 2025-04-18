package main

import (
	"fmt"
	"os"

	"github.com/tobiashort/isatty"
)

func main() {
	fmt.Println(isatty.IsTerminal(os.Stdout))
}
