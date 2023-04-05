/*
Double Linked List implementation
*/
package fundamentals

import (
	"fmt"
	_ "fmt" //refer to : https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc
)

// Define a Node struct with Value and Next
type DoubleNode struct {
	Prev  *DoubleNode
	Next  *DoubleNode
	Value int
}

// Define a linked list with Head and Tail
type DoubleLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
	Size int
}

// Add elements to the last of the list
func (l *DoubleLinkedList) Append(value int) {
	newNode := &DoubleNode{Value: value}
	if l.Size == 0 {
		l.Head = newNode
		l.Tail = newNode
	}
	l.Tail.Next = newNode
	newNode.Prev = l.Tail
	l.Tail = newNode
	l.Size += 1
}

// Add elements to the front of the list
func (l *DoubleLinkedList) Prepend(value int) {
	newNode := &DoubleNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}

// Delete Head
func (l *DoubleLinkedList) DeleteHead() {
	if l.Size == 0 {
		return
	}
	if l.Size == 1 {
		l.Head = nil
		l.Tail = nil
	}
	l.Head.Next.Prev = nil
	l.Head = l.Head.Next
	l.Size--
}

// Delete Tail
func (l *DoubleLinkedList) DeleteTail() {
	if l.Size == 0 {
		return
	}
	if l.Size == 1 {
		l.Head = nil
		l.Tail = nil
	}
	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
	l.Size--
}

// Delete node from the double linked list
func (l *DoubleLinkedList) DeleteIndex(index int) {
	fmt.Println("Deleting position: ", index)
	if index == 0 {
		l.DeleteHead()
	}
	currentNode := l.Head
	for n := index - 2; n >= 0; n-- {
		currentNode = currentNode.Next
	}
	currentNode.Prev.Next = currentNode.Next
	l.Size--
}

// Show item
func (l *DoubleLinkedList) ShowItem() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println("---------------------")
}

// func main() {
// 	fmt.Println("start")
// 	list := DoubleLinkedList{}

// 	list.Append(3)
// 	list.Append(4)
// 	list.Append(5)
// 	list.Append(6)
// 	list.Append(7)
// 	list.ShowItem()
// 	list.Prepend(2)
// 	list.Prepend(1)
// 	list.ShowItem()
// 	list.DeleteHead()
// 	list.DeleteTail()
// 	list.ShowItem()
// 	list.DeleteIndex(3)
// 	list.ShowItem()

// }
