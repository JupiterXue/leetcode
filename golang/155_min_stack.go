package main

import "fmt"

type MinStack struct {
	stack []int
}


func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	minValue := this.stack[0]
	for _, v := range this.stack {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	param4 := obj.GetMin()
	obj.Pop()
	param3 := obj.Top()
	fmt.Println(param4)
	fmt.Println(param3)
}
