package iconspec_test

import (
	"path/filepath"
	"testing"

	"github.com/arsmoriendy/de/pkgs/searchicon/iconspec"
)

func TestDirSpec(t *testing.T) {
	idxFn, _ := filepath.Abs("../../../test/data/sample_xdg_data_dir/icons/SampleAdwaita/index.theme")
	subdir := "scalable/devices"

	ds := iconspec.NewDir(idxFn, subdir)

	// Size check {
	size, err := ds.Size()
	if err != nil {
		t.Fatal(err)
	}

	exp := 128

	if size != exp {
		t.Fatalf("\nExpected:\t%v\nGot Result:\t%v", exp, size)
	}
	// }

	// Scale check {
	scale := ds.Scale()

	exp = 1

	if scale != exp {
		t.Fatalf("\nExpected:\t%v\nGot Result:\t%v", exp, scale)
	}
	// }
}
