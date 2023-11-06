package main

func (l *LinkedList) MiddleOfLinkedList() (val int) {
	if l.head == nil {
		panic("Empty List")
	} else if l.head.Next() == nil {
		return l.head.Value()
	}
	hare := l.head
	tortoise := l.head

	for hare != nil && hare.Next() != nil {
		hare = hare.Next().Next()
		tortoise = tortoise.Next()
	}
	return tortoise.Value()
}
