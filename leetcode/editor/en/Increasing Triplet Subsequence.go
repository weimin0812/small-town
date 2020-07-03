package en

import (
	"math"
)

//Given an unsorted array return whether an increasing subsequence of length 3 e
//xists or not in the array. 
//
// Formally the function should: 
//
// Return true if there exists i, j, k 
//such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false
//. 
//
// Note: Your algorithm should run in O(n) time complexity and O(1) space comple
//xity. 
//
// 
// Example 1: 
//
// 
//Input: [1,2,3,4,5]
//Output: true
// 
//
// 
// Example 2: 
//
// 
//Input: [5,4,3,2,1]
//Output: false
// 
// 
//

//leetcode submit region begin(Prohibit modification and deletion)
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	minOne, minTwo := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num < minOne {
			minOne = num
		}
		if minOne < num && num < minTwo {
			minTwo = num
		}
		if num > minTwo {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
//func increasingTriplet(nums []int) bool {
//	if len(nums) < 3 {
//		return false
//	}
//	dp := make([]int, len(nums))
//	for i, num := range nums {
//		dp[i] = 1
//		for j := 0; j < i; j++ {
//			if num > nums[j] {
//				dp[i] = max(dp[i], dp[j]+1)
//				if dp[i] > 2 {
//					return true
//				}
//			}
//		}
//		if dp[i] > 2 {
//			return true
//		}
//	}
//	return false
//}
//func max(a, b int) int {
//	if a >= b {
//		return a
//	}
//	return b
//}
