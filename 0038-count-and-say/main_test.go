package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	n        int
	expected string
}

func TestCountAndSay(t *testing.T) {
	cases := []TestCase{
		TestCase{n: 1, expected: "1"},
		TestCase{n: 4, expected: "1211"},
		TestCase{n: 5, expected: "111221"},
	}

	for _, cas := range cases {
		res := countAndSay(cas.n)
		assert.Equal(t, cas.expected, res)
	}
}
