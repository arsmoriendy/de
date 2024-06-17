package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/arsmoriendy/de/pkgs/getdefiles"
)

func main() {
	home := path.Join(os.Getenv("HOME"), ".local/share/applications/")

	if len(os.Args) != 2 {
		log.Fatal("Please specify output format")
	}

	rstring := getdefiles.GetAllDeIn(
		os.Args[1],
		home,
		"/usr/share/applications/",
		"/usr/local/share/applications",
	)

	fmt.Print(rstring)
}
