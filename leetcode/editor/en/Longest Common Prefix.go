package en

import (
	"sort"
)

//Write a function to find the longest common prefix string amongst an array of
//strings. 
//
// If there is no common prefix, return an empty string "". 
//
// Example 1: 
//
// 
//Input: ["flower","flow","flight"]
//Output: "fl"
// 
//
// Example 2: 
//
// 
//Input: ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.
// 
//
// Note: 
//
// All given inputs are in lowercase letters a-z. 
// Related Topics String 
// 👍 2554 👎 1857

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]
	ret := ""
	for i := 0; i < min(len(first), len(last)); i++ {
		if first[i] == last[i] {
			ret += first[i : i+1]
		} else {
			break
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
