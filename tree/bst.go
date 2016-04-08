package tree

import (
	"github.com/kucuny/gods/queue"
	"sync"
)

type Comparer func(source, target interface{}) bool
type Runner func(value *Node)

func IntegerComparer(source, target interface{}) bool {
	return source.(int) > target.(int)
}

func Integer64Comparer(source, target interface{}) bool {
	return source.(int64) > target.(int64)
}

func Float32Comparer(source, target interface{}) bool {
	return source.(float32) > target.(float32)
}

func Float64Comparer(source, target interface{}) bool {
	return source.(float64) > target.(float64)
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

func (b *BinarySearchTree) Search(value interface{}) *Node {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.root == nil {
		return nil
	}

	return b.search(b.root, value)
}

func (b *BinarySearchTree) TraversePreOrder(runner Runner) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.preOrder(b.root, runner)
}

func (b *BinarySearchTree) TraversePreOrderResult(runner Runner, resultChan chan interface{}) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.preOrderResult(b.root, runner, resultChan)
	close(resultChan)
}

func (b *BinarySearchTree) TraversePostOrder(runner Runner) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.postOrder(b.root, runner)
}

func (b *BinarySearchTree) TraversePostOrderResult(runner Runner, resultChan chan interface{}) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.postOrderResult(b.root, runner, resultChan)
	close(resultChan)
}

func (b *BinarySearchTree) TraverseInOrder(runner Runner) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.inOrder(b.root, runner)
}

func (b *BinarySearchTree) TraverseInOrderResult(runner Runner, resultChan chan interface{}) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.inOrderResult(b.root, runner, resultChan)
	close(resultChan)
}

func (b *BinarySearchTree) TraverseLevelOrder(runner Runner) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.levelOrder(b.root, runner)
}

func (b *BinarySearchTree) TraverseLevelOrderResult(runner Runner, resultChan chan interface{}) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.levelOrderResult(b.root, runner, resultChan)
	close(resultChan)
}

func (b *BinarySearchTree) Min() interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.root == nil {
		return nil
	}

	return b.min(b.root)
}

func (b *BinarySearchTree) Max() interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.root == nil {
		return nil
	}

	return b.max(b.root)
}

func (b *BinarySearchTree) FindMin(node *Node) interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.root == nil {
		return nil
	}

	return b.min(node)
}

func (b *BinarySearchTree) FindMax(node *Node) interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.root == nil {
		return nil
	}

	return b.max(node)
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

func (b *BinarySearchTree) search(node *Node, value interface{}) *Node {
	if node == nil {
		return nil
	}

	if node.value == value {
		return node
	}

	if b.comparer(node.value, value) {
		return b.search(node.left, value)
	} else {
		return b.search(node.right, value)
	}
}

func (b *BinarySearchTree) preOrder(node *Node, runner Runner) {
	if node == nil {
		return
	}

	runner(node)
	b.preOrder(node.left, runner)
	b.preOrder(node.right, runner)
}

func (b *BinarySearchTree) preOrderResult(node *Node, runner Runner, resultChan chan interface{}) {
	if node == nil {
		return
	}

	resultChan <- node.value
	runner(node)
	b.preOrderResult(node.left, runner, resultChan)
	b.preOrderResult(node.right, runner, resultChan)
}

func (b *BinarySearchTree) postOrder(node *Node, runner Runner) {
	if node == nil {
		return
	}

	b.postOrder(node.left, runner)
	b.postOrder(node.right, runner)
	runner(node)
}

func (b *BinarySearchTree) postOrderResult(node *Node, runner Runner, resultChan chan interface{}) {
	if node == nil {
		return
	}

	b.postOrderResult(node.left, runner, resultChan)
	b.postOrderResult(node.right, runner, resultChan)
	resultChan <- node.value
	runner(node)
}

func (b *BinarySearchTree) inOrder(node *Node, runner Runner) {
	if node == nil {
		return
	}

	b.inOrder(node.left, runner)
	runner(node)
	b.inOrder(node.right, runner)
}

func (b *BinarySearchTree) inOrderResult(node *Node, runner Runner, resultChan chan interface{}) {
	if node == nil {
		return
	}

	b.inOrderResult(node.left, runner, resultChan)
	resultChan <- node.value
	runner(node)
	b.inOrderResult(node.right, runner, resultChan)
}

func (b *BinarySearchTree) levelOrder(node *Node, runner Runner) {
	if node == nil {
		return
	}

	q := queue.NewQueue()
	q.Push(node)

	for q.Len() > 0 {
		n := (*Node)(q.Pop().(*Node))

		runner(n)

		if n.left != nil {
			q.Push(n.left)
		}

		if n.right != nil {
			q.Push(n.right)
		}
	}
}

func (b *BinarySearchTree) levelOrderResult(node *Node, runner Runner, resultChan chan interface{}) {
	if node == nil {
		return
	}

	q := queue.NewQueue()
	q.Push(node)

	for q.Len() > 0 {
		n := (*Node)(q.Pop().(*Node))

		resultChan <- n.value
		runner(n)

		if n.left != nil {
			q.Push(n.left)
		}

		if n.right != nil {
			q.Push(n.right)
		}
	}
}

func (b *BinarySearchTree) min(node *Node) interface{} {
	if node.left == nil {
		return node.value
	}

	return b.min(node.left)
}

func (b *BinarySearchTree) max(node *Node) interface{} {
	if node.right == nil {
		return node.value
	}

	return b.max(node.right)
}
