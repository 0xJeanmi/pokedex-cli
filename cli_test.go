package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input: "  ",
			expected: []string{},
		},
		{
			input: "  *",
			expected: []string{"*"},
		},
		{
			input: " 213              213 213 ",
			expected: []string{"213", "213", "213"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Error: expected %v but got %v", c.expected, actual)
			continue
		}
		
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error: expected %v but got %v", expectedWord, word)
				break
			}
		}

		fmt.Println("Test passed")
	}
}

