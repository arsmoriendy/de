package main

import (
	"fmt"

	"github.com/arsmoriendy/de/pkgs/parsede"
)

const usage = `Usage:
        de <format>

        <format>:
                Output format

                E.g. "Name: {Name}" would result in "Name: <Desktop Entry Name>"
`

func main() {
	opts := parseOpts()

	rstring := parsede.ParseDeFilesIn(
		opts.format,
		&opts.filters,
		opts.paths...,
	)

	fmt.Print(rstring)
}
