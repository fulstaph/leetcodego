package leetcodego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTripletSubsequence(t *testing.T) {
	// Test cases
	assert.True(t, increasingTriplet([]int{1, 2, 3}))
	assert.False(t, increasingTriplet([]int{5, 4, 3, 2, 1}))
	assert.True(t, increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	assert.False(t, increasingTriplet([]int{2, 1, 5, 0, 4, 3}))
}
