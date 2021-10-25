package main

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//
//	newHead := &ListNode{}
//	newPoint := newHead
//	count := 0
//	point := head
//	for point.Next != nil {
//		count++
//		point = point.Next
//	}
//
//	fmt.Println(count)
//
//	for count >= 0 {
//		tmpPoint := head
//		for i := 0; i <= count; i++ {
//			tmpPoint = tmpPoint.Next
//		}
//		newPoint.Next = &ListNode{tmpPoint.Val, nil}
//		newPoint = newPoint.Next
//		count--
//	}
//	return newHead.Next
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var cur *ListNode
	pre := head
	for pre != nil {
		tmp := pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
	}
	return cur
}
