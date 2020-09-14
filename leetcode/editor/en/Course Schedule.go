package en

//There are a total of numCourses courses you have to take, labeled from 0 to nu
//mCourses-1. 
//
// Some courses may have prerequisites, for example to take course 0 you have to
// first take course 1, which is expressed as a pair: [0,1] 
//
// Given the total number of courses and a list of prerequisite pairs, is it pos
//sible for you to finish all courses? 
//
// 
// Example 1: 
//
// 
//Input: numCourses = 2, prerequisites = [[1,0]]
//Output: true
//Explanation: There are a total of 2 courses to take. 
//             To take course 1 you should have finished course 0. So it is poss
//ible.
// 
//
// Example 2: 
//
// 
//Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
//Output: false
//Explanation: There are a total of 2 courses to take. 
//             To take course 1 you should have finished course 0, and to take c
//ourse 0 you should
//             also have finished course 1. So it is impossible.
// 
//
// 
// Constraints: 
//
// 
// The input prerequisites is a graph represented by a list of edges, not adjace
//ncy matrices. Read more about how a graph is represented. 
// You may assume that there are no duplicate edges in the input prerequisites. 
//
// 1 <= numCourses <= 10^5 
// 
// Related Topics Depth-first Search Breadth-first Search Graph Topological Sort
// 
// 👍 4337 👎 189

//leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	in := make([]int, numCourses)
	for _, p := range prerequisites {
		in[p[0]]++
	}
	var q []int
	for i, d := range in {
		if d == 0 {
			q = append(q, i)
		}
	}
	count := 0
	for len(q) > 0 {
		count++
		c := q[0]
		q = q[1:]
		for _, p := range prerequisites {
			if p[1] == c {
				in[p[0]]--
				if in[p[0]] == 0 {
					q = append(q, p[0])
				}
			}
		}
	}
	return count == numCourses
}

//leetcode submit region end(Prohibit modification and deletion)
