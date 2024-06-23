// Parser for the "Exec" key on a desktop entry.
// Ref: https://specifications.freedesktop.org/desktop-entry-spec/latest/ar01s07.html
package parseexec

import (
	"fmt"
	"os"

	"github.com/arsmoriendy/de/pkgs/parser"
)

func ParseExec(entry *map[string]string, opts *parser.Options, filename string) string {
	rstring := ""

	captureFieldCode := false
	for _, c := range (*entry)["Exec"] {
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
				if len(opts.Names) < 1 {
					break
				}

				rstring += opts.Names[0]
			case 'F':
				for i, fname := range opts.Names {
					if i > 0 {
						rstring += " "
					}

					rstring += fname
				}
			case 'u':
				if len(opts.Urls) < 1 {
					break
				}

				rstring += opts.Urls[0]
			case 'U':
				for i, url := range opts.Urls {
					if i > 0 {
						rstring += " "
					}

					rstring += url
				}
			case 'i':
				icon, found := (*entry)["Icon"]

				if !found {
					break
				}

				rstring += fmt.Sprintf(`--icon "%v"`, icon)
			// translated entry name
			case 'c':
				lang := os.Getenv("LANG")

				if lang == "" {
					break
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
			captureFieldCode = false
		}
	}

	return rstring
}
