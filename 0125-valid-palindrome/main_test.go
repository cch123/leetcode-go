package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	s        string
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	cases := []TestCase{
		TestCase{s: "A man, a plan, a canal: Panama", expected: true},
		TestCase{s: "race a car", expected: false},
		TestCase{s: "", expected: true},
	}
	for _, cas := range cases {
		res := isPalindrome(cas.s)
		assert.Equal(t, cas.expected, res, "input: "+cas.s)
	}
}
