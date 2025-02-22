package task0021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists_01(list1 *ListNode, list2 *ListNode) *ListNode {
	append := func(head *ListNode, node *ListNode) {
		if head == nil {
			head = node
		} else {
			head.Next = node
		}
	}

	var head *ListNode
	current := head
	for list1 != nil || list2 != nil {
		var new ListNode

		if list1 == nil {
			new.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			new.Val = list1.Val
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			new.Val = list2.Val
			list2 = list2.Next
		} else {
			new.Val = list1.Val
			list1 = list1.Next
		}

		append(current, &new)
		if head == nil {
			head = &new
		}
		current = &new
	}

	return head
}
