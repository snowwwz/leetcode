package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		arg1     []int
		arg2     int
		expected int
	}{
		{
			arg1:     []int{1, 3, 5, 6},
			arg2:     5,
			expected: 2,
		},
		{
			arg1:     []int{1, 3, 5, 6},
			arg2:     2,
			expected: 1,
		},
		{
			arg1:     []int{1, 3, 5, 6},
			arg2:     7,
			expected: 4,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, searchInsert(c.arg1, c.arg2))
	}
}
