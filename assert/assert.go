package assert

// ErrorHelper enables using the functions with testing.T,
// testing.B and testing.F and for unit testing.
type ErrorHelper interface {
	Helper()
	Errorf(string, ...any)
}

// Equal checks if two values are equal. It calls e.Errorf() if not.
func Equal[C comparable](e ErrorHelper, got, want C) {

	e.Helper()

	if got != want {
		e.Errorf("ERROR: got: \"%v\", want: \"%v\"", got, want)
	}
}

// NotEqual checks if two values are not equal, otherwise
// it calls e.Errorf().
func NotEqual[C comparable](e ErrorHelper, got, want C) {

	e.Helper()

	if got == want {
		e.Errorf("ERROR: got: \"%v\", want something unequal", got)
	}
}

// EqualList checks if two lists are equal.
// If the lenght is different, the elements are not compared.
func EqualList[L ~[]C, C comparable](e ErrorHelper, got, want L) {

	e.Helper()

	if len(got) != len(want) {
		e.Errorf("ERROR: length: got: %d, want: %d",
			len(got), len(want))
		return
	}

	for i, w := range want {
		if got[i] != w {
			e.Errorf("ERROR: [%d] got: \"%v\", want: \"%v\"", i, got[i], w)
		}
	}
}

// EqualMap checks if two maps are equal.
// If the lenght is different, the elements are not compared.
func EqualMap[M ~map[K]V, K, V comparable](e ErrorHelper,
	got, want M) {

	e.Helper()

	if len(got) != len(want) {
		e.Errorf("ERROR: length: got: %d, want: %d", len(got), len(want))
		return
	}

	for k, w := range want {
		g, ok := got[k]
		if !ok {
			e.Errorf("ERROR: wanted key missing: \"%v\"", k)
			continue
		}
		if g != w {
			e.Errorf("ERROR: [%v]: got: \"%v\", want: \"%v\"", k, g, w)
		}
	}
}

// Error calls e.Errorf() if an error ist wanted and not received
// or if an error is not wanted but received.
func Error(e ErrorHelper, err error, want bool) {

	e.Helper()

	if want {
		IsError(e, err)
		return
	}
	NoError(e, err)
}

// IsError calls e.Errorf() if err contains no error.
// In table tests, use function 'Error' with boolean values.
func IsError(e ErrorHelper, err error) {
	e.Helper()
	if err == nil {
		e.Errorf("ERROR: error not detected")
	}
}

// NoError calls e.Errorf() if err contains an error.
// In table tests, use function 'Error' with boolean values.
func NoError(e ErrorHelper, err error) {
	e.Helper()
	if err != nil {
		e.Errorf("ERROR: unexpected error: \"%v\"", err)
	}
}
