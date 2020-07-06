package en

import (
	"math"
)

//Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1: 
//
// 
//Input: 123
//Output: 321
// 
//
// Example 2: 
//
// 
//Input: -123
//Output: -321
// 
//
// Example 3: 
//
// 
//Input: 120
//Output: 21
// 
//
// Note: 
//Assume we are dealing with an environment which could only store integers with
//in the 32-bit signed integer range: [−231, 231 − 1]. For the purpose of this pro
//blem, assume that your function returns 0 when the reversed integer overflows. 
// Related Topics Math 
// 👍 3390 👎 5343

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	ret := 0
	for x != 0 {
		digit := x % 10
		if ret > math.MaxInt32/10 || ret == math.MaxInt32/10 && digit > math.MaxInt32%10 {
			return 0
		}
		if ret < math.MinInt32/10 || ret == math.MinInt32/10 && digit < math.MinInt32%10 {
			return 0
		}
		ret = ret*10 + digit
		x = x / 10
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
