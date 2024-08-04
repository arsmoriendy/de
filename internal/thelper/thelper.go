package thelper

import (
	"fmt"
	"testing"
)

type Thelper struct {
	Tptr *testing.T
}

func (th Thelper) FatalIfErr(err error) {
	if err != nil {
		th.Tptr.Fatal(err)
	}
}

// *As expected*. Checks if `in` is the same as `exp`.
func AsExp[T comparable](in T, exp T) error {
	if in != exp {
		return fmt.Errorf("\nExpected:\t%v\nGot Result:\t%v", exp, in)
	}

	return nil
}

// *Ok(no errors) and as expected*.
// Wrapper for handling `f`'s error and running asexp on it's return value.
func OkAsExp[T comparable](f func() (T, error), exp T) error {
	in, err := f()

	if err != nil {
		return err
	}

	return AsExp(in, exp)
}

// AsExp for custom comparable types
func GenericAsExp[T interface{}](in T, exp T, comp bool) error {
	if !comp {
		return fmt.Errorf("\nExpected:\t%v\nGot Result:\t%v", exp, in)
	}

	return nil
}
