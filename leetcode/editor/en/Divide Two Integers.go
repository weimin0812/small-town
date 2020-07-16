package en

import (
	"math"
)

//Given two integers dividend and divisor, divide two integers without using mul

//tiplication, division and mod operator. 
//
// Return the quotient after dividing dividend by divisor. 
//
// The integer division should truncate toward zero, which means losing its frac
//tional part. For example, truncate(8.345) = 8 and truncate(-2.7335) = -2. 
//
// Example 1: 
//
// 
//Input: dividend = 10, divisor = 3
//Output: 3
//Explanation: 10/3 = truncate(3.33333..) = 3.
// 
//
// Example 2: 
//
// 
//Input: dividend = 7, divisor = -3
//Output: -2
//Explanation: 7/-3 = truncate(-2.33333..) = -2.
// 
//
// Note: 
//
// 
// Both dividend and divisor will be 32-bit signed integers. 
// The divisor will never be 0. 
// Assume we are dealing with an environment which could only store integers wit
//hin the 32-bit signed integer range: [−231, 231 − 1]. For the purpose of this pr
//oblem, assume that your function returns 231 − 1 when the division result overfl
//ows. 
// 
// Related Topics Math Binary Search 
// 👍 1190 👎 5236

//leetcode submit region begin(Prohibit modification and deletion)
func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return math.MaxInt32
	}
	flag := 1
	if divisor < 0 {
		flag *= -1
		divisor *= -1
	}
	if dividend < 0 {
		flag *= -1
		dividend *= -1
	}
	ret := 0
	for divisor <= dividend {
		t := divisor
		count := 1
		for t<<1 <= dividend {
			t <<= 1
			count <<= 1
		}
		ret += count
		dividend -= t
	}
	if ret > math.MaxInt32 {
		if flag == 1 {
			return math.MaxInt32
		}
		return math.MinInt32
	}
	return flag * ret
}

//leetcode submit region end(Prohibit modification and deletion)
