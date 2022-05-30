package assert_test

import (
	"testing"

	"github.com/jochenczemmel/gounit/assert"
)

// TestSpy implements assert.ErrorHelper
type TestSpy struct {
	ErrorCalled bool
}

func (t *TestSpy) Errorf(string, ...any) {
	t.ErrorCalled = true
}

func (t *TestSpy) Helper() {}

func TestEqual(t *testing.T) {
	candidates := []struct {
		name           string
		value1, value2 string
		wantError      bool
	}{
		{"equal", "a", "a", false},
		{"not equal", "a", "b", true},
	}
	for _, c := range candidates {
		t.Run(c.name, func(t *testing.T) {
			ts := &TestSpy{}
			assert.Equal(ts, c.value1, c.value2)
			if c.wantError && !ts.ErrorCalled {
				t.Errorf("ERROR: error not detected")
			}
			if !c.wantError && ts.ErrorCalled {
				t.Errorf("ERROR: false alarm")
			}
		})
	}
}
