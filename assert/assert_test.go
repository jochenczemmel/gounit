package assert_test

import (
	"fmt"
	"testing"

	"github.com/jochenczemmel/gounit/assert"
)

// TestSpy implements assert.ErrorHelper
type TestSpy struct {
	ErrorCalled  bool
	ErrorMessage string
	FailCalled   bool
	FailMessage  string
}

func (t *TestSpy) Errorf(format string, args ...any) {
	t.ErrorCalled = true
	if t.ErrorMessage == "" {
		t.ErrorMessage = fmt.Sprintf(format, args...)
		return
	}
	t.ErrorMessage += "/" + fmt.Sprintf(format, args...)
}

func (t *TestSpy) Failf(format string, args ...any) {
	t.FailCalled = true
	if t.FailMessage == "" {
		t.FailMessage = fmt.Sprintf(format, args...)
		return
	}
	t.FailMessage += "/" + fmt.Sprintf(format, args...)
}

func (t *TestSpy) Helper() {}

func TestEqual(t *testing.T) {
	candidates := []struct {
		name           string
		value1, value2 string
		wantError      bool
		wantMessage    string
	}{
		{
			name:   "equal",
			value1: "a",
			value2: "a",
		},
		{
			name:        "not equal",
			value1:      "a",
			value2:      "b",
			wantError:   true,
			wantMessage: `ERROR: got: "a", want: "b"`,
		},
	}
	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {

			ts := &TestSpy{}
			assert.Equal(ts, c.value1, c.value2)
			if c.wantError {
				if !ts.ErrorCalled {
					t.Errorf("ERROR: error not detected")
				}
				if ts.ErrorMessage != c.wantMessage {
					t.Errorf("ERROR: got: \"%s\", want: \"%s\"",
						ts.ErrorMessage, c.wantMessage)
				}
			}
			if !c.wantError && ts.ErrorCalled {
				t.Errorf("ERROR: false alarm")
				t.Logf("NOTE: error message is: %v", ts.ErrorMessage)
			}

			ts = &TestSpy{}
			assert.EqualFail(ts, c.value1, c.value2)
			if c.wantError {
				if !ts.FailCalled {
					t.Errorf("ERROR: error not detected")
				}
				if ts.FailMessage != c.wantMessage {
					t.Errorf("ERROR: got: \"%s\", want: \"%s\"",
						ts.FailMessage, c.wantMessage)
				}
			}
			if !c.wantError && ts.FailCalled {
				t.Errorf("ERROR: false alarm")
				t.Logf("NOTE: error message is: %v", ts.FailMessage)
			}
		})
	}
}

type intList []int

func TestEqualList(t *testing.T) {

	candidates := []struct {
		name           string
		value1, value2 intList
		wantError      bool
		wantMessage    string
		// value1, value2 []int
	}{
		{
			name:   "equal",
			value1: intList{1, 2, 3},
			value2: intList{1, 2, 3},
		},
		{
			name:        "different length",
			value1:      []int{1, 2, 3},
			value2:      []int{1, 2},
			wantError:   true,
			wantMessage: `ERROR: length: got: 3, want: 2`,
		},
		{
			name:        "different value",
			value1:      []int{1, 2, 3},
			value2:      []int{1, 3, 3},
			wantError:   true,
			wantMessage: `ERROR: [1]: got: "2", want: "3"`,
		},
		{
			name:      "different values",
			value1:    []int{1, 2, 3},
			value2:    []int{1, 3, 2},
			wantError: true,
			wantMessage: `ERROR: [1]: got: "2", want: "3"/` +
				`ERROR: [2]: got: "3", want: "2"`,
		},
		{
			name:        "empty got list",
			value1:      []int{},
			value2:      []int{1, 2, 3},
			wantError:   true,
			wantMessage: `ERROR: length: got: 0, want: 3`,
		},
		{
			name:        "empty want list",
			value1:      []int{1, 2, 3},
			value2:      []int{},
			wantError:   true,
			wantMessage: `ERROR: length: got: 3, want: 0`,
		},
		{
			name:        "nil got list",
			value2:      []int{1, 2, 3},
			wantError:   true,
			wantMessage: `ERROR: length: got: 0, want: 3`,
		},
		{
			name:        "nil want list",
			value1:      []int{1, 2, 3},
			wantError:   true,
			wantMessage: `ERROR: length: got: 3, want: 0`,
		},
	}

	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			ts := &TestSpy{}
			assert.EqualList(ts, c.value1, c.value2)
			if c.wantError {
				if !ts.ErrorCalled {
					t.Errorf("ERROR: error not detected")
				}
				if ts.ErrorMessage != c.wantMessage {
					t.Errorf("ERROR: got: \"%s\", want: \"%s\"",
						ts.ErrorMessage, c.wantMessage)
				}
			}
			if !c.wantError && ts.ErrorCalled {
				t.Errorf("ERROR: false alarm")
				t.Logf("NOTE: error message is: %v", ts.ErrorMessage)
			}
		})
	}
}

type intStringMap map[int]string

func TestEqualMap(t *testing.T) {

	candidates := []struct {
		name           string
		value1, value2 intStringMap
		wantError      bool
		wantMessage    string
		// value1, value2 map[int]string
	}{
		{
			name:   "equal",
			value1: intStringMap{1: "one", 2: "two", 3: "three"},
			value2: intStringMap{1: "one", 2: "two", 3: "three"},
		},
		{
			name:        "different length",
			value1:      map[int]string{1: "one", 2: "two", 3: "three"},
			value2:      map[int]string{1: "one", 3: "three"},
			wantError:   true,
			wantMessage: `ERROR: length: got: 3, want: 2`,
		},
		{
			name:        "different key",
			value1:      map[int]string{1: "one", 2: "two", 3: "three"},
			value2:      map[int]string{1: "one", 4: "two", 3: "three"},
			wantError:   true,
			wantMessage: `ERROR: key missing: "4"`,
		},
		{
			name:        "empty value missing key",
			value1:      map[int]string{1: "one", 2: "two", 3: "three"},
			value2:      map[int]string{1: "one", 4: "", 3: "three"},
			wantError:   true,
			wantMessage: `ERROR: key missing: "4"`,
		},
		{
			name:        "different value",
			value1:      map[int]string{1: "one", 2: "two", 3: "three"},
			value2:      map[int]string{1: "ONE", 2: "two", 3: "three"},
			wantError:   true,
			wantMessage: `ERROR: [1]: got: "one", want: "ONE"`,
		},
		{
			name:        "nil got map",
			value2:      map[int]string{1: "one", 2: "two", 3: "three"},
			wantError:   true,
			wantMessage: `ERROR: length: got: 0, want: 3`,
		},
		{
			name:        "nil want map",
			value1:      map[int]string{1: "one", 2: "two", 3: "three"},
			wantError:   true,
			wantMessage: `ERROR: length: got: 3, want: 0`,
		},
	}

	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			ts := &TestSpy{}
			assert.EqualMap(ts, c.value1, c.value2)
			if c.wantError {
				if !ts.ErrorCalled {
					t.Errorf("ERROR: error not detected")
				}
				if ts.ErrorMessage != c.wantMessage {
					t.Errorf("ERROR: got: \"%s\", want: \"%s\"",
						ts.ErrorMessage, c.wantMessage)
				}
			}
			if !c.wantError && ts.ErrorCalled {
				t.Errorf("ERROR: false alarm")
				t.Logf("NOTE: error message is: %v", ts.ErrorMessage)
			}
		})
	}
}

func TestNotEqual(t *testing.T) {
	candidates := []struct {
		name           string
		value1, value2 string
		wantError      bool
		wantMessage    string
	}{
		{
			name:        "equal",
			value1:      "a",
			value2:      "a",
			wantError:   true,
			wantMessage: `ERROR: got: "a", want something unequal`,
		},
		{
			name:   "not equal",
			value1: "a",
			value2: "b",
		},
	}
	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			ts := &TestSpy{}
			assert.NotEqual(ts, c.value1, c.value2)
			if c.wantError {
				if !ts.ErrorCalled {
					t.Errorf("ERROR: error not detected")
				}
				if ts.ErrorMessage != c.wantMessage {
					t.Errorf("ERROR: got: \"%s\", want: \"%s\"",
						ts.ErrorMessage, c.wantMessage)
				}
			}
			if !c.wantError && ts.ErrorCalled {
				t.Errorf("ERROR: false alarm")
				t.Logf("NOTE: error message is: %v", ts.ErrorMessage)
			}
		})
	}
}

func TestError(t *testing.T) {
	candidates := []struct {
		name        string
		inError     error
		inWant      bool
		wantError   bool
		wantMessage string
	}{
		{
			name: "ok no error",
		},
		{
			name:    "ok error",
			inError: fmt.Errorf("errormessage"),
			inWant:  true,
		},
		{
			name:        "not ok error",
			inError:     fmt.Errorf("errormessage"),
			wantError:   true,
			wantMessage: `ERROR: unexpected error: "errormessage"`,
		},
		{
			name:        "not ok no error",
			inWant:      true,
			wantError:   true,
			wantMessage: `ERROR: error not detected`,
		},
	}
	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			ts := &TestSpy{}
			assert.Error(ts, c.inError, c.inWant)
			if c.wantError {
				if !ts.ErrorCalled {
					t.Errorf("ERROR: error not detected")
				}
				if ts.ErrorMessage != c.wantMessage {
					t.Errorf("ERROR: got: \"%s\", want: \"%s\"",
						ts.ErrorMessage, c.wantMessage)
				}
			}
			if !c.wantError && ts.ErrorCalled {
				t.Errorf("ERROR: false alarm")
				t.Logf("NOTE: error message is: %v", ts.ErrorMessage)
			}
		})
	}
}

func TestErrorFail(t *testing.T) {
	candidates := []struct {
		name        string
		inError     error
		inWant      bool
		wantError   bool
		wantMessage string
	}{
		{
			name: "ok no error",
		},
		{
			name:    "ok error",
			inError: fmt.Errorf("errormessage"),
			inWant:  true,
		},
		{
			name:        "not ok error",
			inError:     fmt.Errorf("errormessage"),
			wantError:   true,
			wantMessage: `ERROR: unexpected error: "errormessage"`,
		},
		{
			name:        "not ok no error",
			inWant:      true,
			wantError:   true,
			wantMessage: `ERROR: error not detected`,
		},
	}
	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			ts := &TestSpy{}
			assert.ErrorFail(ts, c.inError, c.inWant)
			if c.wantError {
				if !ts.FailCalled {
					t.Errorf("ERROR: error not detected")
				}
				if ts.FailMessage != c.wantMessage {
					t.Errorf("ERROR: got: \"%s\", want: \"%s\"",
						ts.FailMessage, c.wantMessage)
				}
			}
			if !c.wantError && ts.FailCalled {
				t.Errorf("ERROR: false alarm")
				t.Logf("NOTE: error message is: %v", ts.FailMessage)
			}
		})
	}
}

func TestTrueFalse(t *testing.T) {
	candidates := []struct {
		name             string
		isTrue           bool
		wantTrueError    bool
		wantTrueMessage  string
		wantFalseError   bool
		wantFalseMessage string
	}{
		{
			name:             "true",
			isTrue:           true,
			wantFalseError:   true,
			wantFalseMessage: `ERROR: got: "true", want: "false"`,
		},
		{
			name:            "false",
			wantTrueError:   true,
			wantTrueMessage: `ERROR: got: "false", want: "true"`,
		},
	}

	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {

			ts := &TestSpy{}
			assert.True(ts, c.isTrue)
			if c.wantTrueError {
				if !ts.ErrorCalled {
					t.Errorf("ERROR: error not detected")
				}
				if ts.ErrorMessage != c.wantTrueMessage {
					t.Errorf("ERROR: got: \"%s\", want: \"%s\"",
						ts.ErrorMessage, c.wantTrueMessage)
				}
			}
			if !c.wantTrueError && ts.ErrorCalled {
				t.Errorf("ERROR: false alarm")
				t.Logf("NOTE: error message is: %v", ts.ErrorMessage)
			}

			ts = &TestSpy{}
			assert.False(ts, c.isTrue)
			if c.wantFalseError {
				if !ts.ErrorCalled {
					t.Errorf("ERROR: error not detected")
				}
				if ts.ErrorMessage != c.wantFalseMessage {
					t.Errorf("ERROR: got: \"%s\", want: \"%s\"",
						ts.ErrorMessage, c.wantFalseMessage)
				}
			}
			if !c.wantFalseError && ts.ErrorCalled {
				t.Errorf("ERROR: false alarm")
				t.Logf("NOTE: error message is: %v", ts.ErrorMessage)
			}

		})
	}
}
