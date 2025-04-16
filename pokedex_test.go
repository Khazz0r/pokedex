package main

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
			input:    " PIKACHU MUST go",
			expected: []string{"pikachu", "must", "go"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) < len(c.expected) {
			t.Errorf("actual slice is not the same length as the expected slice")
			t.Fail()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual word does not match the expected word")
				t.Fail()
			}
		}
	}

}
