package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA.Next ==nil && headB.Next ==nil {
		if headA.Val == headB.Val {
			return headA
		} else {
			return nil
		}
	}

	APoint, BPoint := headA, headB
	for APoint != BPoint {
		if APoint == nil {
			APoint = headB
		} else {
			APoint = APoint.Next
		}

		if BPoint == nil {
			BPoint = headA
		} else {
			BPoint = BPoint.Next
		}
	}
	return APoint
}
