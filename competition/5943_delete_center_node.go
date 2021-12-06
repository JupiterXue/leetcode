package main

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	len := 0
	point := head
	for point != nil {
		len++
		point = point.Next
	}
	point = head
	for i := 0; i < len/2 - 1; i++ {
		point = point.Next
	}
	if point.Next == nil {
		return nil
	}
	point.Next = point.Next.Next

	return head
}

