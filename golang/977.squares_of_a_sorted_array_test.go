package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		arg      []int
		expected []int
	}{
		{
			arg:      []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			arg:      []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, sortedSquares(c.arg))
	}
}
