package list

import (
	"sync"
)

type Node struct {
	next, prev *Node
	Value      interface{}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) GetValue() interface{} {
	return n.Value
}

type LinkedList struct {
	head, tail *Node
	count      int
	mutex      sync.Mutex
}

func NewLinkedList() *LinkedList {
	head := &Node{Value: nil}
	tail := &Node{Value: nil}

	head.next = tail
	tail.prev = head

	return &LinkedList{
		head:  head,
		tail:  tail,
		count: 0,
	}
}

func (l *LinkedList) Len() int {
	return l.count
}

func (l *LinkedList) Front() *Node {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	front := l.head.next

	if front == l.tail {
		return nil
	}

	return front
}

func (l *LinkedList) Back() *Node {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	back := l.tail.prev

	if back == l.head {
		return nil
	}

	return back
}

func (l *LinkedList) PushFront(insertNode *Node) *Node {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	head := l.head
	headNext := head.next
	head.next = insertNode
	insertNode.prev = head
	insertNode.next = headNext
	headNext.prev = insertNode
	l.count++

	return insertNode
}

func (l *LinkedList) PushBack(insertNode *Node) *Node {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	tail := l.tail
	tailPrev := tail.prev
	tailPrev.next = insertNode
	insertNode.prev = tailPrev
	insertNode.next = tail
	tail.prev = insertNode
	l.count++

	return insertNode
}

func (l *LinkedList) InsertBefore(insertNode *Node, mark *Node) *Node {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	prevNode := mark.prev
	prevNode.next = insertNode
	insertNode.prev = prevNode
	insertNode.next = mark
	mark.prev = insertNode

	l.count++

	return insertNode
}

func (l *LinkedList) InsertAfter(insertNode *Node, mark *Node) *Node {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	nextNode := mark.next
	nextNode.prev = insertNode
	insertNode.prev = mark
	insertNode.next = nextNode
	mark.next = insertNode

	l.count++

	return insertNode
}

func (l *LinkedList) Remove(removeItem *Node) *Node {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	removeItem.prev.next = removeItem.next
	removeItem.next.prev = removeItem.prev
	removeItem.next, removeItem.prev = nil, nil

	l.count--

	return removeItem
}
