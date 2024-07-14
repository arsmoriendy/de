package iconspec

// If init is true, return the value of *a.
//
// If init is false, initializes *a to the value of initv(). Returns the value of *a.
func initRet[T string](a *T, init *bool, initv func() (T, error)) (T, error) {
	if *init {
		return *a, nil
	}

	v, err := initv()
	if err == nil {
		*a = v
		*init = true
	}
	return *a, nil
}
