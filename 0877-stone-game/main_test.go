package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	piles    []int
	expected bool
}

func TestStoneGame(t *testing.T) {
	cases := []TestCase{
		TestCase{piles: []int{5, 3, 4, 5}, expected: true},
		TestCase{piles: []int{5}, expected: true},
		TestCase{piles: []int{5, 2}, expected: true},
		TestCase{piles: []int{5, 6, 5}, expected: true},
	}

	for _, cas := range cases {
		res := stoneGame(cas.piles)
		assert.Equal(t, cas.expected, res, "input: "+fmt.Sprint(cas.piles))
	}

}
