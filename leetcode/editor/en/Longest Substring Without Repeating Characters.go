package en

//Given a string, find the length of the longest substring without repeating cha
//racters. 
//
// 
// Example 1: 
//
// 
//Input: "abcabcbb"
//Output: 3 
//Explanation: The answer is "abc", with the length of 3. 
// 
//
// 
// Example 2: 
//
// 
//Input: "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
// 
//
// 
// Example 3: 
//
// 
//Input: "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3. 
//             Note that the answer must be a substring, "pwke" is a subsequence
// and not a substring.
// 
// 
// 
// 
// Related Topics Hash Table Two Pointers String Sliding Window 
// 👍 9398 👎 568

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	ret, l, r := 0, 0, 0
	m := make(map[byte]int)
	for r < len(s) {
		if index, ok := m[s[r]]
			ok {
			l = max(l, index+1)
		}
		m[s[r]] = r
		ret = max(ret, r-l+1)
		r++
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
