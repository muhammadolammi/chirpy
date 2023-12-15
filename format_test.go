package main

import (
	"fmt"
	"testing"
)

func TestFormatter(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "hi this is kerfuffle",
			expected: "hi this is ****",
		}, {
			input:    "hi this is fornax",
			expected: "hi this is ****",
		},
		{
			input:    "hi this is sharbert",
			expected: "hi this is ****",
		},
		{
			input:    "hi this is sharbert , kerfuffle , fornax",
			expected: "hi this is **** , **** , ****",
		},
		{
			input:    "hi this is sharbert!",
			expected: "hi this is sharbert!",
		},
		{
			input:    "hi this is sharbert,",
			expected: "hi this is sharbert,",
		},
		{
			input:    "hi this is Sharbert",
			expected: "hi this is ****",
		},
	}
	for _, cs := range cases {
		actual := formatString(cs.input)
		if actual != cs.expected {
			t.Errorf("actual is not expected. actual: %v, expected : %v", actual, cs.expected)
			fmt.Println(actual)
		}

	}

}
