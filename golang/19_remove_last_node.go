package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	point := head
	nodeLen := 0
	for point != nil {
		nodeLen++
		point = point.Next
	}

	newHead := &ListNode{0, head}
	newPoint := newHead
	for i := 0; i < nodeLen - n; i++ {
		newPoint = newPoint.Next
	}
	newPoint.Next = newPoint.Next.Next
	return newHead.Next
}