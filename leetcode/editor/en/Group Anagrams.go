package en

import (
	"sort"
)

//Given an array of strings strs, group the anagrams together. You can return th
//e answer in any order. 
//
// An Anagram is a word or phrase formed by rearranging the letters of a differe
//nt word or phrase, typically using all the original letters exactly once. 
//
// 
// Example 1: 
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
//Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2: 
// Input: strs = [""]
//Output: [[""]]
// Example 3: 
// Input: strs = ["a"]
//Output: [["a"]]
// 
// 
// Constraints: 
//
// 
// 1 <= strs.length <= 104 
// 0 <= strs[i].length <= 100 
// strs[i] consists of lower-case English letters. 
// 
// Related Topics Hash Table String 
// 👍 3927 👎 195

//leetcode submit region begin(Prohibit modification and deletion)

type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	if len(strs) == 0 {
		return res
	}
	m := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		sort.Sort(sortRunes(runes))
		k := string(runes)
		v := m[k]
		v = append(v, str)
		m[k] = v
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
