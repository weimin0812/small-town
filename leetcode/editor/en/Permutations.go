package en

//Given a collection of distinct integers, return all possible permutations.
//
// Example: 
//
// 
//Input: [1,2,3]
//Output:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// 
// Related Topics Backtracking 
// 👍 4233 👎 109

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	var ret [][]int
	if len(nums) == 0 {
		return ret
	}
	var t []int
	visited := make([]bool, len(nums))
	permHelper(nums, 0, t, &visited, &ret)
	return ret
}

func permHelper(nums []int, index int, t []int, visited *[]bool, ret *[][]int) {
	if index == len(nums) {
		temp := make([]int, len(t))
		copy(temp, t)
		*ret = append(*ret, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !((*visited)[i]) {
			(*visited)[i] = true
			t = append(t, nums[i])
			permHelper(nums, index+1, t, visited, ret)
			t = t[:len(t)-1]
			(*visited)[i] = false
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
