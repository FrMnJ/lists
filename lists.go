package lists

import (
	"fmt"
	"strings"
)

// Generic Node for List implementation
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

// Returns the node as String
func (n Node[T]) String() string {
	return fmt.Sprintf("(%v)", n.Value)
}

// Generic Circular Doubly Linked List
type List[T comparable] struct {
	sentinel *Node[T]
	length   uint
}

// Create a new empty Linked List
func New[T comparable]() *List[T] {
	sentinel := new(Node[T])
	sentinel.Next = sentinel
	sentinel.Prev = sentinel
	return &List[T]{
		sentinel: sentinel,
		length:   0,
	}
}

// Get a pointer to the first element of the Linked List
func (l *List[T]) Head() *Node[T] {
	if l.length == 0 {
		return nil
	}
	return l.sentinel.Next
}

// Return a pointer to the last element of the Linked List
func (l *List[T]) Tail() *Node[T] {
	if l.length == 0 {
		return nil
	}
	return l.sentinel.Prev
}

// Return the number of elements in the Linked List
func (l List[T]) Length() uint {
	return l.length
}

// Add a node to the back of the Linked List
func (l *List[T]) Add(value T) {
	newNode := new(Node[T])
	newNode.Value = value
	newNode.Next = l.sentinel
	newNode.Prev = l.sentinel.Prev
	l.sentinel.Prev.Next = newNode
	l.sentinel.Prev = newNode
	l.length++
}

// Add a node to the beginning of the Linked List
func (l *List[T]) Insert(value T) {
	newNode := new(Node[T])
	newNode.Value = value
	newNode.Prev = l.sentinel
	newNode.Next = l.sentinel.Next
	l.sentinel.Next.Prev = newNode
	l.sentinel.Next = newNode
	l.length++
}

// Remove a the first occurrence of a node with the Value provided
// If the Value is not found, return false
// Otherwise, return true
func (l *List[T]) Remove(value T) bool {
	for iterator := l.sentinel.Next; iterator.Next != nil; iterator = iterator.Next {
		if iterator.Value == value {
			iterator.Next.Prev = iterator.Prev
			iterator.Prev.Next = iterator.Next
			l.length--
			return true
		}
	}
	return false
}

// Return a pointer to the node that contain the Value provided
// If not found, return nil
func (l List[T]) Find(value T) *Node[T] {
	iterator := l.sentinel.Next
	for ; iterator.Next != l.sentinel; iterator = iterator.Next {
		if iterator.Value == value {
			return iterator
		}
	}
	return nil
}

// Convert the Linked List to a string
func (l List[T]) String() string {
	sb := strings.Builder{}
	iterator := l.sentinel.Next
	for ; iterator.Next != l.sentinel; iterator = iterator.Next {
		sb.WriteString(fmt.Sprintf("%v->", iterator))
	}
	sb.WriteString(fmt.Sprintf("%v", iterator))
	return sb.String()
}
