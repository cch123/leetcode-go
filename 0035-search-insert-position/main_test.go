package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	target   int
	expected int
}

func TestSearchInsert(t *testing.T) {
	cases := []TestCase{
		TestCase{nums: []int{1, 3, 5, 6}, target: 5, expected: 2},
		TestCase{nums: []int{1, 3, 5, 6}, target: 2, expected: 1},
		TestCase{nums: []int{1, 3, 5, 6}, target: 7, expected: 4},
		TestCase{nums: []int{1, 3, 5, 6}, target: 0, expected: 0},
	}

	for _, cas := range cases {
		assert.Equal(t, cas.expected, searchInsert(cas.nums, cas.target), "input"+fmt.Sprint(cas.nums)+"target"+fmt.Sprint(cas.target))
	}

}
