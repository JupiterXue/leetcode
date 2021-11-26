package main

//func hasCycle(head *ListNode) bool {
//	n := 0
//	for head != nil {
//		head = head.Next
//		n++
//		if n > 10000 {
//			return true
//		}
//	}
//	return false
//}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
