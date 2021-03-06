package en

//Given an unsorted integer array, find the smallest missing positive integer.
//
// Example 1: 
//
// 
//Input: [1,2,0]
//Output: 3
// 
//
// Example 2: 
//
// 
//Input: [3,4,-1,1]
//Output: 2
// 
//
// Example 3: 
//
// 
//Input: [7,8,9,11,12]
//Output: 1
// 
//
// Note: 
//
// Your algorithm should run in O(n) time and uses constant extra space. 
// Related Topics Array 
// 👍 3492 👎 798

//leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	l := len(nums)
	for i := 0; i < l; {
		num := nums[i]
		if 1 <= num && num <= l {
			if num != nums[num-1] {
				nums[i], nums[num-1] = nums[num-1], nums[i]
			} else {
				i++
			}
		} else {
			i++
		}
	}
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return l + 1
}

//leetcode submit region end(Prohibit modification and deletion)
