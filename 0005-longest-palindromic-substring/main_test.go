package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  string
	output string
}

func TestLongestPalindrome(t *testing.T) {
	cases := []TestCase{
		TestCase{input: "babad", output: "bab"},
		TestCase{input: "cbbd", output: "bb"},
		TestCase{input: "abcba", output: "abcba"},
		TestCase{input: "aaaa", output: "aaaa"},
	}

	for _, cas := range cases {
		res := longestPalindrome(cas.input)
		assert.Equal(t, cas.output, res, "input: "+cas.input)
	}
}
