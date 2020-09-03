package en

import (
	"sort"
)

//Given a collection of numbers that might contain duplicates, return all possib
//le unique permutations. 
//
// Example: 
//
// 
//Input: [1,1,2]
//Output:
//[
//  [1,1,2],
//  [1,2,1],
//  [2,1,1]
//]
// 
// Related Topics Backtracking 
// 👍 2108 👎 63

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	var t []int
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	permUniqHelp(nums, 0, t, &visited, &res)
	return res
}

func permUniqHelp(nums []int, index int, t []int, visited *[]bool, res *[][]int) {
	if index == len(nums) {
		temp := make([]int, len(t))
		copy(temp, t)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !((*visited)[i]) {
			if i > 0 && nums[i] == nums[i-1] && !((*visited)[i-1]) {
				continue
			}
			(*visited)[i] = true
			t = append(t, nums[i])
			permUniqHelp(nums, index+1, t, visited, res)
			t = t[:len(t)-1]
			(*visited)[i] = false
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
