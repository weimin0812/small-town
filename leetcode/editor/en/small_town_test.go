package en

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSumCount(t *testing.T) {
	a := []int{1, 2}
	b := []int{-2, -1}
	c := []int{-1, 2}
	d := []int{0, 2}
	count := fourSumCount(a, b, c, d)
	assert.True(t, count == 2)
}

func TestFizzBuzz(t *testing.T) {
	n := 15
	buzz := fizzBuzz(n)
	assert.True(t, len(buzz) == n)
}

func TestTopKFrequent(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	k := 10
	frequent := topKFrequent(nums, k)
	assert.True(t, len(frequent) == k)
}