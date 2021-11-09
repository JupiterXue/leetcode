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

//func removeElements(head *ListNode, val int) *ListNode {
//	newHead := &ListNode{Next: head}
//	for point := newHead; point.Next != nil; {
//		if point.Next.Val == val {
//			point.Next = point.Next.Next
//		} else {
//			point = point.Next
//		}
//	}
//	return newHead.Next
//}

//func removeElements(head *ListNode, val int) *ListNode {
//	var cur, pre *ListNode = head, nil
//	for cur != nil {
//		if cur.Val == val {
//			if pre == nil {
//				head = head.Next
//				cur = head
//			} else {
//				pre.Next = cur.Next
//				cur = cur.Next
//			}
//		} else {
//			pre = cur
//			cur = cur.Next
//		}
//	}
//	return head
//}


//func removeElements(head *ListNode, val int) *ListNode {
//	newH := &ListNode{Next: head}
//	point := newH
//	for point.Next != nil {
//		if point.Next.Val == val {
//			point = point.Next.Next
//		} else {
//			point = point.Next
//		}
//	}
//	return newH.Next
//}

//func removeElements(head *ListNode, val int) *ListNode {
//	newHead := &ListNode{Next: head}
//	for point := newHead.Next; point.Next != nil; {
//		if point.Next.Val == val {
//			point.Next = point.Next.Next
//		} else {
//			point = point.Next
//		}
//	}
//	return newHead.Next
//}

func removeElements(head *ListNode, val int) *ListNode {
	nHead := &ListNode{}
	nHead.Next = head
	point := nHead
	for point.Next != nil {
		if point.Next.Val == val {
			point.Next = point.Next.Next
		} else {
			point = point.Next
		}
	}
	return nHead.Next
}
