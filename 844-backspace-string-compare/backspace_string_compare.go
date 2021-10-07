package _44_backspace_string_compare

import "container/list"

type stack struct {
	list *list.List
}

func newStakc() *stack {
	l := list.New()
	return &stack{
		l,
	}
}

func (s *stack) push(v interface{}) {
	s.list.PushBack(v)
}

func (s *stack) pop() interface{} {
	v := s.list.Back()
	if v != nil {
		s.list.Remove(v)
		return v.Value
	}
	return nil
}

func (s *stack) isEmpty() bool {
	return s.list.Len() == 0
}

func (s *stack) length() int {
	return s.list.Len()
}

func (s *stack) peek() interface{} {
	v := s.list.Back()
	if v != nil {
		return v.Value
	}
	return nil
}

func backspaceCompare(S string, T string) bool {
	s1, s2 := compare(S), compare(T)
	if s1.length() != s2.length() {
		return false
	}
	for !s1.isEmpty() && !s2.isEmpty() {
		v1 := s1.pop().(byte)
		v2 := s2.pop().(byte)
		if v1 != v2 {
			return false
		}
	}
	if s1.isEmpty() && s2.isEmpty() {
		return true
	}
	return false
}

func compare(str string) *stack {
	s := newStakc()
	for i := 0; i < len(str); i++ {
		if str[i] != '#' {
			s.push(str[i])
		} else {
			s.pop()
		}
	}
	return s
}
