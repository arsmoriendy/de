package main

import (
	"fmt"
	"os"
	"path"

	"github.com/arsmoriendy/de/pkgs/getdefiles"
)

func main() {
	home := path.Join(os.Getenv("HOME"), ".local/share/applications/")

	rstring := getdefiles.GetAllDeIn(
		home,
		"/usr/share/applications/",
		"/usr/local/share/applications",
	)

	fmt.Print(rstring)
}
