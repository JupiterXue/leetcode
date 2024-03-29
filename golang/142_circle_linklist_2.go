package main

//func detectCycle(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//	slow, fast := head, head
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//		if slow == fast {
//			slow = head
//			for slow != fast {
//				slow = slow.Next
//				fast = fast.Next
//			}
//			return slow
//		}
//	}
//	return nil
//}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head // 重置
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
