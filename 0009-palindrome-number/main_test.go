package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input    int
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	cases := []TestCase{
		TestCase{input: 0, expected: true},
		TestCase{input: 121, expected: true},
		TestCase{input: -121, expected: false},
	}

	for _, cas := range cases {
		res := isPalindrome(cas.input)
		assert.Equal(t, cas.expected, res, "input: "+fmt.Sprint(cas.input))
	}
}
