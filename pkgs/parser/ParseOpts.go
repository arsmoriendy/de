package parser

import (
	"flag"
	"os"
	"path"
)

type Options struct {
	Paths   []string          // -p
	Format  string            // -f
	Filters map[string]string // -F
	Names   []string          // -n
	Urls    []string          // -u
}

func ParseOpts() Options {
	// defaults for Func flags
	ropts := Options{
		Paths: []string{
			"/usr/share/applications/",
			"/usr/local/share/applications/",
			path.Join(os.Getenv("HOME"), ".local/share/applications/"),
		},
		Filters: map[string]string{},
		Names:   []string{},
		Urls:    []string{},
	}

	hasp := false // does the arguments have p
	flag.Func(
		"p",
		"`directory` to search for desktop entry files. Multiple  instances of this flag can be omited",
		func(s string) error {
			// clear default
			if !hasp {
				ropts.Paths = nil
				hasp = true
			}

			ropts.Paths = append(ropts.Paths, s)

			return nil
		},
	)

	flag.StringVar(
		&ropts.Format,
		"f",
		"{Name}={Icon}={Exec}",
		"output `format`. {} will be replaced with the value of the key",
	)

	flag.Func(
		"F",
		"regexp filters to match each entry with. Multiple instances of this flag can be omited. Format: \"`Key:RegexpValue`\" (RegexpValue syntax: https://pkg.go.dev/regexp/syntax)",
		func(s string) error {
			key := ""
			getkey := true

			for _, c := range s {

				switch c {
				case ':':
					getkey = false
					continue
				}

				c := string(c)
				if getkey {
					key += c
				} else {
					ropts.Filters[key] += c
				}
			}

			return nil
		},
	)

	flag.Func(
		"n",
		"file `name` as an argument for the Exec key. Multiple instances of this flag can be omited",
		func(s string) error {
			ropts.Names = append(ropts.Names, s)
			return nil
		},
	)

	flag.Func(
		"u",
		"`url` as an argument for the Exec key. Multiple instances of this flag can be omited",
		func(s string) error {
			ropts.Urls = append(ropts.Urls, s)
			return nil
		},
	)

	flag.Parse()

	return ropts
}
