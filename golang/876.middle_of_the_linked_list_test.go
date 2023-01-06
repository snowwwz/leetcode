package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleNode(t *testing.T) {
	new Linkedlist()
	cases := []struct {
		arg      string
		expected string
	}{
		{
			arg:      "Let's take LeetCode contest",
			expected: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			arg:      "God Ding",
			expected: "doG gniD",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, middleNode(c.arg))
	}
}
