package en

//Implement strStr().
//
// Return the index of the first occurrence of needle in haystack, or -1 if need
//le is not part of haystack. 
//
// Example 1: 
//
// 
//Input: haystack = "hello", needle = "ll"
//Output: 2
// 
//
// Example 2: 
//
// 
//Input: haystack = "aaaaa", needle = "bba"
//Output: -1
// 
//
// Clarification: 
//
// What should we return when needle is an empty string? This is a great questio
//n to ask during an interview. 
//
// For the purpose of this problem, we will return 0 when needle is an empty str
//ing. This is consistent to C's strstr() and Java's indexOf(). 
// Related Topics Two Pointers String 
// 👍 1600 👎 1912

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	i, index := 0, 0

	for ; i < len(haystack); {
		if haystack[i] == needle[index] {
			i++
			index++
		} else {
			i = i - index + 1
			index = 0
		}

		if index == len(needle) {
			return i - index
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
