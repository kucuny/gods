package tree

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BinarySearchTreeTestSuite struct {
	suite.Suite
	bst *BinarySearchTree
}

func (suite *BinarySearchTreeTestSuite) SetupSuite() {
	suite.bst = NewBinarySearchTree(IntegerComparer)
}

func (suite *BinarySearchTreeTestSuite) TestBinarySearchTree() {
	suite.Equal(0, suite.bst.Len())

	suite.bst.Insert(10)
	suite.Equal(1, suite.bst.Len())

	suite.bst.Insert(15)
	suite.Equal(2, suite.bst.Len())

	suite.bst.Insert(5)
	suite.Equal(3, suite.bst.Len())

	suite.Equal(10, suite.bst.root.value)
	suite.Equal(15, suite.bst.root.right.value)
	suite.Equal(5, suite.bst.root.left.value)
}

func TestBinarySearchTreeTestSuite(t *testing.T) {
	suite.Run(t, new(BinarySearchTreeTestSuite))
}