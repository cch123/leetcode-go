package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected int
}

func TestMaxSubArray(t *testing.T) {
	cases := []TestCase{
		TestCase{
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		TestCase{
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		TestCase{
			nums:     []int{-1},
			expected: -1,
		},
	}

	for _, cas := range cases {
		res := maxSubArray(cas.nums)
		assert.Equal(t, cas.expected, res)
	}
}
