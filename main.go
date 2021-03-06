package main

import (
	"errors"
	"fmt"
)

/*
In simple words, we can say, linked list is a
collection of nodes. Node consists of two parts:

	1. Data
	2. Pointer

*/

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
	len int
}



func (l *LinkedList) Append(val int)  {
	newNode := Node{}
	newNode.value = val
	if l.len ==0{
		l.head = &newNode
		l.len++
		return
	}
	pointer := l.head
	for i := 0; i<l.len; i++ {
		if pointer.next == nil {
			pointer.next =  &newNode
			l.len++
			return
		}
		pointer = pointer.next
	}

}
func (l *LinkedList) PrintList() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	pointer := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", pointer.value)
		pointer = pointer.next
	}

}

func (l *LinkedList) Search(val int) int {
	pointer := l.head
	for i := 0; i < l.len; i++ {
		if pointer.value == val {
			return i // or just return value
		}
		pointer = pointer.next
	}
	return -1
}

// GetAt returns node at given position from linked list
func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (l *LinkedList) DeleteAt(pos int) error {
	// validate the position
	if pos < 0 {
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("No nodes in list")
	}
	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = l.GetAt(pos).next
	fmt.Println("prevNode: ",prevNode)
	fmt.Println("prevNode.next: ",prevNode.next)
	fmt.Println("l.len before decrementing: ", l.len)
	l.len-- //reduce the length after removing
	fmt.Println("l.len after decrementing: ", l.len)

	return nil
}

func (l *LinkedList) DeleteVal(val int) error {
	ptr := l.head
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = ptr.next
			}
			l.len--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}
///check well
func (l *LinkedList) Prepend(val int) {
	newNode := &Node{value: val}
	if l.head == nil {
		l.head = newNode
	}
	newNode.next = l.head
	l.head = newNode
}
// InsertAt Insert a node after a given node
func (l *LinkedList) InsertAt(val int, pos int) {
	// 10 2 3 : 1 => 10, 15, 2, 3
	newNode := &Node{value: val}
	if pos == 0 {
		newNode.next = l.head
		l.head = newNode
	}
	count := 0
	current := l.head
	for current != nil {
		if count == pos-1 {
			newNode.next = current.next
			current.next = newNode
			break
		}
		count++
		current = current.next
	}
}

func main()  {
	l := LinkedList{}
	l.Append(15)
	l.Append(27)
	l.Append(45)
	l.PrintList()
	fmt.Println(l.Search(27))
	l.PrintList()
	fmt.Println(l.GetAt(1))
	l.DeleteAt(1)
	l.PrintList()

}