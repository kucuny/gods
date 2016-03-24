package gods

import (
	"fmt"
)

type Node interface{}

type Queue struct {
	data []Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (s *Queue) Len() int {
	return len(s.data)
}

func (s *Queue) Push(item Node) {
	s.data = append(s.data, item)
}

func (s *Queue) Pop() Node {
	if s.Len() == 0 {
		return nil
	}

	popItem := s.data[0]
	s.data = s.data[1:]

	return popItem
}
