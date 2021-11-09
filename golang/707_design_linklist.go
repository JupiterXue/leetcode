package main

type MyLinkedList struct {
	val int
	next *MyLinkedList
}


func Constructor() MyLinkedList {
	return MyLinkedList{0, nil}
}


func (this *MyLinkedList) Get(index int) int {
	point := this
	for index > 0 {
		index--
		point = point.next
	}
	return point.val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	this = &MyLinkedList{val, this}
	//this.next = tmp
}


func (this *MyLinkedList) AddAtTail(val int)  {
	point := this
	for point.next != nil {
		point = point.next
	}
	tmp := &MyLinkedList{val, nil}
	point.next = tmp
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	point := this
	for index > 0 {
		index--
		point = point.next
	}
	tmp := &MyLinkedList{val, point.next}
	point.next = tmp
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	point := this
	for index > 0 {
		index--
		point = point.next
	}
	point.next = point.next.next
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
