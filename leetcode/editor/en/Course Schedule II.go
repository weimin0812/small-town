package en

//There are a total of n courses you have to take labelled from 0 to n - 1.
//
// Some courses may have prerequisites, for example, if prerequisites[i] = [ai, 
//bi] this means you must take the course bi before the course ai. 
//
// Given the total number of courses numCourses and a list of the prerequisite p
//airs, return the ordering of courses you should take to finish all courses. 
//
// If there are many valid answers, return any of them. If it is impossible to f
//inish all courses, return an empty array. 
//
// 
// Example 1: 
//
// 
//Input: numCourses = 2, prerequisites = [[1,0]]
//Output: [0,1]
//Explanation: There are a total of 2 courses to take. To take course 1 you shou
//ld have finished course 0. So the correct course order is [0,1].
// 
//
// Example 2: 
//
// 
//Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
//Output: [0,2,1,3]
//Explanation: There are a total of 4 courses to take. To take course 3 you shou
//ld have finished both courses 1 and 2. Both courses 1 and 2 should be taken afte
//r you finished course 0.
//So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3
//].
// 
//
// Example 3: 
//
// 
//Input: numCourses = 1, prerequisites = []
//Output: [0]
// 
//
// 
// Constraints: 
//
// 
// 1 <= numCourses <= 2000 
// 0 <= prerequisites.length <= numCourses * (numCourses - 1) 
// prerequisites[i].length == 2 
// 0 <= ai, bi < numCourses 
// ai != bi 
// All the pairs [ai, bi] are distinct. 
// 
// Related Topics Depth-first Search Breadth-first Search Graph Topological Sort
// 
// 👍 2662 👎 141

//leetcode submit region begin(Prohibit modification and deletion)
func findOrder(numCourses int, prerequisites [][]int) []int {
	var ret []int
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
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		ret = append(ret, c)
		for _, p := range prerequisites {
			if p[1] == c {
				in[p[0]]--
				if in[p[0]] == 0 {
					q = append(q, p[0])
				}
			}
		}
	}
	if len(ret) == numCourses {
		return ret
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
