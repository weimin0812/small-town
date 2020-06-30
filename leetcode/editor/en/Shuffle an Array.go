package en

import (
	"math/rand"
)

//Shuffle a set of numbers without duplicates.
// 
//
// Example:
// 
//// Init an array with set 1, 2, and 3.
//int[] nums = {1,2,3};
//Solution solution = new Solution(nums);
//
//// Shuffle the array [1,2,3] and return its result. Any permutation of [1,2,3]
// must equally likely to be returned.
//solution.shuffle();
//
//// Resets the array back to its original configuration [1,2,3].
//solution.reset();
//
//// Returns the random shuffling of array [1,2,3].
//solution.shuffle();
// 
//

//leetcode submit region begin(Prohibit modification and deletion)
type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
	return Solution{original: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.original
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	length := len(this.original)
	shuffle := make([]int, length)
	copy(shuffle, this.original)
	for i, _ := range shuffle {
		index := i + rand.Intn(length-i)
		shuffle[i], shuffle[index] = shuffle[index], shuffle[i]
	}
	return shuffle
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
//leetcode submit region end(Prohibit modification and deletion)
