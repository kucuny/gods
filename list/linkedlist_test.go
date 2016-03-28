package list

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type LinkedListTestSuite struct {
	suite.Suite
	link *LinkedList
}

func (suite *LinkedListTestSuite) SetupSuite() {
	suite.link = NewLinkedList()
}

func (suite *LinkedListTestSuite) TestLinkedList() {
	suite.Nil(suite.link.Front())
	suite.Nil(suite.link.Back())

	node1 := &Node{Value: 30}
	suite.link.PushFront(node1)
	node2 := &Node{Value: 40}
	suite.link.PushBack(node2)

	suite.Equal(2, suite.link.Len())
	suite.Equal(node1, suite.link.Front())
	suite.Equal(node2, suite.link.Back())

	suite.Equal(node2, node1.Next())
	suite.Equal(node1, node2.Prev())

	insertNode1 := &Node{Value: 100}
	insertNode2 := &Node{Value: 200}

	insertNode1 = suite.link.InsertBefore(insertNode1, node2)
	insertNode2 = suite.link.InsertAfter(insertNode2, node1)

	suite.Equal(4, suite.link.Len())
	suite.Equal(node2, insertNode1.Next())
	suite.Equal(node1, insertNode2.Prev())

	removeNode := insertNode1
	resRemove := suite.link.Remove(removeNode)

	suite.Equal(3, suite.link.Len())
	suite.Nil(resRemove.prev)
	suite.Nil(resRemove.next)
	suite.Equal(100, resRemove.Value)

	head := suite.link.head
	tail := suite.link.tail

	var nextNode *Node

	nextNode = head.Next()
	suite.Equal(node1, nextNode)

	nextNode = nextNode.Next()
	suite.Equal(insertNode2, nextNode)

	nextNode = nextNode.Next()
	suite.Equal(node2, nextNode)

	nextNode = nextNode.Next()
	suite.Equal(tail, nextNode)
}

func TestLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(LinkedListTestSuite))
}
