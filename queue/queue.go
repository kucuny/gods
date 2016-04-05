package queue

import (
	"sync"
)

type node interface{}

type Queue struct {
	count int
	data  []node
	mutex *sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		count: 0,
		mutex: new(sync.Mutex),
	}
}

func (q *Queue) Len() int {
	return q.count
}

func (q *Queue) AtomicLen() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return q.count
}

func (q *Queue) Push(item node) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.data = append(q.data, item)
	q.count++
}

func (q *Queue) Pop() node {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.Len() == 0 {
		return nil
	}

	popItem := q.data[0]
	q.data = q.data[1:]

	q.count--

	return popItem
}

func (q *Queue) Peek() node {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if q.Len() < 1 {
		return nil
	}

	return q.data[0]
}
