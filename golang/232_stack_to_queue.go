package main

import "fmt"

type MyQueue struct {
	queue []int
}


func Constructor2() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
	this.queue = append(this.queue, x)
}


func (this *MyQueue) Pop() int {
	res := this.queue[0]
	this.queue = this.queue[1:]
	return res
	//if len(this.queue) > 0 {
	//	res := this.queue[0]
	//	this.queue = this.queue[1:]
	//	return res
	//} else {
	//	return nil
	//}
}


func (this *MyQueue) Peek() int {
	return this.queue[0]
}


func (this *MyQueue) Empty() bool {
	if len(this.queue) == 0 {
		return true
	} else {
		return false
	}
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor2()
	obj.Push(1)
	obj.Push(1)
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}
