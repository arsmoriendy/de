package searchicon

import (
	"fmt"
	"os"
	"path"
)

// Don't edit directly. Not guaranteed to be absolute paths.
var baseDirs string
var isBaseDirsInit = false

func initBaseDirs() {
	if isBaseDirsInit == true {
		return
	}

	home := os.Getenv("HOME")
	xdgDataDirs := "/usr/local/share:/usr/share"

	if envXdgDataDirs, isset := os.LookupEnv("XDG_DATA_DIRS"); isset {
		xdgDataDirs = envXdgDataDirs
	}

	baseDirs = fmt.Sprintf("%v:%v:%v",
		path.Join(home, ".icons"),
		appendXdgDataDirs(&xdgDataDirs, "icons"),
		"/usr/share/pixmaps",
	)

	isBaseDirsInit = true
}

func appendXdgDataDirs(xdgDataDirs *string, s string) string {
	rstring := ""

	xdgDataDir := ""
	for _, c := range *xdgDataDirs {
		switch c {
		case ':':
			rstring += path.Join(xdgDataDir, s) + ":"
			xdgDataDir = ""
		default:
			xdgDataDir += string(c)
		}
	}
	rstring += path.Join(xdgDataDir, s)

	return rstring
}
