package main

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) isEmpty() bool {
	return l.head == nil
}

func (l *LinkedList) prepend(value int) *LinkedList {

	l.head = &Node{
		value: value,
		next:  l.head,
	}

	return l
}

func (l *LinkedList) append(value int) *LinkedList {

	if l.head == nil {
		l.head = &Node{value: value}
		return l
	}

	tail := l.head

	for tail.next != nil {
		tail = tail.next
	}

	tail.next = &Node{value: value}

	return l
}

func (l *LinkedList) count() int {
	var count int

	for tail := l.head; tail != nil; count++ {
		tail = tail.next
	}

	return count
}

func (l *LinkedList) first() int {
	if l.head == nil {
		return 0
	}

	return l.head.value
}

func (l *LinkedList) last() int {
	if l.head == nil {
		return 0
	}

	tail := l.head

	for tail.next != nil {
		tail = tail.next
	}

	return tail.value
}
