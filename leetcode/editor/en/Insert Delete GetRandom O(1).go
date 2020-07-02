package en
//Design a data structure that supports all following operations in average O(1)
// time. 
//
// 
//
// 
// insert(val): Inserts an item val to the set if not already present. 
// remove(val): Removes an item val from the set if present. 
// getRandom: Returns a random element from current set of elements (it's guaran
//teed that at least one element exists when this method is called). Each element 
//must have the same probability of being returned. 
// 
//
// 
//
// Example: 
//
// 
//// Init an empty set.
//RandomizedSet randomSet = new RandomizedSet();
//
//// Inserts 1 to the set. Returns true as 1 was inserted successfully.
//randomSet.insert(1);
//
//// Returns false as 2 does not exist in the set.
//randomSet.remove(2);
//
//// Inserts 2 to the set, returns true. Set now contains [1,2].
//randomSet.insert(2);
//
//// getRandom should return either 1 or 2 randomly.
//randomSet.getRandom();
//
//// Removes 1 from the set, returns true. Set now contains [2].
//randomSet.remove(1);
//
//// 2 was already in the set, so return false.
//randomSet.insert(2);
//
//// Since 2 is the only number in the set, getRandom always return 2.
//randomSet.getRandom();
// 
// Related Topics Array Hash Table Design


//leetcode submit region begin(Prohibit modification and deletion)
type RandomizedSet struct {
    
}


/** Initialize your data structure here. */
func Constructor1() RandomizedSet {
	return RandomizedSet{
	}

}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	return false
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	return false
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return 0
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)
