package cqueue

import (
	"fmt"
)

type CQueue struct {
	in  Stack
	out Stack
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)

}

func (this *CQueue) DeleteHead() int {
	val, err := this.out.Pop()
	if err == nil {
		return val
	}
	for {
		val, err = this.in.Pop()
		if err != nil {
			break
		}
		this.out.Push(val)
	}
	val, err = this.out.Pop()
	if err != nil {
		return -1
	}
	return val
}

type Stack struct {
	list []int
}

func (s *Stack) Push(v int) {
	s.list = append(s.list, v)
}

func (s *Stack) Pop() (int, error) {
	if len(s.list) != 0 {
		v := s.list[len(s.list)-1]
		s.list = s.list[:len(s.list)-1]
		return v, nil
	}
	return 0, fmt.Errorf("stack is nil")
}
