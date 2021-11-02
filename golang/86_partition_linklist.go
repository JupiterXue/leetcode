package main

func partition2(head *ListNode, x int) *ListNode {
	first := &ListNode{0, nil}
	firstPoint := first
	second := &ListNode{0, nil}
	secondPoint := second
	for head != nil {
		if head.Val >= x {
			tmp := &ListNode{head.Val, nil}
			secondPoint.Next = tmp
			secondPoint = secondPoint.Next
		} else {
			tmp := &ListNode{head.Val, nil}
			firstPoint.Next = tmp
			firstPoint = firstPoint.Next
		}
		head = head.Next
	}
	firstPoint.Next = second.Next

	return first.Next
}