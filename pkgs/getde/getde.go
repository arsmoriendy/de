package getde

import (
	"bufio"
	"log"
	"os"
)

func GetDe(entryfile os.File, parseEntry func(*map[string]string) string) string {
	rstring := ""

	scanner := bufio.NewScanner(&entryfile)

	entry := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 1 {
			continue
		}

		switch line[0] {
		case '[':
			if len(entry) != 0 {
				rstring = rstring + parseEntry(&entry) + "\n"
				clear(entry)
			}
			continue
		case '#':
			continue
		}

		key, value := parseLine(&line)

		entry[key] = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	rstring = rstring + parseEntry(&entry) + "\n"

	return rstring
}
