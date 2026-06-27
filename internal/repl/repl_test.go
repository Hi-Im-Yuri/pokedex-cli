package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: nil,
		},
		{
			input:    "  Charmander Bulbasar  PIKACHU",
			expected: []string{"charmander", "bulbasar", "pikachu"},
		},
		{
			input:    "trapinch    torchick kyogre evee",
			expected: []string{"trapinch", "torchick", "kyogre", "evee"},
		},
		{
			input:    "WOOloo",
			expected: []string{"wooloo"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		//checks if the length of the slices matches, failure allows returning failed test early
		if len(actual) != len(c.expected) {
			t.Fatal("expected length of slice did not match actual slice")
		}

		//checks if failure input gracefully returns nil as expected
		if actual == nil && c.expected == nil {
			continue
		}

		//ensures that the slice values actually match
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected: %s\nActual: %v", expectedWord, word)
			}
		}
	}

}
