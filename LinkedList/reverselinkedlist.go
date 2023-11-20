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

func (l* LinkedList) reverseRecursiveUtil(head *Node) *Node {
	if head==nil || head.next==nil{
		return head 
	}
	new_head := l.reverseRecursiveUtil(head.next)
	head.next.next = head
	head.next = nil
	return new_head
}

func (l* LinkedList) ReverseRecursive(){
	if l.length < 2{
		return
	}
	l.head=l.reverseRecursiveUtil(l.head)
}