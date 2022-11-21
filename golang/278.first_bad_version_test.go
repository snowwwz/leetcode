package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstBadVersion(t *testing.T) {
	cases := []struct {
		arg      int
		expected int
	}{
		{
			arg:      5,
			expected: 4,
		},
		{
			arg:      1,
			expected: 1,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, firstBadVersion(c.arg))
	}
}
