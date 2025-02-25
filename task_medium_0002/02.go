package task0002

func addTwoNumbers_02(l1 *ListNode, l2 *ListNode) *ListNode {
	append := func(head *ListNode, node *ListNode) {
		if head == nil {
			head = node
		} else {
			head.Next = node
		}
	}

	var result ListNode
	current := &result
	var carry = 0
	for l1 != nil || l2 != nil || carry != 0 {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		digit := (val1 + val2 + carry) % 10
		carry = (val1 + val2 + carry) / 10

		new := ListNode{Val: digit}
		append(current, &new)
		current = current.Next
	}

	return result.Next
}
