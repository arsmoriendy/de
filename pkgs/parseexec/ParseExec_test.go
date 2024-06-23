package parseexec

import (
	"testing"

	"github.com/arsmoriendy/de/pkgs/parser"
)

func TestLowerFMulti(t *testing.T) {
	s := "program %f -f"
	entry := map[string]string{}
	opts := parser.Options{
		Names: []string{
			"/example/path/1",
			"/example/path/2",
			"/example/path/3",
		},
	}
	filename := ""

	exec := ParseExec(s, &entry, &opts, filename)
	exp := "program /example/path/1 -f"

	if exec != exp {
		t.Errorf("\nExpected:\t%v\nGot Result:\t%v", exp, exec)
	}
}

func FuzzSingleFile(f *testing.F) {

	f.Add("/example/path")
	f.Add("/usr/local/bin/")
	f.Add("/root/ ")

	f.Fuzz(func(t *testing.T, file string) {
		s := "program %f -f"
		entry := map[string]string{}
		opts := parser.Options{
			Names: []string{file},
		}
		filename := ""

		exec := ParseExec(s, &entry, &opts, filename)

		if exec != "program "+file+" -f" {
			t.Errorf("\nRaw: \t%v\nName: \t%v\nParsed: \t%v\n", s, file, exec)
		}
	})
}
