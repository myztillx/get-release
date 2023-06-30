package release

import (
	"errors"
	"testing"
)

type tests []struct {
	Name        string
	Input       string
	Expected    string
	ExpectedErr error
	F           func(string) (string, error)
}

func (tl tests) validate(t *testing.T) {
	for _, tc := range tl {
		actual, actualErr := tc.F(tc.Input)

		if actual != tc.Expected || !errors.Is(actualErr, tc.ExpectedErr) {
			t.Errorf("\nInput: %s\nExpected: (%s, %v)\nActual:   (%s, %v)",
				tc.Input, tc.Expected, tc.ExpectedErr, actual, actualErr)
		}
	}
}
