package main

import (
	"fmt"

	"github.com/arsmoriendy/de/pkgs/parsede"
)

func main() {
	opts := parseOpts()

	rstring := parsede.ParseDeFilesIn(
		opts.format,
		&opts.filters,
		opts.paths...,
	)

	fmt.Print(rstring)
}
