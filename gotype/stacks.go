package gotype

import (
	"reflect"

	structform "github.com/urso/go-structform"
)

type keyStack struct {
	current string
	stack   []string
	stack0  [32]string
}

type idxStack struct {
	current int
	stack   []int
	stack0  [32]int
}

type reflectValueStack struct {
	current reflect.Value
	stack   []reflect.Value
	stack0  [32]reflect.Value
}

type unfoldStateStack struct {
	current unfoldState
	stack   []unfoldState
	stack0  [32]unfoldState
}

type unfoldStateDetailStack struct {
	current unfoldStateDetail
	stack   []unfoldStateDetail
	stack0  [32]unfoldStateDetail
}

type structTypeStack struct {
	current structform.BaseType
	stack   []structform.BaseType
	stack0  [32]structform.BaseType
}

func (s *keyStack) init() {
	s.current = ""
	s.stack = s.stack[:0]
}

func (s *keyStack) push(name string) {
	s.stack = append(s.stack, s.current)
	s.current = name
}

func (s *keyStack) pop() string {
	name := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return name
}

func (s *idxStack) init() {
	s.current = -1
	s.stack = s.stack[:0]
}

func (s *idxStack) push() {
	s.stack = append(s.stack, s.current)
	s.current = 0
}

func (s *idxStack) pop() int {
	name := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return name
}

func (s *reflectValueStack) init(v reflect.Value) {
	s.current = v
	s.stack = s.stack0[:1]
	s.stack[0] = reflect.Value{}
}

func (s *reflectValueStack) push(v reflect.Value) {
	s.stack = append(s.stack, s.current)
	s.current = v
}

func (s *reflectValueStack) pop() reflect.Value {
	v := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return v
}

func (s *unfoldStateStack) init(v unfoldState) {
	s.current = v
	s.stack = s.stack0[:0]
}

func (s *unfoldStateStack) push(v unfoldState) {
	s.stack = append(s.stack, s.current)
	s.current = v
}

func (s *unfoldStateStack) pop() unfoldState {
	v := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return v
}

func (s *unfoldStateDetailStack) init(v unfoldStateDetail) {
	s.current = v
	s.stack = s.stack0[:0]
}

func (s *unfoldStateDetailStack) push(v unfoldStateDetail) {
	s.stack = append(s.stack, s.current)
	s.current = v
}

func (s *unfoldStateDetailStack) pop() unfoldStateDetail {
	v := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return v
}
