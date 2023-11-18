package main

func (l* LinkedList) ReverseWithAStack(){
	if l.length < 2 {
		return
	}
	curr := l.head
	stack := make([](*Node),0,l.length)
	for curr!=nil {
		stack = append(stack, curr)
		curr = curr.Next()
	}
	j := len(stack)-1
	l.head = stack[j]
	curr = l.head
	j--
	for j>-1 {
		curr.next = nil
		curr.next = stack[j]
		curr=curr.next
		j--
	}
	curr.next = nil
}

func (l* LinkedList) ReverseWithPointers(){
	if l.length < 2{
		return
	}

	var prev *Node = nil
	var curr *Node = l.head
	var next *Node = curr.next

	for curr!=nil{
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.head = prev

}