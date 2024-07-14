package iconspec

// If init is false, initializes *a with the result of initv.
// Returns *a, and the returned error from initv.
func initRet[T string | int](a *T, init bool, initv func() (T, error)) (T, error) {
	if init {
		return *a, nil
	}

	v, err := initv()
	if err == nil {
		*a = v
	}
	return *a, err
}
