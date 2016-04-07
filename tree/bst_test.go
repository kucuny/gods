package tree

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BinarySearchTreeTestSuite struct {
	suite.Suite
	bst *BinarySearchTree
}

func (suite *BinarySearchTreeTestSuite) SetupTest() {
	suite.bst = NewBinarySearchTree(IntegerComparer)
}

func (suite *BinarySearchTreeTestSuite) TestBinarySearchTreeBasic() {
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

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeTraversePreOrderGetResult() {
	suite.bst.Insert(20)
	suite.bst.Insert(6)
	suite.bst.Insert(7)
	suite.bst.Insert(35)
	suite.bst.Insert(10)
	suite.bst.Insert(25)
	suite.bst.Insert(8)

	suite.Equal(7, suite.bst.Len())

	runner := func(value *Node) {}

	resultChan := make(chan interface{})
	var result []interface{}
	go suite.bst.TraversePreOrderResult(runner, resultChan)

	for r := range resultChan {
		result = append(result, r)
	}

	suite.Equal([]interface{}{20, 6, 7, 10, 8, 35, 25}, result)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeTraverseInOrderGetResult() {
	suite.bst.Insert(10)
	suite.bst.Insert(4)
	suite.bst.Insert(20)
	suite.bst.Insert(7)
	suite.bst.Insert(15)
	suite.bst.Insert(5)

	suite.Equal(6, suite.bst.Len())

	runner := func(value *Node) {}
	resultChan := make(chan interface{})
	var result []interface{}
	go suite.bst.TraverseInOrderResult(runner, resultChan)

	for r := range resultChan {
		result = append(result, r)
	}

	suite.Equal([]interface{}{4, 5, 7, 10, 15, 20}, result)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeTraversePostOrderGetResult() {
	suite.bst.Insert(12)
	suite.bst.Insert(1)
	suite.bst.Insert(7)
	suite.bst.Insert(20)
	suite.bst.Insert(4)

	suite.Equal(5, suite.bst.Len())

	runner := func(value *Node) {}
	resultChan := make(chan interface{})
	var result []interface{}
	go suite.bst.TraversePostOrderResult(runner, resultChan)

	for r := range resultChan {
		result = append(result, r)
	}

	suite.Equal([]interface{}{4, 7, 1, 20, 12}, result)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeTraverseLevelOrderGetResult() {
	suite.bst.Insert(10)
	suite.bst.Insert(15)
	suite.bst.Insert(5)
	suite.bst.Insert(3)

	suite.Equal(4, suite.bst.Len())

	runner := func(value *Node) {}
	resultChan := make(chan interface{})
	var result []interface{}
	go suite.bst.TraverseLevelOrderResult(runner, resultChan)

	for r := range resultChan {
		result = append(result, r)
	}

	suite.Equal([]interface{}{10, 5, 15, 3}, result)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeSearchSuccess() {
	suite.bst.Insert(10)
	suite.bst.Insert(15)
	suite.bst.Insert(5)
	suite.bst.Insert(3)

	suite.Equal(4, suite.bst.Len())

	searched := suite.bst.Search(10)
	suite.Equal(10, searched.value)

	searched = suite.bst.Search(3)
	suite.Equal(3, searched.value)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeSearchFailure() {
	suite.bst.Insert(10)
	suite.bst.Insert(15)
	suite.bst.Insert(5)
	suite.bst.Insert(3)

	suite.Equal(4, suite.bst.Len())

	searched := suite.bst.Search(100)
	suite.Nil(searched)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeGetMinValueIsExist() {
	suite.bst.Insert(10)
	suite.bst.Insert(15)
	suite.bst.Insert(5)
	suite.bst.Insert(3)

	suite.Equal(4, suite.bst.Len())

	minValue := suite.bst.Min()
	suite.Equal(3, minValue)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeGetMinValueIsNotExist() {
	suite.Equal(0, suite.bst.Len())

	minValue := suite.bst.Min()
	suite.Nil(minValue)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeGetMaxValueIsExist() {
	suite.bst.Insert(10)
	suite.bst.Insert(15)
	suite.bst.Insert(5)
	suite.bst.Insert(3)

	suite.Equal(4, suite.bst.Len())

	minValue := suite.bst.Max()
	suite.Equal(15, minValue)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeGetMaxValueIsNotExist() {
	suite.Equal(0, suite.bst.Len())

	minValue := suite.bst.Max()
	suite.Nil(minValue)
}

func TestBinarySearchTreeTestSuite(t *testing.T) {
	suite.Run(t, new(BinarySearchTreeTestSuite))
}
