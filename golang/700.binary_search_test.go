package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		arg1     []int
		arg2     int
		expected int
	}{
		{
			arg1:     []int{-1, 0, 3, 5, 9, 12},
			arg2:     9,
			expected: 4,
		},
		{
			arg1:     []int{-1, 0, 3, 5, 9, 12},
			arg2:     2,
			expected: -1,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, search(c.arg1, c.arg2))
	}
}
