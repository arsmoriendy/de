package searchicon

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/arsmoriendy/de/pkgs/parser"
)

// Gets a value from an index.theme file.
//
// HKV = Header Key Value.
func getHKV(idxFile *os.File, header string, key string) (string, error) {
	scanner := bufio.NewScanner(idxFile)

	// e = entry
	eheader := ""
	ekey := ""
	evalue := ""
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 1 {
			continue
		}

		switch line[0] {
		case '[':
			eheader = getHeader(line)
			continue
		case '#':
			continue
		}

		if eheader != header {
			continue
		}

		ekey, evalue = parser.ParseLine(&line)

		if ekey != key {
			continue
		}

		return evalue, nil
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return "", fmt.Errorf(`Key "%v" in header "%v" not found`, key, header)
}

func getHeader(line string) string {
	rstring := ""

	for _, c := range line {
		switch c {
		case '[':
		case ']':
		default:
			rstring += string(c)
		}
	}

	return rstring
}
