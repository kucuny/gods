package stack

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type StackTestSuite struct {
	suite.Suite
	s *Stack
}

func (suite *StackTestSuite) SetupSuite() {
	suite.s = NewStack()
}

func (suite *StackTestSuite) TestStack() {
	suite.Equal(0, suite.s.Len())

	suite.s.Push("abcd")
	suite.s.Push("1235")
	suite.s.Push(1234)
	suite.s.Push(345.123)

	suite.Equal(4, suite.s.Len())
	suite.Equal(345.123, suite.s.Pop())
	suite.Equal(3, suite.s.Len())
	suite.Equal(1234, suite.s.Pop())
	suite.Equal(2, suite.s.Len())
	suite.Equal("1235", suite.s.Pop())
	suite.Equal(1, suite.s.Len())
	suite.Equal("abcd", suite.s.Peek())
	suite.Equal(1, suite.s.Len())
}

func TestStackTestSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}
