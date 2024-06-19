package searchicon

import (
	"errors"
	"io/fs"
)

// If f returns fs.Skipall, stops iteration
func walkDirsStr(dirsStr string, f func(dir string) error) {
	dir := ""
	for _, c := range dirsStr {
		switch c {
		case ':':
			err := f(dir)
			if errors.Is(err, fs.SkipAll) {
				return
			}

			dir = ""
		default:
			dir += string(c)
		}
	}
	f(dir)
}
