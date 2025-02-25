package task0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers_01(l1 *ListNode, l2 *ListNode) *ListNode {
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
	for l1 != nil && l2 != nil {
		digit := (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10

		new := ListNode{Val: digit}
		append(current, &new)

		current = current.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if carry == 0 {
			current.Next = l1
			break
		}

		mod := (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10

		new := ListNode{Val: mod}
		append(current, &new)

		current = current.Next
		l1 = l1.Next
	}

	for l2 != nil {
		if carry == 0 {
			current.Next = l2
			break
		}

		mod := (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10

		new := ListNode{Val: mod}
		append(current, &new)

		current = current.Next
		l2 = l2.Next
	}

	if carry != 0 {
		new := ListNode{Val: carry}
		append(current, &new)
		current = current.Next
	}

	return result.Next
}
