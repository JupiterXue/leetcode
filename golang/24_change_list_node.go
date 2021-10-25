package main

import (
	"fmt"
	"strconv"
	"strings"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h1, h2 := head, head.Next
	h1.Next = swapPairs(h2.Next)
	h2.Next = h1
	return h2

	//if head == nil {
	//	return head
	//}
	//
	//if head.Next == nil {
	//	return head
	//}
	//
	//nodeList := make([]*ListNode, 0)
	//for head != nil && head.Next != nil {
	//	nodeList = append(nodeList, head)
	//	head = head.Next.Next
	//}
	//
	////nextNode := new(ListNode)
	//nextNode := (*ListNode)(nil)
	//for i := len(nodeList) -1; i >= 0; i-- {
	//	h1, h2 := nodeList[i], nodeList[i].Next
	//	h1.Next = nextNode
	//	h2.Next = h1
	//	nextNode = h2
	//}
	//return nextNode
}

type ListNode struct {
	Val int
	Next *ListNode
}

func PrintlistNode(l0 *ListNode, name string) {
	var result []int
	node := l0
	for ;; {
		if node == nil {
			break
		}
		result = append(result, node.Val)
		node = node.Next
	}
	fmt.Print(name + "=")
	fmt.Println(result)
}

func createListNode(str string) *ListNode {
	result := new(ListNode)
	point := result
	dataList := strings.Split(str, ",")
	for _, v := range dataList {
		value, _ := strconv.Atoi(v)
		tmpNode := &ListNode{
			value,
			nil,
		}
		point.Next = tmpNode
		point = point.Next
	}
	return result.Next
}

func main() {
	head := createListNode("1,2,3,4")

	PrintlistNode(head, "head")

	PrintlistNode(swapPairs(head), "result")
	PrintlistNode(head, "head")
}

