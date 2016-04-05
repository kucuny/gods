package tree

import (
	"fmt"
	"github.com/kucuny/gods/queue"
	"sync"
)

type Comparer func(value1, value2 interface{}) bool

func IntegerComparer(value1, value2 interface{}) bool {
	return value1.(int) > value2.(int)
}

func Integer64Comparer(value1, value2 interface{}) bool {
	return value1.(int64) > value2.(int64)
}

func Float32Comparer(value1, value2 interface{}) bool {
	return value1.(float32) > value2.(float32)
}

func Float64Comparer(value1, value2 interface{}) bool {
	return value1.(float64) > value2.(float64)
}

type Node struct {
	left, right *Node
	value       interface{}
}

func NewNode(value interface{}) *Node {
	return &Node{value: value}
}

type BinarySearchTree struct {
	root     *Node
	count    int
	comparer Comparer
	mutex    *sync.Mutex
}

func NewBinarySearchTree(comparer Comparer) *BinarySearchTree {
	return &BinarySearchTree{
		root:     nil,
		count:    0,
		comparer: comparer,
		mutex:    new(sync.Mutex),
	}
}

func (b *BinarySearchTree) Len() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.count
}

func (b *BinarySearchTree) insert(node *Node, value interface{}) bool {
	if node == nil {
		return false
	}

	if b.comparer(node.value, value) {
		if node.left == nil {
			node.left = &Node{value: value}
			return true
		} else {
			return b.insert(node.left, value)
		}
	} else if !b.comparer(node.value, value) {
		if node.right == nil {
			node.right = &Node{value: value}
			return true
		} else {
			return b.insert(node.right, value)
		}
	}

	return false
}

func (b *BinarySearchTree) Insert(value interface{}) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	var isSuccess bool = false

	if b.root == nil {
		b.root = &Node{value: value}
		isSuccess = true
	} else {
		isSuccess = b.insert(b.root, value)
	}

	if isSuccess {
		b.count++
	}

	return isSuccess
}

func (b *BinarySearchTree) TraversePreOrder() {
	b.preOrder(b.root)
}

func (b *BinarySearchTree) TraversePostOrder() {
	b.postOrder(b.root)
}

func (b *BinarySearchTree) TraverseInOrder() {
	b.inOrder(b.root)
}

func (b *BinarySearchTree) TraverseLevelOrder() {
	b.levelOrder(b.root)
}

func (b *BinarySearchTree) preOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.value)
	b.preOrder(node.left)
	b.preOrder(node.right)
}

func (b *BinarySearchTree) postOrder(node *Node) {
	if node == nil {
		return
	}

	b.postOrder(node.left)
	b.postOrder(node.right)
	fmt.Println(node.value)
}

func (b *BinarySearchTree) inOrder(node *Node) {
	if node == nil {
		return
	}

	b.inOrder(node.left)
	fmt.Println(node.value)
	b.inOrder(node.right)
}

func (b *BinarySearchTree) levelOrder(node *Node) {
	if node == nil {
		return
	}

	q := queue.NewQueue()
	q.Push(node)

	fmt.Println(q.Len())

	for q.Len() > 0 {
		n := q.Pop()
		e := n.(*Node)

		fmt.Println(e.value)

		if node.left != nil {
			q.Push(node.left)
		}

		if node.right != nil {
			q.Push(node.right)
		}
	}
}
