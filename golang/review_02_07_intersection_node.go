package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//if headA == nil || headB == nil {
	//	return nil
	//}
	pointA, pointB := headA, headB
	for pointA != pointB {
		if pointA == nil {
			pointA = headB
		} else {
			pointA = pointA.Next
		}

		if pointB == nil {
			pointB = headA
		} else {
			pointB = pointB.Next
		}
	}
	return pointA
}
