package parsede

import (
	"os"
	"path"
	"regexp"
)

func ParseDeFilesIn(format string, filters *map[string]string, paths ...string) string {
	rstring := ""
	validDF := regexp.MustCompile(`\.desktop$`)
	parseEntry := genFormatEntryFunc(format)
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

			rstring = rstring + parseDeFile(*file, parseEntry, &regMap)
		}
	}

	return rstring
}
