package main

import (
	"fmt"

	"github.com/arsmoriendy/de/pkgs/parsede"
	"github.com/arsmoriendy/de/pkgs/parser"
)

func main() {
	opts := parser.ParseOpts()

	rstring := parsede.ParseDeFiles(
		opts.Format,
		&opts.Filters,
		opts.Paths...,
	)

	fmt.Print(rstring)
}
