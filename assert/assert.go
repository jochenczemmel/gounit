// package assert provides some utility functions for unit testing.
// It uses the methods 'Helper', 'Errorf' and 'Fatalf', it can be used
// with testing.T, testing.B and testing.F.
package assert

// Error messages
const (
	msgErrorNotDetected = "ERROR: error not detected"
	msgUnexpectedError  = "ERROR: unexpected error: \"%v\""
	msgGotWant          = "ERROR: got: \"%v\", want: \"%v\""
	msgLengthGotWant    = "ERROR: length: got: %d, want: %d"
	msgSliceMapGotWant  = "ERROR: [%v]: got: \"%v\", want: \"%v\""
	msgWantUnequal      = "ERROR: got: \"%v\", want something unequal"
	msgKeyMissing       = "ERROR: key missing: \"%v\""
)

// ErrorHelper enables using the functions with testing.T,
// testing.B and testing.F and for unit testing.
type ErrorHelper interface {
	Helper()
	Errorf(string, ...any)
}

// FailHelper enables using the functions with testing.T,
// testing.B and testing.F and for unit testing.
type FailHelper interface {
	Helper()
	Failf(string, ...any)
}

// Equal checks if two values are equal. It calls e.Errorf() if not.
func Equal[C comparable](e ErrorHelper, got, want C) {
	e.Helper()
	if got != want {
		e.Errorf(msgGotWant, got, want)
	}
}

// Equal checks if two values are equal. It calls f.Failf() if not.
func EqualFail[C comparable](f FailHelper, got, want C) {
	f.Helper()
	if got != want {
		f.Failf(msgGotWant, got, want)
	}
}

// NotEqual checks if two values are not equal, otherwise
// it calls e.Errorf().
func NotEqual[C comparable](e ErrorHelper, got, want C) {
	e.Helper()
	if got == want {
		e.Errorf(msgWantUnequal, got)
	}
}

// EqualList checks if two lists are equal.
// If the lenght is different, the elements are not compared.
func EqualList[L ~[]C, C comparable](e ErrorHelper, got, want L) {

	e.Helper()

	if len(got) != len(want) {
		e.Errorf(msgLengthGotWant, len(got), len(want))
		return
	}

	for i, w := range want {
		if got[i] != w {
			e.Errorf(msgSliceMapGotWant, i, got[i], w)
		}
	}
}

// EqualMap checks if two maps are equal.
// If the lenght is different, the elements are not compared.
func EqualMap[M ~map[K]V, K, V comparable](e ErrorHelper,
	got, want M) {

	e.Helper()

	if len(got) != len(want) {
		e.Errorf(msgLengthGotWant, len(got), len(want))
		return
	}

	for k, w := range want {
		g, ok := got[k]
		if !ok {
			e.Errorf(msgKeyMissing, k)
			continue
		}
		if g != w {
			e.Errorf(msgSliceMapGotWant, k, g, w)
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
		e.Errorf(msgErrorNotDetected)
	}
}

// NoError calls e.Errorf() if err contains an error.
// In table tests, use function 'Error' with boolean values.
func NoError(e ErrorHelper, err error) {
	e.Helper()
	if err != nil {
		e.Errorf(msgUnexpectedError, err)
	}
}

// ErrorFail calls e.Failf() if an error ist wanted and not received
// or if an error is not wanted but received.
func ErrorFail(f FailHelper, err error, want bool) {
	f.Helper()
	if want {
		IsErrorFail(f, err)
		return
	}
	NoErrorFail(f, err)
}

// IsErrorFail calls f.Failf() if err contains no error.
// In table tests, use function 'ErrorFail' with boolean values.
func IsErrorFail(f FailHelper, err error) {
	f.Helper()
	if err == nil {
		f.Failf(msgErrorNotDetected)
	}
}

// NoErrorFail calls f.Failf() if err contains an error.
// In table tests, use function 'ErrorFail' with boolean values.
func NoErrorFail(f FailHelper, err error) {
	f.Helper()
	if err != nil {
		f.Failf(msgUnexpectedError, err)
	}
}

// True checks if expression is true.
// In table tests, use 'Equal' or 'EqualFail' with boolean values.
func True(t ErrorHelper, got bool) {
	t.Helper()
	if !got {
		t.Errorf(msgGotWant, got, true)
	}
}

// False checks if expression is not true.
// In table tests, use 'Equal' or 'EqualFail' with boolean values.
func False(t ErrorHelper, got bool) {
	t.Helper()
	if got {
		t.Errorf(msgGotWant, got, false)
	}
}
