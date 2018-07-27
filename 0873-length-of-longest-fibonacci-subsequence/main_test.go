package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected int
}

func TestLenLongestFibSubseq(t *testing.T) {
	cases := []TestCase{
		TestCase{nums: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: 5},
		TestCase{nums: []int{1, 3, 7, 11, 12, 14, 18}, expected: 3},
	}

	for _, cas := range cases {
		res := lenLongestFibSubseq(cas.nums)
		assert.Equal(t, cas.expected, res, "input: "+fmt.Sprint(cas.nums))
	}
}
