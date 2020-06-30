package en

//Calculate the sum of two integers a and b, but you are not allowed to use the
//operator + and -. 
//
// 
// Example 1: 
//
// 
//Input: a = 1, b = 2
//Output: 3
// 
//
// 
// Example 2: 
//
// 
//Input: a = -2, b = 3
//Output: 1
// 
// 
// 
// Related Topics Bit Manipulation

//leetcode submit region begin(Prohibit modification and deletion)
func getSum(a int, b int) int {
	sum, carry := a^b, (a&b)<<1
	for carry != 0 {
		t := sum
		sum = sum ^ carry
		carry = (carry & t) << 1
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
