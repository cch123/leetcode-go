package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected int
}

func TestMaxProduct(t *testing.T) {
	cases := []TestCase{
		TestCase{nums: []int{2, -5, -2, -4, 3}, expected: 24},
		TestCase{nums: []int{2, 3, -2, 4}, expected: 6},
		TestCase{nums: []int{-2, 0, -1}, expected: 0},
	}

	for _, cas := range cases {
		res := maxProduct(cas.nums)
		assert.Equal(t, cas.expected, res)
	}
}
