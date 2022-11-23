package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		arg1     []int
		arg2     int
		expected []int
	}{
		{
			arg1:     []int{1, 2, 3, 4, 5, 6, 7},
			arg2:     3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			arg1:     []int{-1, -100, 3, 99},
			arg2:     2,
			expected: []int{3, 99, -1, -100},
		},
	}

	for _, c := range cases {
		rotate(c.arg1, c.arg2)
		assert.Equal(t, c.expected, c.arg1)
	}
}
