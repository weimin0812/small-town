package en

import (
	"math"
)

//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
// Find the median of the two sorted arrays. The overall run time complexity sho
//uld be O(log (m+n)). 
//
// You may assume nums1 and nums2 cannot be both empty. 
//
// Example 1: 
//
// 
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
// 
//
// Example 2: 
//
// 
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5
// 
// Related Topics Array Binary Search Divide and Conquer 
// 👍 7106 👎 1101

//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	l1, l2 := len(nums1), len(nums2)
	half := (l1 + l2 + 1) / 2
	l, h := 0, l1
	for l <= h {
		i := l + (h-l)/2
		j := half - i
		if i > 0 && nums1[i-1] > nums2[j] {
			h = i - 1
		} else if i < l1 && nums2[j-1] > nums1[i] {
			l = i + 1
		} else {
			lmax := 0.0
			if i == 0 {
				lmax = float64(nums2[j-1])
			} else if j == 0 {
				lmax = float64(nums1[i-1])
			} else {
				lmax = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (l1+l2)%2 == 1 {
				return lmax
			}
			rmin := 0.0
			if i == l1 {
				rmin = float64(nums2[j])
			} else if j == l2 {
				rmin = float64(nums1[i])
			} else {
				rmin = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			return (lmax + rmin) / 2.0

		}
	}
	return 0

}

//leetcode submit region end(Prohibit modification and deletion)
