package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode struct {
//	Val int
//	Next *ListNode
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	result := new(ListNode)
//	point := result
//	sum := 0
//	for {
//		if l1 != nil && l2 != nil{
//		 	sum += l1.Val + l2.Val
//			l1 = l1.Next
//			l2 = l2.Next
//		} else {
//			if l1 != nil {
//				sum += l1.Val
//				l1 = l1.Next
//			}
//
//			if l2 != nil {
//				sum += l2.Val
//				l2 = l2.Next
//			}
//		}
//		point.Val = sum % 10
//		sum /= 10
//		if l1 == nil && l2 == nil && sum == 0 {
//			break
//		} else {
//			point.Next = new(ListNode)
//			point = point.Next
//		}
//	}
//	return result
//}
//
//func PrintlistNode(l0 *ListNode, name string) {
//	var result []int
//	node := l0
//	for ;; {
//		if node == nil {
//			break
//		}
//		result = append(result, node.Val)
//		node = node.Next
//	}
//	fmt.Print(name + "=")
//	fmt.Println(result)
//}
//
//func createListNode(str string) *ListNode {
//	result := new(ListNode)
//	point := result
//	dataList := strings.Split(str, ",")
//	for _, v := range dataList {
//		value, _ := strconv.Atoi(v)
//		tmpNode := &ListNode{
//			value,
//			nil,
//		}
//		point.Next = tmpNode
//		point = point.Next
//	}
//	return result.Next
//}
//
//func main() {
//	//l1 := createListNode("2,4,3")
//	//l2 := createListNode("5,6,4")
//
//	l1 := createListNode("9,9,9,9,9,9,9")
//	l2 := createListNode("9,9,9,9")
//
//	PrintlistNode(l1, "l1")
//	PrintlistNode(l2, "l2")
//
//	PrintlistNode(addTwoNumbers(l1, l2), "result")
//	PrintlistNode(l1, "l1")
//	PrintlistNode(l2, "l2")
//}


func LinkListToArray(l *ListNode) []int {
	point := l
	res := []int{}
	for point != nil {
		//res = append([]int{point.Val}, res...) // 头插
		res = append(res, point.Val) // 尾插
		point = point.Next
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listL1 := LinkListToArray(l1)
	listL2 := LinkListToArray(l2)
	newListNode := &ListNode{0, nil}
	point := newListNode
	more := 0
	if len(listL1) > len(listL2) {
		for i := 0; i < len(listL1); i++ {
			tmp := &ListNode{0, nil}
			sum := 0
			if i > len(listL2) - 1 {
				sum = listL1[i] + more
			} else {
				sum = listL1[i] + listL2[i] + more
			}
			if more > 0 {
				more--
			}
			if sum > 9 {
				sum %= 10
				more++
			}
			tmp.Val = sum
			point.Next = tmp
			point = point.Next
		}
	} else {
		for i := 0; i < len(listL2); i++ {
			tmp := &ListNode{0, nil}
			sum := 0
			if i > len(listL1) - 1 {
				sum = listL2[i] + more
			} else {
				sum = listL1[i] + listL2[i] + more
			}
			if more > 0 {
				more--
			}
			if sum > 9 {
				sum %= 10
				more++
			}
			tmp.Val = sum
			point.Next = tmp
			point = point.Next
		}
	}
	if more > 0 {
		tmp := &ListNode{1, nil}
		point.Next = tmp
	}
	return newListNode.Next
}