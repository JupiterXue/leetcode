package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	point := head
	for point.Next != nil {
		if point.Val == point.Next.Val {
			point.Next = point.Next.Next
		} else {
			point = point.Next
		}
	}
	return head
}