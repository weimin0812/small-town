package en

//Given a matrix of m x n elements (m rows, n columns), return all elements of t
//he matrix in spiral order. 
//
// Example 1: 
//
// 
//Input:
//[
// [ 1, 2, 3 ],
// [ 4, 5, 6 ],
// [ 7, 8, 9 ]
//]
//Output: [1,2,3,6,9,8,7,4,5]
// 
//
// Example 2: 
// 
//Input:
//[
//  [1, 2, 3, 4],
//  [5, 6, 7, 8],
//  [9,10,11,12]
//]
//Output: [1,2,3,4,8,12,11,10,9,5,6,7]
// Related Topics Array 
// 👍 2658 👎 589

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	row := len(matrix)
	if len(matrix[0]) == 0 {
		return res
	}
	col := len(matrix[0])
	l, r, t, b := 0, col-1, 0, row-1
	for l <= r && t <= b {
		// top row
		for i := l; i <= r; i++ {
			res = append(res, matrix[t][i])
		}
		// right column
		for i := t + 1; i <= b; i++ {
			res = append(res, matrix[i][r])
		}
		// bottom row
		if b != t {
			for i := r - 1; i >= l; i-- {
				res = append(res, matrix[b][i])
			}
		}
		// left column
		if l != r {
			for i := b - 1; i > t; i-- {
				res = append(res, matrix[i][l])
			}
		}
		l++
		r--
		t++
		b--
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
