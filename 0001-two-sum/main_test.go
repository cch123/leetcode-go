package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Case test definition
type Case struct {
	nums           []int
	target         int
	expectedResult []int
}

func TestTwoSum(t *testing.T) {
	var cases = []Case{
		Case{nums: []int{1, 2, 3, 4, 5},
			target:         8,
			expectedResult: []int{2, 4},
		},
		Case{nums: []int{1, 2, 3, 4, 5},
			target:         9,
			expectedResult: []int{3, 4},
		},
		Case{nums: []int{3, 2, 4},
			target:         6,
			expectedResult: []int{1, 2},
		},
	}

	for _, cas := range cases {
		result := twoSum(cas.nums, cas.target)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, result[0], cas.expectedResult[0])
		assert.Equal(t, result[1], cas.expectedResult[1])
	}

}
