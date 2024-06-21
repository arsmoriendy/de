package parsede

import (
	"os"
	"path"
	"regexp"

	"github.com/arsmoriendy/de/pkgs/parser"
)

func ParseDeFiles(opts *parser.Options) string {
	rstring := ""
	validDF := regexp.MustCompile(`\.desktop$`)
	parseEntry := genFormatEntryFunc((*opts).Format)
	regMap := genRegMap(&(*opts).Filters)

	for _, p := range (*opts).Paths {
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
