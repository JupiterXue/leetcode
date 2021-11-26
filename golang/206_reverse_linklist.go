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

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	var cur *ListNode
//	pre := head
//	for pre != nil {
//		tmp := pre.Next
//		pre.Next = cur
//		cur = pre
//		pre = tmp
//	}
//	return cur
//}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	// 双指针
//	var cur *ListNode
//	pre := head
//	for pre != nil {
//		tmp := pre.Next
//		pre.Next = cur
//		cur = pre
//		pre = tmp
//	}
//	return cur
//}

//func reverseList(head *ListNode) *ListNode {
//	if head == nil ||head.Next == nil {
//		return head
//	}
//
//	point := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//	return point
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	point := reverseList(head.Next) // 定位到最后个节点作为头
	head.Next.Next = head // 指针反向
	head.Next = nil // 最后个节点指空
	return point
}
