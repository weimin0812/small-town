package en

//Given a string s, find the longest palindromic substring in s. You may assume
//that the maximum length of s is 1000. 
//
// Example 1: 
//
// 
//Input: "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.
// 
//
// Example 2: 
//
// 
//Input: "cbbd"
//Output: "bb"
// 
// Related Topics String Dynamic Programming 
// 👍 6883 👎 535

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}

	// dp[i][j] = dp[i+1][j-1]
	ret := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && (i+1 > j-1 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(ret) {
				ret = s[i : j+1]
			}
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
