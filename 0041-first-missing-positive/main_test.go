package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected int
}

func TestFirstMissingPositive(t *testing.T) {
	cases := []TestCase{
		TestCase{nums: []int{}, expected: 1},
		TestCase{nums: []int{1}, expected: 2},
		TestCase{nums: []int{1, 2, 0}, expected: 3},
		TestCase{nums: []int{3, 4, -1, 1}, expected: 2},
		TestCase{nums: []int{7, 8, 9, 11, 12}, expected: 1},
	}

	for _, cas := range cases {
		res := firstMissingPositive(cas.nums)
		assert.Equal(t, cas.expected, res)
	}
}
