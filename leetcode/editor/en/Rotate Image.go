package en

//You are given an n x n 2D matrix representing an image, rotate the image by 90
// degrees (clockwise). 
//
// You have to rotate the image in-place, which means you have to modify the inp
//ut 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation. 
//
// 
// Example 1: 
//
// 
//Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
//Output: [[7,4,1],[8,5,2],[9,6,3]]
// 
//
// Example 2: 
//
// 
//Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
//Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
// 
//
// Example 3: 
//
// 
//Input: matrix = [[1]]
//Output: [[1]]
// 
//
// Example 4: 
//
// 
//Input: matrix = [[1,2],[3,4]]
//Output: [[3,1],[4,2]]
// 
//
// 
// Constraints: 
//
// 
// matrix.length == n 
// matrix[i].length == n 
// 1 <= n <= 20 
// -1000 <= matrix[i][j] <= 1000 
// 
// Related Topics Array 
// 👍 3418 👎 251

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	col := len(matrix[0])
	if row != col {
		return
	}
	for i := 0; i < row; i++ {
		for j := i + 1; j < col; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	halfCol := col / 2
	for i := 0; i < row; i++ {
		for j := 0; j < halfCol; j++ {
			matrix[i][j], matrix[i][col-j-1] = matrix[i][col-j-1], matrix[i][j]
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
