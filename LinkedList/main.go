package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (n *Node) Value() int {
	return n.data
}

func (n *Node) Next() (next *Node) {
	next = n.next
	return next
}

func NewList() *LinkedList {
	list := &LinkedList{head: nil, length: 0}
	return list
}

func NewNode(value int) *Node {
	return &Node{data: value, next: nil}
}

func (l *LinkedList) AppendNode(n *Node) {
	if n == nil {
		return
	}

	if l.head == nil {
		l.head = n
		return
	}

	curr := l.head

	for curr.next != nil {
		curr = curr.Next()
	}

	curr.next = n
}

func (l *LinkedList) Add(value int) {
	node := NewNode(value)
	l.AppendNode(node)
	l.length++
}

func (l *LinkedList) Remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.Value() == value {
		l.head = l.head.Next()
		l.length--
		return
	}

	curr := l.head

	for curr.Next() != nil && curr.Next().Value() != value {
		curr = curr.Next()
	}

	if curr.Next() != nil {
		curr.next = curr.Next().Next()
	}

	l.length--
}

func main() {
	list := NewList()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	fmt.Println("Middle of Linked List:", list.MiddleOfLinkedList())
	printList(list)
	list.Remove(3)
	printList(list)
	list.Remove(1)
	printList(list)
	fmt.Println(list.head)

}

func printList(l *LinkedList) {
	curr := l.head
	fmt.Println("")
	for ; curr != nil; curr = curr.Next() {
		fmt.Printf("%v", curr.Value())
	}
	fmt.Println("")
}
