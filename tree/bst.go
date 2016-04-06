package tree

import (
	"github.com/kucuny/gods/queue"
	"sync"
)

type Comparer func(value1, value2 interface{}) bool
type Runner func(value interface{})

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

func (b *BinarySearchTree) TraversePreOrder(runner Runner) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.preOrder(b.root, runner)
}

func (b *BinarySearchTree) TraversePostOrder(runner Runner) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.postOrder(b.root, runner)
}

func (b *BinarySearchTree) TraverseInOrder(runner Runner) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.inOrder(b.root, runner)
}

func (b *BinarySearchTree) TraverseLevelOrder(runner Runner) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.levelOrder(b.root, runner)
}

func (b *BinarySearchTree) preOrder(node *Node, runner Runner) {
	if node == nil {
		return
	}

	runner(node.value)
	b.preOrder(node.left, runner)
	b.preOrder(node.right, runner)
}

func (b *BinarySearchTree) postOrder(node *Node, runner Runner) {
	if node == nil {
		return
	}

	b.postOrder(node.left, runner)
	b.postOrder(node.right, runner)
	runner(node.value)
}

func (b *BinarySearchTree) inOrder(node *Node, runner Runner) {
	if node == nil {
		return
	}

	b.inOrder(node.left, runner)
	runner(node.value)
	b.inOrder(node.right, runner)
}

func (b *BinarySearchTree) levelOrder(node *Node, runner Runner) {
	if node == nil {
		return
	}

	q := queue.NewQueue()
	q.Push(node)

	for q.Len() > 0 {
		n := (*Node)(q.Pop().(*Node))

		runner(n.value)

		if n.left != nil {
			q.Push(n.left)
		}

		if n.right != nil {
			q.Push(n.right)
		}
	}
}
