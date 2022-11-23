package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		arg      []int
		expected []int
	}{
		{
			arg:      []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			arg:      []int{0},
			expected: []int{0},
		},
	}

	for _, c := range cases {
		moveZeroes(c.arg)
		assert.Equal(t, c.expected, c.arg)
	}
}
