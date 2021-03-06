package structures

import (
	"errors"
	"math/big"

	"github.com/azraeljack/algorithm-practice/common"
)

var (
	errStackEmpty           = errors.New("stack is empty")
	errStackFull            = errors.New("stack is full")
	errStackLimitOutOfBound = errors.New("stack limit must be greater than 0")
)

// linked list implementation of Stack
type Stack struct {
	root        *common.ForwardDirectionalNode
	size, limit *big.Int
}

func NewStack(limit *big.Int) (*Stack, error) {
	if limit.Cmp(common.Big0) <= 0 {
		return nil, errStackLimitOutOfBound
	}
	return &Stack{
		root:  nil,
		size:  new(big.Int).Set(common.Big0),
		limit: new(big.Int).Set(limit),
	}, nil
}

func (s *Stack) Empty() bool {
	if s.root == nil {
		return true
	}
	return false
}

func (s *Stack) Peek() (interface{}, error) {
	if s.root == nil {
		return nil, errStackEmpty
	}
	return s.root.Data, nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.root == nil {
		return nil, errStackEmpty
	}
	pointer := s.root
	s.root = s.root.Next
	s.size.Sub(s.size, common.Big1)
	return pointer.Data, nil
}

func (s *Stack) Size() *big.Int {
	return s.size
}

func (s *Stack) Push(data interface{}) error {
	if s.size.Cmp(s.limit) >= 0 {
		return errStackFull
	}
	s.root = &common.ForwardDirectionalNode{
		Next: s.root,
		Data: data,
	}
	s.size.Add(s.size, common.Big1)
	return nil
}
