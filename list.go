package list

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList) print() {
	curr := l.head
	for l.length != 0 {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
		l.length--
	}

	fmt.Println()
}

func (l *LinkedList) delete(value int) {
	prev := l.head
	if l.length == 0 {
		return
	}

	if prev.data == value {
		l.head = l.head.next
	}

	for prev.next.data != value {
		if prev.next.next == nil {
			return
		}
		prev = prev.next
	}

	prev.next = prev.next.next
}
