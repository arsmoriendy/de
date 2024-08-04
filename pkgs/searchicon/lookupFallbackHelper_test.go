package searchicon

import (
	"os"
	fp "path/filepath"
	"testing"

	b "github.com/arsmoriendy/de/internal/bind"
	th "github.com/arsmoriendy/de/internal/thelper"
)

func TestLookupFallbackHelper(t *testing.T) {
	// init basedirs
	xdgDataDirs, _ := fp.Abs("../../test/data/sample_xdg_data_dir/")
	os.Setenv("XDG_DATA_DIRS", xdgDataDirs)
	initBaseDirs()

	theme := "SampleAdwaita"
	themeDir := fp.Join(xdgDataDirs, "icons", theme)

	ths := th.Thelper{
		Tptr: t,
	}

	// test "audio-headset" icon
	ths.FatalIfErr(th.OkAsExp(b.Bind[string](lookupFallbackHelper, "audio-headset", theme),
		fp.Join(themeDir, "16x16/devices/audio-headset.png", // exp
	)))

}
