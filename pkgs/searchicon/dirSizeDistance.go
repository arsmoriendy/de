package searchicon

import (
	"errors"
	"fmt"
	"math"
	"os"

	"github.com/arsmoriendy/de/pkgs/searchicon/iconspec"
)

var dirSizeDistanceErr error = errors.New("Coultn't determine directory size distance")

func dirSizeDistance(idxFile *os.File, subdir string, iconsize int, iconscale int) (int, error) {
	ds := iconspec.NewDir(idxFile.Name(), subdir)

	type_, err := ds.Type()
	if err != nil {
		err = fmt.Errorf("%w: %w", dirSizeDistanceErr, err)
		return 0, err
	}

	scale := ds.Scale()

	switch type_ {
	case "Fixed":
		size, err := ds.Size()
		if err != nil {
			err = fmt.Errorf("%w: %w", dirSizeDistanceErr, err)
			return 0, err
		}

		dist := math.Abs(float64(scale*size - iconsize*iconscale))
		return int(dist), nil
	case "Scalable", "Scaled":
		minsize, err := ds.MinSize()
		if err != nil {
			err = fmt.Errorf("%w: %w", dirSizeDistanceErr, err)
			return 0, err
		}

		maxsize, err := ds.MaxSize()
		if err != nil {
			err = fmt.Errorf("%w: %w", dirSizeDistanceErr, err)
			return 0, err
		}

		if iconsize*iconscale < minsize*scale {
			return minsize*scale - iconsize*iconscale, nil
		}
		if iconsize*iconscale > maxsize*scale {
			return iconsize*iconscale - maxsize*scale, nil
		}
		return 0, nil

	case "Threshold":
		size, err := ds.Size()
		if err != nil {
			err = fmt.Errorf("%w: %w", dirSizeDistanceErr, err)
			return 0, err
		}

		threshold := ds.Threshold()

		minsize, err := ds.MinSize()
		if err != nil {
			err = fmt.Errorf("%w: %w", dirSizeDistanceErr, err)
			return 0, err
		}

		maxsize, err := ds.MaxSize()
		if err != nil {
			err = fmt.Errorf("%w: %w", dirSizeDistanceErr, err)
			return 0, err
		}

		if iconsize*iconscale < (size-threshold)*scale {
			return minsize*scale - iconsize*iconscale, nil
		}
		if iconsize*iconscale > (size+threshold)*scale {
			return iconscale*iconsize - maxsize*scale, nil
		}
		return 0, nil
	}

	return 0, nil
}
