package en

//Merge k sorted linked lists and return it as one sorted list. Analyze and desc
//ribe its complexity. 
//
// Example: 
//
// 
//Input:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//Output: 1->1->2->3->4->4->5->6
// 
// Related Topics Linked List Divide and Conquer Heap 
// 👍 4721 👎 288

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}
	d, l := 1, len(lists)
	for d < l {
		for i := 0; i < l-d; i += 2 * d {
			lists[i] = mergeTwo(lists[i], lists[i+d])
		}
		d = d * 2
	}
	return lists[0]
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	current := ret
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return ret.Next
}

//leetcode submit region end(Prohibit modification and deletion)
