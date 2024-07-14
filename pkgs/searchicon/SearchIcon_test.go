package searchicon

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/arsmoriendy/de/pkgs/searchicon/gethkv"
)

const iconname = "firefox"
const size = "16x16"
const scale = "2"
const theme = "SampleAdwaita"

func TestSearchIcon(t *testing.T) {
	xdgDataDirs, err := filepath.Abs("../../test/data/sample_xdg_data_dir/")

	os.Setenv("XDG_DATA_DIRS", xdgDataDirs)

	initBaseDirs()

	tDir := getThemeDirs(theme)
	eTDir := filepath.Join(xdgDataDirs, "icons/SampleAdwaita")
	if tDir != eTDir {
		t.Fatalf(
			"Theme dir mismatch \nExpected:\t%v\nGot Result:\t%v",
			eTDir,
			tDir,
		)
	}

	idxFile, err := getThemeIdx(tDir)
	if err != nil {
		t.Fatal(err)
	}

	inherits, err := searchicon.GetHKV(idxFile.Name(), "Icon Theme", "Inherits")
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal(inherits, err)
}
