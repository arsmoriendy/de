package getdefiles

import (
	"os"
	"path"
	"regexp"

	"github.com/arsmoriendy/de/pkgs/getde"
)

func GetAllDeIn(format string, filters *map[string]string, paths ...string) string {
	rstring := ""
	validDF := regexp.MustCompile(`\.desktop$`)
	parseEntry := createParseEntry(format)
	regMap := genRegMap(filters)

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

			rstring = rstring + getde.GetDe(*file, parseEntry, &regMap)
		}
	}

	return rstring
}
