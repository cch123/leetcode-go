package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected int
}

func TestLongestConsecutive(t *testing.T) {
	cases := []TestCase{
		TestCase{nums: []int{100, 4, 200, 1, 3, 2}, expected: 4},
		TestCase{nums: []int{0}, expected: 1},
		TestCase{nums: []int{}, expected: 0},
		TestCase{nums: []int{0, 1, 2, 3, 4}, expected: 5},
		TestCase{nums: []int{4, 3, 2, -1, 1, 0}, expected: 6},
		TestCase{nums: []int{4, 3, 2, -1, 1, 0, 5}, expected: 7},
		TestCase{nums: []int{4, 4, 4, 4, 4, 4, 3, 2, -1, 1, 0, 5}, expected: 7},
		TestCase{nums: []int{-7, -1, 3, -9, -4, 7, -3, 2, 4, 9, 4, -9, 8, -7, 5, -1, -7}, expected: 4},
	}

	for _, cas := range cases {
		res := longestConsecutive(cas.nums)
		assert.Equal(t, cas.expected, res)
	}
}
