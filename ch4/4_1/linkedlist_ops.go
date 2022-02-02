package main

func (l *LinkedList) AddToList(n Node) {
	n.next = &l.head
	l.head = n
}

func (l *LinkedList) IsEmpty() bool {
	if &l.head == nil {
		return true
	}
	return false
}

//dequeue
func (l *LinkedList) RemoveFirst() Node {
	return l.tail
}
