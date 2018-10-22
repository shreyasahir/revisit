package main

import (
	"fmt"
)

//Node is single entity of linklist
type Node struct {
	Next  *Node
	Value int
}

//List is linklist
type List struct {
	Head *Node
}

func (l *List) insertAsecnding(val int) {
	n := &Node{Value: val}
	if l.Head == nil {
		l.Head = n
	} else {
		curr := l.Head
		var prev *Node
		for curr.Value < n.Value && curr != nil {
			prev = curr
			curr = curr.Next
		}
		if prev == nil {
			l.Head = n
		} else {
			prev.Next = n
		}
		if curr != nil {
			n.Next = curr
		}
	}
}

func (l *List) insert(val int) {
	n := &Node{Value: val}
	if l.Head == nil {
		l.Head = n
	} else {
		curr := l.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = n
	}

}

func (l *List) printList() {
	curr := l.Head

	for curr != nil {
		fmt.Println("Link list", curr.Value)
		curr = curr.Next
	}
}
