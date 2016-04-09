package tree

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BinarySearchTreeTestSuite struct {
	suite.Suite
	bstInt     *BinarySearchTree
	bstInt64   *BinarySearchTree
	bstFloat32 *BinarySearchTree
	bstFloat64 *BinarySearchTree
}

func (suite *BinarySearchTreeTestSuite) SetupTest() {
	suite.bstInt = NewBinarySearchTree(IntegerComparer)
	suite.bstInt64 = NewBinarySearchTree(Integer64Comparer)
	suite.bstFloat32 = NewBinarySearchTree(Float32Comparer)
	suite.bstFloat64 = NewBinarySearchTree(Float64Comparer)
}

func (suite *BinarySearchTreeTestSuite) TestBinarySearchTreeInt() {
	suite.Equal(0, suite.bstInt.Len())

	suite.bstInt.Insert(10)
	suite.Equal(1, suite.bstInt.Len())

	suite.bstInt.Insert(15)
	suite.Equal(2, suite.bstInt.Len())

	suite.bstInt.Insert(5)
	suite.Equal(3, suite.bstInt.Len())

	suite.Equal(10, suite.bstInt.root.Value)
	suite.Equal(15, suite.bstInt.root.right.Value)
	suite.Equal(5, suite.bstInt.root.left.Value)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntTraversePreOrderGetResult() {
	suite.bstInt.Insert(20)
	suite.bstInt.Insert(6)
	suite.bstInt.Insert(7)
	suite.bstInt.Insert(35)
	suite.bstInt.Insert(10)
	suite.bstInt.Insert(25)
	suite.bstInt.Insert(8)

	suite.Equal(7, suite.bstInt.Len())

	runner := func(value *Node) {}

	resultChan := make(chan interface{})
	var result []interface{}
	go suite.bstInt.TraversePreOrderResult(runner, resultChan)

	for r := range resultChan {
		result = append(result, r)
	}

	suite.Equal([]interface{}{20, 6, 7, 10, 8, 35, 25}, result)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntTraverseInOrderGetResult() {
	suite.bstInt.Insert(10)
	suite.bstInt.Insert(4)
	suite.bstInt.Insert(20)
	suite.bstInt.Insert(7)
	suite.bstInt.Insert(15)
	suite.bstInt.Insert(5)

	suite.Equal(6, suite.bstInt.Len())

	runner := func(value *Node) {}
	resultChan := make(chan interface{})
	var result []interface{}
	go suite.bstInt.TraverseInOrderResult(runner, resultChan)

	for r := range resultChan {
		result = append(result, r)
	}

	suite.Equal([]interface{}{4, 5, 7, 10, 15, 20}, result)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntTraversePostOrderGetResult() {
	suite.bstInt.Insert(12)
	suite.bstInt.Insert(1)
	suite.bstInt.Insert(7)
	suite.bstInt.Insert(20)
	suite.bstInt.Insert(4)

	suite.Equal(5, suite.bstInt.Len())

	runner := func(value *Node) {}
	resultChan := make(chan interface{})
	var result []interface{}
	go suite.bstInt.TraversePostOrderResult(runner, resultChan)

	for r := range resultChan {
		result = append(result, r)
	}

	suite.Equal([]interface{}{4, 7, 1, 20, 12}, result)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntTraverseLevelOrderGetResult() {
	suite.bstInt.Insert(10)
	suite.bstInt.Insert(15)
	suite.bstInt.Insert(5)
	suite.bstInt.Insert(3)

	suite.Equal(4, suite.bstInt.Len())

	runner := func(value *Node) {}
	resultChan := make(chan interface{})
	var result []interface{}
	go suite.bstInt.TraverseLevelOrderResult(runner, resultChan)

	for r := range resultChan {
		result = append(result, r)
	}

	suite.Equal([]interface{}{10, 5, 15, 3}, result)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntSearchSuccess() {
	suite.bstInt.Insert(10)
	suite.bstInt.Insert(15)
	suite.bstInt.Insert(5)
	suite.bstInt.Insert(3)

	suite.Equal(4, suite.bstInt.Len())

	searched := suite.bstInt.Search(10)
	suite.Equal(10, searched.Value)

	searched = suite.bstInt.Search(3)
	suite.Equal(3, searched.Value)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntSearchFailure() {
	suite.bstInt.Insert(10)
	suite.bstInt.Insert(15)
	suite.bstInt.Insert(5)
	suite.bstInt.Insert(3)

	suite.Equal(4, suite.bstInt.Len())

	searched := suite.bstInt.Search(100)
	suite.Nil(searched)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntGetMinValueIsExist() {
	suite.bstInt.Insert(10)
	suite.bstInt.Insert(15)
	suite.bstInt.Insert(5)
	suite.bstInt.Insert(3)

	suite.Equal(4, suite.bstInt.Len())

	minValue := suite.bstInt.Min()
	suite.Equal(3, minValue.Value)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntGetMinValueIsNotExist() {
	suite.Equal(0, suite.bstInt.Len())

	minValue := suite.bstInt.Min()
	suite.Nil(minValue)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntGetMaxValueIsExist() {
	suite.bstInt.Insert(10)
	suite.bstInt.Insert(15)
	suite.bstInt.Insert(5)
	suite.bstInt.Insert(3)

	suite.Equal(4, suite.bstInt.Len())

	minValue := suite.bstInt.Max()
	suite.Equal(15, minValue.Value)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntGetMaxValueIsNotExist() {
	suite.Equal(0, suite.bstInt.Len())

	minValue := suite.bstInt.Max()
	suite.Nil(minValue)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntGetFindMinValueIsExist() {
	suite.bstInt.Insert(10)
	suite.bstInt.Insert(15)
	suite.bstInt.Insert(7)
	suite.bstInt.Insert(40)
	suite.bstInt.Insert(5)
	suite.bstInt.Insert(3)
	suite.bstInt.Insert(6)
	suite.bstInt.Insert(12)
	suite.bstInt.Insert(11)
	suite.bstInt.Insert(13)
	suite.bstInt.Insert(22)
	suite.bstInt.Insert(45)
	suite.bstInt.Insert(20)
	suite.bstInt.Insert(24)

	suite.Equal(14, suite.bstInt.Len())

	searchedNode := suite.bstInt.Search(40)
	minValue := suite.bstInt.FindMin(searchedNode)
	suite.Equal(20, minValue.Value)
}

func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntGetFindMaxValueIsExist() {
	suite.bstInt.Insert(10)
	suite.bstInt.Insert(15)
	suite.bstInt.Insert(7)
	suite.bstInt.Insert(40)
	suite.bstInt.Insert(5)
	suite.bstInt.Insert(3)
	suite.bstInt.Insert(6)
	suite.bstInt.Insert(12)
	suite.bstInt.Insert(11)
	suite.bstInt.Insert(13)
	suite.bstInt.Insert(22)
	suite.bstInt.Insert(45)
	suite.bstInt.Insert(20)
	suite.bstInt.Insert(24)

	suite.Equal(14, suite.bstInt.Len())

	searchedNode := suite.bstInt.Search(15)

	minValue := suite.bstInt.FindMax(searchedNode)
	suite.Equal(45, minValue.Value)
}

//func (suite *BinarySearchTreeTestSuite) TestBianrySearchTreeIntRemoveNodeIsSuccess() {
//	suite.bstInt.Insert(10)
//	suite.bstInt.Insert(15)
//	suite.bstInt.Insert(7)
//	suite.bstInt.Insert(40)
//	suite.bstInt.Insert(5)
//	suite.bstInt.Insert(3)
//	suite.bstInt.Insert(6)
//	suite.bstInt.Insert(12)
//	suite.bstInt.Insert(11)
//	suite.bstInt.Insert(13)
//	suite.bstInt.Insert(22)
//	suite.bstInt.Insert(45)
//	suite.bstInt.Insert(20)
//	suite.bstInt.Insert(24)
//
//	suite.Equal(14, suite.bstInt.Len())
//
//	suite.bstInt.Remove(3)
//	suite.Equal(13, suite.bstInt.Len())
//	suite.Equal(5, suite.bstInt.Min().Value)
//}

func TestBinarySearchTreeTestSuite(t *testing.T) {
	suite.Run(t, new(BinarySearchTreeTestSuite))
}
