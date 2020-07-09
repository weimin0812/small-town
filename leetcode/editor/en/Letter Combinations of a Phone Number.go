package en

//Given a string containing digits from 2-9 inclusive, return all possible lette
//r combinations that the number could represent. 
//
// A mapping of digit to letters (just like on the telephone buttons) is given b
//elow. Note that 1 does not map to any letters. 
//
// 
//
// Example: 
//
// 
//Input: "23"
//Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 
//
// Note: 
//
// Although the above answer is in lexicographical order, your answer could be i
//n any order you want. 
// Related Topics String Backtracking 
// 👍 3889 👎 411

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	ret := []string{}
	if len(digits) == 0 {
		return ret
	}
	numbers := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "qprs", "tuv", "wxyz"}
	helper(digits, numbers, 0, &ret, "")
	return ret
}

func helper(digits string, numbers []string, index int, ret *[]string, s string) {
	if index == len(digits) {
		*ret = append(*ret, s)
		return
	}
	chars := numbers[digits[index]-'0']
	for i := range chars {
		s = s + chars[i:i+1]
		helper(digits, numbers, index+1, ret, s)
		s = s[:len(s)-1]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
