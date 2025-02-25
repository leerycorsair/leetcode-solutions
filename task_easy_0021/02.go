package task0021

func mergeTwoLists_02(list1 *ListNode, list2 *ListNode) *ListNode {
	append := func(head *ListNode, node *ListNode) {
		if head == nil {
			head = node
		} else {
			head.Next = node
		}
	}

	var head ListNode
	current := &head
	for list1 != nil && list2 != nil {
		var new ListNode
		if list1.Val > list2.Val {
			new.Val = list2.Val
			list2 = list2.Next
		} else {
			new.Val = list1.Val
			list1 = list1.Next
		}

		append(current, &new)
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return head.Next
}
