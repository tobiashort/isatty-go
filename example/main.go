package main

import (
	"fmt"

	"github.com/tobiashort/isatty-go"
)

func main() {
	fmt.Println(isatty.IsTerminal())
}
