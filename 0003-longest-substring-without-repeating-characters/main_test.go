package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	s              string
	expectedResult int
}

func TestLengthfLongestSubstring(t *testing.T) {
	cases := []TestCase{
		TestCase{s: "abcabcbb", expectedResult: 3},
		TestCase{s: "bbbbb", expectedResult: 1},
		TestCase{s: "pwwkew", expectedResult: 3},
	}

	for _, cas := range cases {
		result := lengthfLongestSubstring(cas.s)
		assert.Equal(t, cas.expectedResult, result, "input: "+cas.s)
	}

}
