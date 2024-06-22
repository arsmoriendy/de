// Parser for the "Exec" key on a desktop entry.
// Ref: https://specifications.freedesktop.org/desktop-entry-spec/latest/ar01s07.html
package parseexec

import (
	"fmt"
	"os"

	"github.com/arsmoriendy/de/pkgs/parser"
)

func ParseExec(s string, entry *map[string]string, opts *parser.Options, filename string) string {
	rstring := ""

	captureFieldCode := false
	for _, c := range s {
		switch c {
		// TODO: handle quoting
		case '"':
		case '%':
			captureFieldCode = true
		default:
			if !captureFieldCode {
				rstring += string(c)
				continue
			}

			switch c {
			case '%':
				rstring += "%"
			case 'f':
				if len(opts.Names) == 1 {
					rstring += opts.Names[0]
				}
			case 'F':
				for _, fname := range opts.Names {
					rstring += fname
				}
			case 'u':
				if len(opts.Names) == 1 {
					rstring += opts.Names[0]
				}
			case 'U':
				for _, url := range opts.Urls {
					rstring += url
				}
			case 'i':
				icon, found := (*entry)["Icon"]

				if !found {
					continue
				}

				rstring += fmt.Sprintf(`--icon "%v"`, icon)
			// translated entry name
			case 'c':
				lang := os.Getenv("LANG")

				if lang == "" {
					continue
				}

				name, found := (*entry)["Name["+lang[0:2]+"]"]

				if !found {
					name = (*entry)["Name"]
				}

				rstring += name
			// entry filename
			case 'k':
				rstring += filename
			}

		}
	}

	return rstring
}
