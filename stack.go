package gods

import (
	"sync"
)

type node interface{}

type Stack struct {
	count int
	data  []node
	mutex sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		count: 0,
	}
}

func (s *Stack) Len() int {
	return s.count
}

func (s *Stack) AtomicLen() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.count
}

func (s *Stack) Push(item node) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, item)
	s.count++
}

func (s *Stack) Pop() node {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	topIndex := s.Len() - 1

	if topIndex < 0 {
		return nil
	}

	popItem := s.data[topIndex]
	s.data = s.data[:topIndex]

	s.count--

	return popItem
}

func (s *Stack) Peek() node {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	topIndex := s.Len() - 1

	if topIndex < 0 {
		return nil
	}

	return s.data[topIndex]
}
