package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	N        int
	expected int
}

func TestConsecutiveNumbersSum(t *testing.T) {
	cases := []TestCase{
		TestCase{N: 1, expected: 1},
		TestCase{N: 2, expected: 1},
		TestCase{N: 3, expected: 2},
		TestCase{N: 5, expected: 2},
		TestCase{N: 9, expected: 3},
		TestCase{N: 15, expected: 4},
	}

	for _, cas := range cases {
		res := consecutiveNumbersSum(cas.N)
		assert.Equal(t, cas.expected, res, "input: "+fmt.Sprint(cas.N))
	}
}
