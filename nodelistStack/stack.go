package nodelistStack

import "sync"

type node struct {
	value interface{}
	prev  *node
}

type stack struct {
	lock   *sync.RWMutex
	top    *node
	length int
}

func NewStack() *stack {
	return &stack{
		length: 0,
		top:    nil,
		lock:   &sync.RWMutex{},
	}
}

func (s *stack) Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.length == 0 {
		s.top = &node{
			value: value,
		}
	} else {
		node := &node{
			value: value,
			prev:  s.top,
		}
		s.top = node
	}
	s.length++
}

func (s *stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}

	return s.top.value
}

func (s *stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	value := s.Peek()
	s.top = s.top.prev
	s.length--
	return value
}

func (s *stack) Len() int {
	return s.length
}
