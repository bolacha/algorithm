package linkedlist

import "fmt"

type (
	Node struct {
		value int
		next  *Node
	}
	LinkedList struct {
		head   *Node
		length int
	}
)

func (l *LinkedList) Prepend(value int) {
	newNode := Node{value, nil}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
	return
}

func (l *LinkedList) Append(value int) {
	newNode := Node{value, nil}

	if l.head == nil {
		l.head = &newNode
		l.length++
		return
	}

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = &newNode
	l.length++
	return
}

func (l *LinkedList) PrintList() {
	if l.head == nil {
		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%d->", currentNode.value)
		currentNode = currentNode.next
	}

	fmt.Print("\n")
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.length == 0 || l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousNode := l.head
	for previousNode.next.value != value {
		if previousNode.next.next == nil {
			return
		}

		previousNode = previousNode.next
	}

	previousNode.next = previousNode.next.next
	l.length--
	return
}
