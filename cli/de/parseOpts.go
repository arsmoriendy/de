package main

import (
	"flag"
	"os"
	"path"
)

type options struct {
	paths   []string
	format  string
	filters map[string]string
}

func parseOpts() options {
	// defaults for Func flags
	ropts := options{
		paths: []string{
			"/usr/share/applications/",
			"/usr/local/share/applications/",
			path.Join(os.Getenv("HOME"), ".local/share/applications/"),
		},
		filters: map[string]string{},
	}

	hasp := false // does the arguments have p
	flag.Func(
		"p",
		"`directory` to search for desktop entry files. Multiple  instances of this flag can be omited",
		func(s string) error {
			// clear default
			if !hasp {
				ropts.paths = nil
				hasp = true
			}

			ropts.paths = append(ropts.paths, s)

			return nil
		},
	)

	flag.StringVar(
		&ropts.format,
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
					ropts.filters[key] += c
				}
			}

			return nil
		},
	)

	flag.Parse()

	return ropts
}
