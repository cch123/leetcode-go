package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	piles    []int
	H        int
	expected int
}

func TestMinEatingSpeed(t *testing.T) {
	cases := []TestCase{
		TestCase{piles: []int{312884470}, H: 968709470, expected: 1},
		TestCase{piles: []int{3, 6, 7, 11}, H: 8, expected: 4},
		TestCase{piles: []int{30, 11, 23, 4, 20}, H: 5, expected: 30},
		TestCase{piles: []int{30, 11, 23, 4, 20}, H: 6, expected: 23},
		TestCase{piles: []int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}, H: 823855818, expected: 14},
	}

	for _, cas := range cases {
		res := minEatingSpeed(cas.piles, cas.H)
		assert.Equal(t, cas.expected, res)
	}
}
