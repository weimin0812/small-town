package en

//Given a string, find the first non-repeating character in it and return its in
//dex. If it doesn't exist, return -1. 
//
// Examples: 
//
// 
//s = "leetcode"
//return 0.
//
//s = "loveleetcode"
//return 2.
// 
//
// 
//
// Note: You may assume the string contains only lowercase English letters. 
// Related Topics Hash Table String

//leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	var count [256]int
	for _, c := range s {
		count[c]++
	}
	for i, c := range s {
		if count[c] == 1 {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
