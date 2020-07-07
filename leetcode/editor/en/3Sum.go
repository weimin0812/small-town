package en

import (
	"sort"
)

//Given an array nums of n integers, are there elements a, b, c in nums such tha
//t a + b + c = 0? Find all unique triplets in the array which gives the sum of ze
//ro. 
//
// Note: 
//
// The solution set must not contain duplicate triplets. 
//
// Example: 
//
// 
//Given array nums = [-1, 0, 1, 2, -1, -4],
//
//A solution set is:
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 
// Related Topics Array Two Pointers 
// 👍 6947 👎 814

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			l, h := i+1, len(nums)-1
			for l < h {
				sum := nums[i] + nums[l] + nums[h]
				if sum == 0 {
					ret = append(ret, []int{nums[i], nums[l], nums[h]})
					l++
					h--
					for l < h && nums[l] == nums[l-1] {
						l++
					}
					for l < h && nums[h] == nums[h+1] {
						h--
					}
				} else if sum > 0 {
					h--
				} else {
					l++
				}
			}
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
