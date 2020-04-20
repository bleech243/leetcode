package leetcode

/*
	2. 两数相加
	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
	示例：
	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head = l1
	var p1, p2 *ListNode
	var scale = 0
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		scale, l1.Val = (l1.Val+l2.Val+scale)/10, (l1.Val+l2.Val+scale)%10
		p1 = l1
		p2 = l2
	}

	if l1 == nil && l2 == nil {
		if scale != 0 {
			p1.Next = &ListNode{scale, nil}
		}
		return head
	}

	if l1 == nil {
		p1.Next = l2
		for ; scale != 0 && l2 != nil; l2 = l2.Next {
			scale, l2.Val = (l2.Val+scale)/10, (l2.Val+scale)%10
			p2 = l2
		}
		if scale != 0 {
			p2.Next = &ListNode{scale, nil}
		}
		return head
	}

	if l2 == nil {
		for ; scale != 0 && l1 != nil; l1 = l1.Next {
			scale, l1.Val = (l1.Val+scale)/10, (l1.Val+scale)%10
			p1 = l1
		}
		if scale != 0 {
			p1.Next = &ListNode{scale, nil}
		}
		return head
	}
	return head
}