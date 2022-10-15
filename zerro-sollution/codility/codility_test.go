package codility_test

import (
	"testing"

	code "always-learning/zerro-sollution/codility"

	"github.com/stretchr/testify/assert"
)

func TestBinaryGap(t *testing.T) {
	// assert.Equal(t, 2, code.BinaryGap(1))
	// assert.Equal(t, 2, code.BinaryGap(9))
	assert.Equal(t, 4, code.BinaryGap(529))
	// assert.Equal(t, 0, code.BinaryGap(32))
	assert.Equal(t, 5, code.BinaryGap(1041))
	// assert.Equal(t, 5, code.BinaryGap(1376796946))
}

func TestCyrcilcRotation(t *testing.T) {
	var expect []int
	// code.CyrcilcRotation([]int{3, 8, 9, 7, 6}, 3)
	expect = []int{9, 7, 6, 3, 8}
	assert.Equal(t, expect, code.CyrcilcRotation([]int{3, 8, 9, 7, 6}, 3))
	expect = []int{5, 1, 2, 3, 4}
	assert.Equal(t, expect, code.CyrcilcRotation([]int{1, 2, 3, 4, 5}, 6))
}
