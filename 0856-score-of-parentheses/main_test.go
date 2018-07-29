package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	S        string
	expected int
}

func TestScoreOfParentheses(t *testing.T) {
	cases := []TestCase{
		TestCase{S: "(())", expected: 2},
		TestCase{S: "(()(()))", expected: 6},
		TestCase{S: "()", expected: 1},
		TestCase{S: "()()", expected: 2},
	}

	for _, cas := range cases {
		res := scoreOfParentheses(cas.S)
		assert.Equal(t, cas.expected, res, "input: "+fmt.Sprint(cas.S))
	}
}
