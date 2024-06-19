package searchicon

import (
	"errors"
	"io/fs"
)

// Calls f(dir) on every dir on a string.
// If f returns fs.Skipall, stops iteration.
func walkDirsStr(dirsStr string, separator byte, f func(dir string) error) {
	dir := ""
	for _, c := range dirsStr {
		switch c {
		case rune(separator):
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
