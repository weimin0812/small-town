package en

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid. 
//
// An input string is valid if: 
//
// 
// Open brackets must be closed by the same type of brackets. 
// Open brackets must be closed in the correct order. 
// 
//
// Note that an empty string is also considered valid. 
//
// Example 1: 
//
// 
//Input: "()"
//Output: true
// 
//
// Example 2: 
//
// 
//Input: "()[]{}"
//Output: true
// 
//
// Example 3: 
//
// 
//Input: "(]"
//Output: false
// 
//
// Example 4: 
//
// 
//Input: "([)]"
//Output: false
// 
//
// Example 5: 
//
// 
//Input: "{[]}"
//Output: true
// 
// Related Topics String Stack 
// 👍 5037 👎 219

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	m := map[int32]int32{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []int32{}
	for _, c := range s {
		if v, ok := m[c]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
