// This file has been generated from 'stacks_tmpl.yml', do not edit
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

type unfoldTypeStack struct {
	current unfoldType
	stack   []unfoldType
	stack0  [32]unfoldType
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

type unfolderStack struct {
	current structform.Visitor
	stack   []structform.Visitor
	stack0  [32]structform.Visitor
}

func (s *keyStack) init() {
	s.current = ""
	s.stack = s.stack0[:0]
}

func (s *keyStack) push(v string) {
	s.stack = append(s.stack, s.current)
	s.current = v
}

func (s *keyStack) pop() string {
	old := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return old
}

func (s *idxStack) init() {
	s.current = -1
	s.stack = s.stack0[:0]
}

func (s *idxStack) push(v int) {
	s.stack = append(s.stack, s.current)
	s.current = v
}

func (s *idxStack) pop() int {
	old := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return old
}

func (s *reflectValueStack) init(v reflect.Value) {
	s.current = v
	s.stack = s.stack0[:0]
}

func (s *reflectValueStack) push(v reflect.Value) {
	s.stack = append(s.stack, s.current)
	s.current = v
}

func (s *reflectValueStack) pop() reflect.Value {
	old := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return old
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
	old := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return old
}

func (s *unfoldTypeStack) init(v unfoldType) {
	s.current = v
	s.stack = s.stack0[:0]
}

func (s *unfoldTypeStack) push(v unfoldType) {
	s.stack = append(s.stack, s.current)
	s.current = v
}

func (s *unfoldTypeStack) pop() unfoldType {
	old := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return old
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
	old := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return old
}

func (s *structTypeStack) init(v structform.BaseType) {
	s.current = v
	s.stack = s.stack0[:0]
}

func (s *structTypeStack) push(v structform.BaseType) {
	s.stack = append(s.stack, s.current)
	s.current = v
}

func (s *structTypeStack) pop() structform.BaseType {
	old := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return old
}

func (s *unfolderStack) init(v structform.Visitor) {
	s.current = v
	s.stack = s.stack0[:0]
}

func (s *unfolderStack) push(v structform.Visitor) {
	s.stack = append(s.stack, s.current)
	s.current = v
}

func (s *unfolderStack) pop() structform.Visitor {
	old := s.current
	last := len(s.stack) - 1
	s.current = s.stack[last]
	s.stack = s.stack[:last]
	return old
}
