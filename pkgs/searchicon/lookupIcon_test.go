package searchicon

import (
	"errors"
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

	theme := "SampleAdwaita"

	th.FatalIfErr(t, th.OkAsExp(func() (string, error) {
		return lookupIcon(
			"audio-volume-low-symbolic",
			128,
			1,
			theme,
		)
	}, filepath.Join(
		xdgDataDirs,
		"icons/SampleAdwaita/symbolic/status/audio-volume-low-symbolic.svg",
	)))

	// Test unavailable icon, should fail
	_, err := lookupIcon(
		"example-unavailable-icon-name",
		128,
		1,
		theme)

	th.FatalIfErr(t, th.GenericAsExp(err, IconNotFound, errors.Is(err, IconNotFound)))

}
