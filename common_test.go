package release

import (
	"errors"
	"testing"
)

type tests []struct {
	Name        string
	InputRepo   string
	InputMatch  string
	Expected    string
	ExpectedErr error
	F           func(string, string) (string, error)
}

func (tl tests) validate(t *testing.T) {
	for _, tc := range tl {
		actual, actualErr := tc.F(tc.InputRepo, tc.InputMatch)

		if actual != tc.Expected || !errors.Is(actualErr, tc.ExpectedErr) {
			t.Errorf("\nInput Repo: %s \tInputMatch: %s\nExpected: (%s, %v)\nActual:   (%s, %v)",
				tc.InputRepo, tc.InputMatch, tc.Expected, tc.ExpectedErr, actual, actualErr)
		}
	}
}
