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
}

//func swapPairs(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	h1, h2 := head, head.Next
//	h1.Next = swapPairs(h2.Next)
//	h2.Next = h1
//	return h2
//}

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

