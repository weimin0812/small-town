package en

//Given n non-negative integers representing an elevation map where the width of
// each bar is 1, compute how much water it is able to trap after raining. 
//
// 
//The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In 
//this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos
// for contributing this image! 
//
// Example: 
//
// 
//Input: [0,1,0,2,1,0,1,3,2,1,2,1]
//Output: 6 
// Related Topics Array Two Pointers Stack 
// 👍 7196 👎 122

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	l, r := 0, len(height)-1
	lMax, rMax := 0, 0
	ret := 0
	for l < r {
		if height[l] <= height[r] {
			if height[l] >= lMax {
				lMax = height[l]
			} else {
				ret += lMax - height[l]
			}
			l++
		} else {
			if height[r] >= rMax {
				rMax = height[r]
			} else {
				ret += rMax - height[r]
			}
			r--
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
