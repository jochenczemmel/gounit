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
}

func (t *TestSpy) Errorf(format string, args ...any) {
	t.ErrorCalled = true
	t.ErrorMessage = fmt.Sprintf(format, args...)
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
			name:      "equal",
			value1:    "a",
			value2:    "a",
			wantError: false,
		},
		{name: "not equal",
			value1:      "a",
			value2:      "b",
			wantError:   true,
			wantMessage: `ERROR: got "a", want "b"`,
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
					t.Errorf("ERROR: got: %s, want: %s",
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
