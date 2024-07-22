package iconspec_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/arsmoriendy/de/pkgs/searchicon/iconspec"
)

func TestDirSpec(t *testing.T) {
	idxFn, _ := filepath.Abs("../../../test/data/sample_xdg_data_dir/icons/SampleAdwaita/index.theme")
	subdir := "scalable/devices"

	ds := iconspec.NewDir(idxFn, subdir)

	// Size
	fataliferr(t, okasexp(ds.Size, 128))

	// Scale
	fataliferr(t, asexp(ds.Scale(), 1))

	// Context
	fataliferr(t, okasexp(ds.Context, "Devices"))

	// Type
	fataliferr(t, okasexp(ds.Type, "Scalable"))

	// MaxSize
	fataliferr(t, okasexp(ds.MaxSize, 512))

	// MinSize
	fataliferr(t, okasexp(ds.MinSize, 8))

	// Threshold
	fataliferr(t, asexp(ds.Threshold(), 2))
}

func fataliferr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// *As expected*. Checks if `in` is the same as `exp`.
func asexp[T comparable](in T, exp T) error {
	if in != exp {
		return fmt.Errorf("\nExpected:\t%v\nGot Result:\t%v", exp, in)
	}

	return nil
}

// *Ok(no errors) and as expected*.
// Wrapper for handling `f`'s error and running asexp on it's return value.
func okasexp[T comparable](f func() (T, error), exp T) error {
	in, err := f()

	if err != nil {
		return err
	}

	return asexp(in, exp)
}
