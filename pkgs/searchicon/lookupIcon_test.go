package searchicon

import (
	"os"
	"path/filepath"
	"testing"

	th "github.com/arsmoriendy/de/internal/thelper"
)

func TestLookupIcon(t *testing.T) {
	// init basedirs
	xdgDataDirs, _ := filepath.Abs("../../test/data/sample_xdg_data_dir/")
	os.Setenv("XDG_DATA_DIRS", xdgDataDirs)
	initBaseDirs()

	th.FatalIfErr(t, th.OkAsExp(func() (string, error) {
		return lookupIcon(
			"audio-volume-low-symbolic",
			128,
			1,
			"SampleAdwaita",
		)
	}, filepath.Join(
		xdgDataDirs,
		"icons/SampleAdwaita/symbolic/status/audio-volume-low-symbolic.svg",
	)))
}
