package main

type Node struct {
	Prev *Node
	Next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Insert(key interface{}) {
	list := &Node{
		Next: l.head,
		key:  key,
	}

	if l.head != nil {
		l.head.Prev = list
	}

	l.head = list

	List := l.head

	for list.Next != nil {
		List = list.Next
	}

	l.tail = List
}
