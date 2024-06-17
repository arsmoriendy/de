package getdefiles

import (
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/arsmoriendy/de/pkgs/getde"
)

func GetAllDeIn(paths ...string) string {
	rstring := ""
	validDF := regexp.MustCompile(`\.desktop$`)
	for _, p := range paths {
		files, err := os.ReadDir(p)
		if err != nil {
			continue
		}

		for _, dentry := range files {
			if dentry.IsDir() {
				continue
			}

			fname := dentry.Name()

			if !validDF.MatchString(fname) {
				continue
			}

			absfname := path.Join(p, fname)

			file, err := os.Open(absfname)
			defer file.Close()
			if err != nil {
				continue
			}

			rstring = rstring + getde.GetDe(*file, parseEntry)
		}
	}

	return rstring
}

func parseEntry(entry *map[string]string) string {
	return fmt.Sprintf(
		"%v=%v=%v",
		(*entry)["Name"],
		(*entry)["Icon"],
		(*entry)["Exec"],
	)
}
