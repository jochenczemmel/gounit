package assert

// ErrorHelper enables using the functions with testing.T,
// testing.B and testing.F and for unit testing.
type ErrorHelper interface {
	Helper()
	Errorf(string, ...any)
}

// Equal checks if two values are equal. It calls Error if not.
func Equal[C comparable](e ErrorHelper, got, want C) {
	e.Helper()
	if got != want {
		e.Errorf("ERROR: got\n%v\nwant\n%v\n", got, want)
	}
}
