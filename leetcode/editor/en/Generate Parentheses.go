package en

//
//Given n pairs of parentheses, write a function to generate all combinations of
// well-formed parentheses.
// 
//
// 
//For example, given n = 3, a solution set is:
// 
// 
//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
// Related Topics String Backtracking 
// 👍 5172 👎 267

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	ret := []string{}
	if n <= 0 {
		return ret
	}
	generateParenthesisHelper(&ret, "", 0, 0, n)
	return ret
}

func generateParenthesisHelper(ret *[]string, s string, open int, close int, n int) {
	if open == n && close == n {
		*ret = append(*ret, s)
		return
	}
	if open < n {
		s = s + "("
		generateParenthesisHelper(ret, s, open+1, close, n)
		s = s[:len(s)-1]
	}
	if close < open {
		s = s + ")"
		generateParenthesisHelper(ret, s, open, close+1, n)
		s = s[:len(s)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
