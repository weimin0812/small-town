package en

import (
	"math"
)

//Given an integer array nums, find the contiguous subarray (containing at least
// one number) which has the largest sum and return its sum. 
//
// Example: 
//
// 
//Input: [-2,1,-3,4,-1,2,1,-5,4],
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
// 
//
// Follow up: 
//
// If you have figured out the O(n) solution, try coding another solution using 
//the divide and conquer approach, which is more subtle. 
// Related Topics Array Divide and Conquer Dynamic Programming 
// 👍 8735 👎 412

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := nums[0]
	t := nums[0]
	for i := 1; i < len(nums); i++ {
		if t < 0 {
			t = 0
		}
		t += nums[i]
		ret = int(math.Max(float64(ret), float64(t)))
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
