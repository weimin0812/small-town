package en

//Implement pow(x, n), which calculates x raised to the power n (i.e. xn).
//
// 
// Example 1: 
//
// 
//Input: x = 2.00000, n = 10
//Output: 1024.00000
// 
//
// Example 2: 
//
// 
//Input: x = 2.10000, n = 3
//Output: 9.26100
// 
//
// Example 3: 
//
// 
//Input: x = 2.00000, n = -2
//Output: 0.25000
//Explanation: 2-2 = 1/22 = 1/4 = 0.25
// 
//
// 
// Constraints: 
//
// 
// -100.0 < x < 100.0 
// -231 <= n <= 231-1 
// -104 <= xn <= 104 
// 
// Related Topics Math Binary Search 
// 👍 1711 👎 3228

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	ret := 1.0
	if n < 0 {
		n = -n
		if x != 0 {
			x = 1 / x
		}
	}
	for n != 0 {
		if n%2 == 1 {
			ret = ret * x
		}
		x *= x
		n /= 2
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
