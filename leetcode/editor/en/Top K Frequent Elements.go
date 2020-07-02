package en

import (
	"sort"
)

//Given a non-empty array of integers, return the k most frequent elements.
//
// Example 1: 
//
// 
//Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
// 
//
// 
// Example 2: 
//
// 
//Input: nums = [1], k = 1
//Output: [1] 
// 
//
// Note: 
//
// 
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements. 
// Your algorithm's time complexity must be better than O(n log n), where n is t
//he array's size. 
// It's guaranteed that the answer is unique, in other words the set of the top 
//k frequent elements is unique. 
// You can return the answer in any order. 
// 
// Related Topics Hash Table Heap

//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	frequentMap := make(map[int]int)
	for i := range nums {
		frequentMap[nums[i]]++
	}
	var unique []int
	for k := range frequentMap {
		unique = append(unique, k)
	}
	sort.Slice(unique, func(i, j int) bool {
		return frequentMap[unique[i]] >= frequentMap[unique[j]]
	})
	var ret []int
	for i := 0; i < k; i++ {
		ret = append(ret, unique[i])
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
//func topKFrequent(nums []int, k int) []int {
//	frequentMap := make(map[int]int)
//	for i := range nums {
//		frequentMap[nums[i]]++
//	}
//	var unique []int
//	for k := range frequentMap {
//		unique = append(unique, k)
//	}
//	sort.Slice(unique, func(i, j int) bool {
//		return frequentMap[unique[i]] >= frequentMap[unique[j]]
//	})
//	var ret []int
//	for i := 0; i < k; i++ {
//		ret = append(ret, unique[i])
//	}
//	return ret
//}
