package en

//Given four lists A, B, C, D of integer values, compute how many tuples (i, j,
//k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero. 
//
// To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ 
//N ≤ 500. All integers are in the range of -228 to 228 - 1 and the result is guar
//anteed to be at most 231 - 1. 
//
// Example: 
//
// 
//Input:
//A = [ 1, 2]
//B = [-2,-1]
//C = [-1, 2]
//D = [ 0, 2]
//
//Output:
//2
//
//Explanation:
//The two tuples are:
//1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
//2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
// 
//
// 
// Related Topics Hash Table Binary Search

//leetcode submit region begin(Prohibit modification and deletion)
func fourSumCount(A []int, B []int, C []int, D []int) int {
	abMap, cdMap := make(map[int]int), make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			abMap[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			cdMap[c+d]++
		}
	}
	ret := 0
	for abk, abCount := range abMap {
		if cdCount, ok := cdMap[-abk]; ok {
			ret += abCount * cdCount
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
