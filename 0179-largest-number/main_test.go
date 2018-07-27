package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected string
}

func TestLargestNumber(t *testing.T) {
	cases := []TestCase{
		TestCase{nums: []int{0, 0}, expected: "0"},
		TestCase{nums: []int{121, 12}, expected: "12121"},
		TestCase{nums: []int{128, 12}, expected: "12812"},
		TestCase{nums: []int{2, 1}, expected: "21"},
		TestCase{nums: []int{20, 1}, expected: "201"},
		TestCase{nums: []int{10, 2}, expected: "210"},
		TestCase{nums: []int{3, 30, 34, 5, 9}, expected: "9534330"},
	}

	for _, cas := range cases {
		res := largestNumber(cas.nums)
		assert.Equal(t, cas.expected, res)
	}
}
