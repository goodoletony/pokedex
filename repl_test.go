package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input: "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input: "     HELlo!!           ",
			expected: []string{"hello!!"},
		},
		{
			input: "   This is a tesT  ",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input: "  This     Is  Another           Test   ",
			expected: []string{"this", "is", "another", "test"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d, but got %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expecting %v, but got %v", expectedWord, word)
			}
		}
	}
}
