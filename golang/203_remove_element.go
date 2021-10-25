package main

// 没对
//func removeElements(head *ListNode, val int) *ListNode {
//	point := new(ListNode)
//	point.Next = head
//	for point.Next != nil {
//		if point.Next.Val == val {
//			if point.Next == nil {
//				point = nil
//			} else {
//				point = point.Next.Next
//			}
//			break
//		}
//		point = point.Next
//	}
//	return point
//}

func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{Next: head}
	for point := newHead; point.Next != nil; {
		if point.Next.Val == val {
			point.Next = point.Next.Next
		} else {
			point = point.Next
		}
	}
	return newHead.Next
}
