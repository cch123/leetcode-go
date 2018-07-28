package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected int
}

func TestFindPeakElement(t *testing.T) {
	cases := []TestCase{
		TestCase{nums: []int{2, 1}, expected: 0},
		TestCase{nums: []int{1, 2}, expected: 1},
		TestCase{nums: []int{1, 2, 3}, expected: 2},
		TestCase{nums: []int{3, 2, 1}, expected: 0},
		TestCase{nums: []int{1, 2, 3, 2, 1}, expected: 2},
		TestCase{nums: []int{1, 2, 2, 3, 2, 1}, expected: 3},
		TestCase{nums: []int{1, 2, 3, 1}, expected: 2},
		TestCase{nums: []int{1, 2, 1, 3, 5, 6, 4}, expected: 1},
	}

	for _, cas := range cases {
		res := findPeakElement(cas.nums)
		assert.Equal(t, cas.expected, res)
	}

}
