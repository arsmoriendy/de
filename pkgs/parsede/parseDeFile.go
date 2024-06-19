package parsede

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/arsmoriendy/de/pkgs/parser"
)

// Filters, and formats a desktop entry file into a string
func parseDeFile(
	deFile os.File,
	formatEntry func(*map[string]string) string,
	regMap *map[string]*regexp.Regexp,
) string {
	rstring := ""

	scanner := bufio.NewScanner(&deFile)

	entry := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 1 {
			continue
		}

		switch line[0] {
		case '[':
			if len(entry) != 0 && entryMatches(&entry, regMap) {
				rstring = rstring + formatEntry(&entry) + "\n"
				clear(entry)
			}
			continue
		case '#':
			continue
		}

		key, value := parser.ParseLine(&line)

		entry[key] = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	if entryMatches(&entry, regMap) {
		rstring = rstring + formatEntry(&entry) + "\n"
	}

	return rstring
}
