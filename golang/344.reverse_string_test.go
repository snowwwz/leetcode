package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		arg      []string
		expected []string
	}{
		{
			arg:      []string{"h", "e", "l", "l", "o"},
			expected: []string{"o", "l", "l", "e", "h"},
		},
		{
			arg:      []string{"H", "a", "n", "n", "a", "h"},
			expected: []string{"h", "a", "n", "n", "a", "H"},
		},
	}

	for _, c := range cases {
		reverseString(c.arg)
		assert.Equal(t, c.expected, c.arg)
	}
}
