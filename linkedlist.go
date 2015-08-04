// Package algostudy provides algostudy
package algostudy

import (
	"errors"
)

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList() LinkedList {
	return LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *LinkedList) InsertHead(data string) {
	newNode := Node{
		Data: data,
		Next: l.head,
	}
	l.head = &newNode
	l.length += 1
}

func (l *LinkedList) InsertTail(data string) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}
	l.tail.Next = newNode
	l.tail = newNode
	l.length += 1
}

func (l *LinkedList) Insert(data string, p int) error {
	var prev, cur *Node
	cur = l.head

	if p < 0 {
		return errors.New("Postition must be a positive integer")
	}

	if cur == nil {
		l.InsertHead(data)
		l.tail = l.head
		return nil
	}

	if p == 0 {
		l.InsertHead(data)
		return nil
	}

	if p == l.length && p > 0 {
		l.InsertTail(data)
		return nil
	}

	prev = cur
	cur = cur.Next
	newNode := Node{
		Data: data,
		Next: nil,
	}

	i := 1
	for cur != nil {
		if p == i {
			newNode.Next = cur
			prev.Next = &newNode
			l.length += 1
			return nil
		}
		prev = cur
		cur = cur.Next
		i++
	}

	return errors.New("Overflow Error!")
}

func (l *LinkedList) RemoveNodeByValue(data string) error {
	var cur, prev *Node

	cur = l.head
	for cur != nil {
		if cur.Data == data {
			if cur == l.head {
				l.head = cur.Next
				l.length -= 1
				return nil
			}

			prev.Next = cur.Next
			l.length -= 1
			return nil
		}
		prev = cur
		cur = cur.Next
	}
	return errors.New("Element not found in list")
}

func (l *LinkedList) RemoveHead() {
	l.head = l.head.Next
	l.length -= 1
}

func (l *LinkedList) RemoveTail() {
	var cur, prev *Node
	cur = l.head
	for cur.Next != nil {
		prev = cur
		cur = cur.Next
	}
	prev.Next = nil
	l.tail = prev
	l.length -= 1
}

func (l *LinkedList) RemoveByPosition(p int) error {
	var cur, prev *Node
	cur = l.head

	if p < 0 {
		return errors.New("Postition must be a positive integer")
	}

	if cur == nil {
		return errors.New("List Empty")
	}

	if p == 0 {
		l.RemoveHead()
		return nil
	}

	if p == l.length-1 {
		l.RemoveTail()
		return nil
	}

	prev = cur
	cur = cur.Next

	i := 1
	for cur != nil {
		if p == i {
			prev.Next = cur.Next
			l.length -= 1
			return nil
		}
		prev = cur
		cur = cur.Next
		i++
	}

	return errors.New("Hit end of list before desired position.")
}

func (l *LinkedList) Get(p int) (string, error) {
	var cur *Node
	cur = l.head

	if p < 0 {
		return "", errors.New("Position must be a positive integer.")
	}

	i := 0
	for cur != nil {
		if p == i {
			return cur.Data, nil
		}
		cur = cur.Next
		i++
	}

	return "", errors.New("Hit end of list before provided index.")
}

func (l *LinkedList) GetHead() string {
	return l.head.Data
}

func (l *LinkedList) GetTail() string {
	return l.tail.Data
}
