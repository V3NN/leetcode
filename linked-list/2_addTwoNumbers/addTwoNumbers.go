package __addTwoNumbers

/**
2. 两数相加

给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)
	curr := res
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		curr.Next = new(ListNode)
		curr = curr.Next

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Val = carry % 10
		carry = carry / 10
	}

	return res.Next
}
