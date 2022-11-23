package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		arg1     []int
		arg2     int
		expected []int
	}{
		{
			arg1:     []int{2, 7, 11, 15},
			arg2:     9,
			expected: []int{1, 2},
		},
		{
			arg1:     []int{2, 3, 4},
			arg2:     6,
			expected: []int{1, 3},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, twoSum(c.arg1, c.arg2))
	}
}
