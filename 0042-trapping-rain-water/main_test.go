package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	height   []int
	expected int
}

func TestTrap(t *testing.T) {
	cases := []TestCase{
		TestCase{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, expected: 6},
	}

	for _, cas := range cases {
		res := trap(cas.height)
		assert.Equal(t, cas.expected, res)
	}
}
