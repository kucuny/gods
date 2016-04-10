package tree

import (
	"github.com/kucuny/gods/queue"
	"sync"
)

type ComparerType int

const (
	ComparerLarger = iota
	ComparerSmaller
	ComparerEqual
)

type Comparer func(source, target interface{}) ComparerType
type Runner func(value *Node)

func IntegerComparer(source, target interface{}) ComparerType {
	if source.(int) < target.(int) {
		return ComparerLarger
	}

	if source.(int) > target.(int) {
		return ComparerSmaller
	}

	return ComparerEqual
}

func Integer64Comparer(source, target interface{}) ComparerType {
	if source.(int64) < target.(int64) {
		return ComparerLarger
	}

	if source.(int64) > target.(int64) {
		return ComparerSmaller
	}

	return ComparerEqual
}

func Float32Comparer(source, target interface{}) ComparerType {
	if source.(float32) < target.(float32) {
		return ComparerLarger
	}

	if source.(float32) > target.(float32) {
		return ComparerSmaller
	}

	return ComparerEqual
}

func Float64Comparer(source, target interface{}) ComparerType {
	if source.(float64) < target.(float64) {
		return ComparerLarger
	}

	if source.(float64) > target.(float64) {
		return ComparerSmaller
	}

	return ComparerEqual
}

type Node struct {
	left, right *Node
	Value       interface{}
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value, left: nil, right: nil}
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
		b.root = NewNode(value)
		isSuccess = true
	} else {
		node := b.insert(b.root, value)

		if node != nil {
			isSuccess = true
		} else {
			isSuccess = false
		}
	}

	if isSuccess {
		b.count++
	}

	return isSuccess
}

func (b *BinarySearchTree) Remove(value interface{}) *Node {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.remove(b.root, value)
	b.count--
	return nil
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

func (b *BinarySearchTree) Min() *Node {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.root == nil {
		return nil
	}

	return b.min(b.root)
}

func (b *BinarySearchTree) Max() *Node {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.root == nil {
		return nil
	}

	return b.max(b.root)
}

func (b *BinarySearchTree) FindMin(node *Node) *Node {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.root == nil {
		return nil
	}

	return b.min(node)
}

func (b *BinarySearchTree) FindMax(node *Node) *Node {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.root == nil {
		return nil
	}

	return b.max(node)
}

func (b *BinarySearchTree) insert(node *Node, value interface{}) *Node {
	if node == nil {
		return NewNode(value)
	}

	if b.comparer(node.Value, value) == ComparerSmaller {
		node.left = b.insert(node.left, value)
		return node
	} else if b.comparer(node.Value, value) == ComparerLarger {
		node.right = b.insert(node.right, value)
		return node
	} else {
		return nil
	}
}

func (b *BinarySearchTree) remove(node *Node, removeValue interface{}) {
	if node == nil {
		return
	}

	if b.comparer(node.Value, removeValue) == ComparerSmaller {
		b.remove(node.left, removeValue)
	} else if b.comparer(node.Value, removeValue) == ComparerLarger {
		b.remove(node.right, removeValue)
	} else if b.comparer(node.Value, removeValue) == ComparerEqual {
		if node.left == nil && node.right == nil {
			node = nil
		} else if node.left != nil {
			*node = *node.left
		} else if node.right != nil {
			*node = *node.right
		} else {
			node = nil
		}
	}

	return
}

func (b *BinarySearchTree) search(node *Node, value interface{}) *Node {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if b.comparer(node.Value, value) == ComparerSmaller {
		return b.search(node.left, value)
	} else if b.comparer(node.Value, value) == ComparerLarger {
		return b.search(node.right, value)
	} else {
		return node
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

	resultChan <- node.Value
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
	resultChan <- node.Value
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
	resultChan <- node.Value
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

		resultChan <- n.Value
		runner(n)

		if n.left != nil {
			q.Push(n.left)
		}

		if n.right != nil {
			q.Push(n.right)
		}
	}
}

func (b *BinarySearchTree) min(node *Node) *Node {
	if node.left == nil {
		return node
	}

	return b.min(node.left)
}

func (b *BinarySearchTree) max(node *Node) *Node {
	if node.right == nil {
		return node
	}

	return b.max(node.right)
}
