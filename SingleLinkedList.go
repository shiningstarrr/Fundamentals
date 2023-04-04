/*
Single Linked List implementation
*/

package fundamentals

import (
	"fmt"
)

// Define a Node struct with Value and Next
type SingleNode struct {
	Value int
	Next  *SingleNode
}

// Define a Linked List with Head and Tail
type SingleLinkedList struct {
	Head *SingleNode
	Tail *SingleNode
}

// Add elements to the last of the list
func (l *SingleLinkedList) Append(value int) {
	newNode := &SingleNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	}
	l.Tail.Next = newNode
	l.Tail = newNode
}

// Add elemetns to the front of the list
func (l *SingleLinkedList) Prepend(value int) {
	newNode := &SingleNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	}
	newNode.Next = l.Head
	l.Head = newNode

}

// Delete node from single linked list
func (l *SingleLinkedList) Delete(value int) *SingleLinkedList {
	fmt.Println("Expect deleted", value)
	newList := &SingleLinkedList{}
	if l.Head == nil {
		return newList
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			currentNode.Next = currentNode.Next.Next
			break
		}
		currentNode = currentNode.Next
	}
	return l
}

// Display elements from the list
func (l *SingleLinkedList) showList() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}
