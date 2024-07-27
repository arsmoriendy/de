package iconspec_test

import (
	"path/filepath"
	"testing"

	"github.com/arsmoriendy/de/pkgs/searchicon/iconspec"
	th "github.com/arsmoriendy/de/internal/thelper"
)

func TestDirSpec(t *testing.T) {
	idxFn, _ := filepath.Abs("../../../test/data/sample_xdg_data_dir/icons/SampleAdwaita/index.theme")
	subdir := "scalable/devices"

	ds := iconspec.NewDir(idxFn, subdir)

	// Size
	th.FatalIfErr(t, th.OkAsExp(ds.Size, 128))

	// Scale
	th.FatalIfErr(t, th.AsExp(ds.Scale(), 1))

	// Context
	th.FatalIfErr(t, th.OkAsExp(ds.Context, "Devices"))

	// Type
	th.FatalIfErr(t, th.OkAsExp(ds.Type, "Scalable"))

	// MaxSize
	th.FatalIfErr(t, th.OkAsExp(ds.MaxSize, 512))

	// MinSize
	th.FatalIfErr(t, th.OkAsExp(ds.MinSize, 8))

	// Threshold
	th.FatalIfErr(t, th.AsExp(ds.Threshold(), 2))
}
