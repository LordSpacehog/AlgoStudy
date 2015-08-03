// Package algostudy provides algostudy
package algostudy

import (
	"errors"
)

type Frame struct {
	Data string
	Next *Frame
}

type Stack struct {
	Head *Frame
	Size int
}

func NewStack() Stack {
	return Stack{
		Head: nil,
		Size: 0,
	}
}

func (s *Stack) Push(data string) {
	newframe := Frame{
		Data: data,
		Next: s.Head,
	}

	s.Head = &newframe
	s.Size += 1
}

func (s *Stack) Pop() (string, error) {
	if s.Head == nil {
		return "", errors.New("Error! Stack Empty!")
	}
	ret := s.Head.Data
	s.Head = s.Head.Next
	s.Size -= 1
	return ret, nil
}
