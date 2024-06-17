package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/arsmoriendy/de/pkgs/getdefiles"
)

const usage = `Usage:
        de <format>

        <format>:
                Output format

                E.g. "Name: {Name}" would result in "Name: <Desktop Entry Name>"
`

func main() {
	home := path.Join(os.Getenv("HOME"), ".local/share/applications/")

	if len(os.Args) < 2 {
		log.Fatalf(
			"Please specify output format\n%v",
			usage,
		)
	}

	rstring := getdefiles.GetAllDeIn(
		os.Args[1],
		home,
		"/usr/share/applications/",
		"/usr/local/share/applications",
	)

	fmt.Print(rstring)
}
