package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	A        [][]int
	expected [][]int
}

func TestFlipAndInvertImage(t *testing.T) {
	cases := []TestCase{
		TestCase{
			A:        [][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}},
			expected: [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}},
		},
		TestCase{
			A:        [][]int{[]int{1, 1, 0, 0}, []int{1, 0, 0, 1}, []int{0, 1, 1, 1}, []int{1, 0, 1, 0}},
			expected: [][]int{[]int{1, 1, 0, 0}, []int{0, 1, 1, 0}, []int{0, 0, 0, 1}, []int{1, 0, 1, 0}},
		},
	}

	for _, cas := range cases {
		res := flipAndInvertImage(cas.A)
		assert.Equal(t, cas.expected, res)
	}
}
