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
