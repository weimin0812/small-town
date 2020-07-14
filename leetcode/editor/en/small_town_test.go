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

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	lis := lengthOfLIS(nums)
	assert.True(t, lis == 4)
}

func TestTriplet(t *testing.T) {
	nums := []int{2, 4, -2, -3}
	triplet := increasingTriplet(nums)
	assert.True(t, !triplet)
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	m := findMedianSortedArrays(nums1, nums2)
	assert.True(t, m == 2.5)
}

func TestIsValid(t *testing.T) {
	s := "()"
	valid := isValid(s)
	assert.True(t, valid)
}

func TestLetterCombinations(t *testing.T) {
	s := "23"
	combinations := letterCombinations(s)
	assert.True(t, len(combinations) == 9)
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	ret := removeDuplicates(nums)
	assert.True(t, ret == 5)
}

func TestStrStr(t *testing.T) {
	haystack := "a"
	needle := "a"
	index := strStr(haystack, needle)
	assert.True(t, index == 2)
}