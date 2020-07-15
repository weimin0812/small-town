package en

//Given an array of integers nums sorted in ascending order, find the starting a
//nd ending position of a given target value. 
//
// Your algorithm's runtime complexity must be in the order of O(log n). 
//
// If the target is not found in the array, return [-1, -1]. 
//
// Example 1: 
//
// 
//Input: nums = [5,7,7,8,8,10], target = 8
//Output: [3,4] 
//
// Example 2: 
//
// 
//Input: nums = [5,7,7,8,8,10], target = 6
//Output: [-1,-1] 
//
// 
// Constraints: 
//
// 
// 0 <= nums.length <= 10^5 
// -10^9 <= nums[i] <= 10^9 
// nums is a non decreasing array. 
// -10^9 <= target <= 10^9 
// 
// Related Topics Array Binary Search 
// 👍 3487 👎 148

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	return []int{binarySearch(nums, target, true), binarySearch(nums, target, false)}
}

func binarySearch(nums []int, target int, findFirst bool) int {
	l, h := 0, len(nums)-1
	for l <= h {
		m := l + (h-l)/2
		if nums[m] == target {
			if findFirst {
				if m == 0 || nums[m-1] != target {
					return m
				}
				h = m - 1
			} else {
				if m == len(nums)-1 || nums[m+1] != target {
					return m
				}
				l = m + 1
			}
		} else if nums[m] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
