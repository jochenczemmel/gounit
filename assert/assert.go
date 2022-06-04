package assert

// ErrorHelper enables using the functions with testing.T,
// testing.B and testing.F and for unit testing.
type ErrorHelper interface {
	Helper()
	Errorf(string, ...any)
}

// Equal checks if two values are equal. It calls t.Error() if not.
func Equal[C comparable](e ErrorHelper, got, want C) {

	e.Helper()

	if got != want {
		e.Errorf("ERROR: got: \"%v\", want: \"%v\"", got, want)
	}
}

// NotEqual checks if two values are not equal, otherwise
// it calls t.Error().
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
