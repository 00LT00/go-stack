package go-stack

import "sync"

type stack struct {
	lock   *sync.RWMutex
	arr    []interface{}
	length int
}

func NewStack() *stack {
	return &stack{lock: &sync.RWMutex{}, arr: nil, length: 0}
}

func (s *stack) Push(value ...interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.arr = append(s.arr, value...)
	s.length += len(value)
}

func (s *stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}

	return s.arr[s.length-1]
}

func (s *stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	value := s.Peek()
	s.length--
	s.arr = s.arr[:s.length]

	return value
}

func (s *stack) Len() int {
	return s.length
}
