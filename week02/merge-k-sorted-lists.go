package week02

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
题目: 合并K个升序链表
主体思路: K个链表进行排序 => 多次进行两个升序链表的排序
*/

//ListNode 单链表
type ListNode struct {
	Val  int
	Next *LinkedNode
}

//KListNode 单链表数组
type KListNode struct {
	lists []*ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	answer := &ListNode{Val: -100000, Next: nil}
	kln := KListNode{lists: lists}
	kln.mergeTwoLists(0, answer)
	return answer.Next
}

// 对两个列表进行排序
func (this *KListNode) mergeTwoLists(idx int, l2 *ListNode) {
	protect := &ListNode{Val: -100000, Next: nil}
	ans := protect

	if idx == len(this.lists) {
		return
	}

	l1 := this.lists[idx]

	for l2 != nil && l1 != nil {
		if l2.Val >= l1.Val {
			ans.Next = l1
			l1 = l1.Next
		} else {
			ans.Next = l2
			l2 = l2.Next
		}
		ans = ans.Next
	}

	// 将较长的链表尾部加到ans尾部
	if l2 != nil {
		ans.Next = l2
	}
	if l1 != nil {
		ans.Next = l1
	}
	// 进行下一链表的排序
	this.mergeTwoLists(idx+1, protect.Next)
}
