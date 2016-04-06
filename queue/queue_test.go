package queue

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type QueueTestSuite struct {
	suite.Suite
	q *Queue
}

type testNode struct {
	strValue string
	intValue int
}

func (suite *QueueTestSuite) SetupSuite() {
	suite.q = NewQueue()
}

func (suite *QueueTestSuite) TestQueue() {
	suite.Equal(0, suite.q.Len())

	suite.q.Push("abcd")
	suite.q.Push("1235")
	suite.q.Push(1234)
	suite.q.Push(345.123)
	suite.q.Push(testNode{strValue: "test1", intValue: 30})

	suite.Equal(5, suite.q.Len())
	suite.Equal("abcd", suite.q.Pop())
	suite.Equal(4, suite.q.Len())
	suite.Equal("1235", suite.q.Pop())
	suite.Equal(3, suite.q.Len())
	suite.Equal(1234, suite.q.Pop())
	suite.Equal(2, suite.q.Len())
	suite.Equal(345.123, suite.q.Peek())
	suite.Equal(2, suite.q.Len())
	suite.Equal(345.123, suite.q.Pop())
	suite.Equal(1, suite.q.Len())
	suite.Equal(testNode{strValue: "test1", intValue: 30}, suite.q.Pop().(testNode))
}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}
