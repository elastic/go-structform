// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package gotype

import (
	"errors"
	"fmt"
	"reflect"

	stunsafe "github.com/elastic/go-structform/internal/unsafe"
)

func makeUserUnfolder(fn reflect.Value) (target reflect.Type, unfolder reflUnfolder, err error) {
	t := fn.Type()

	if fn.Kind() != reflect.Func {
		return nil, nil, errors.New("function type required")
	}

	switch {
	case t.NumIn() == 2 && t.NumOut() == 1:
		unfolder, err = makeUserPrimitiveUnfolder(fn)
	case t.NumIn() == 1 && t.NumOut() == 2:
		unfolder, err = makeUserProcessingUnfolder(fn)
	default:
		return nil, nil, fmt.Errorf("invalid number of arguments in unfolder: %v", fn)
	}

	return t.In(0), unfolder, err
}

func makeUserPrimitiveUnfolder(fn reflect.Value) (reflUnfolder, error) {
	t := fn.Type()

	if fn.Kind() != reflect.Func {
		return nil, errors.New("function type required")
	}

	if t.NumIn() != 2 {
		return nil, fmt.Errorf("function '%v' must accept 2 arguments", fn)
	}
	if t.NumOut() != 1 || (t.NumOut() > 0 && t.Out(0) != tError) {
		return nil, fmt.Errorf("function '%v' does not return errors", fn)
	}

	ta0 := t.In(0)
	if ta0.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("first argument in function '%v' must be a pointer", t.Name())
	}

	constr := lookupUserPrimitiveConstructor(t.In(1))
	if constr == nil {
		return nil, fmt.Errorf("%v is no supported primitive type", t.In(1))
	}

	unfolder := constr(fn)
	return liftGoUnfolder(unfolder), nil
}

func makeUserProcessingUnfolder(fn reflect.Value) (reflUnfolder, error) {
	if err := checkProcessingUnfolder(fn); err != nil {
		return nil, err
	}

	return liftGoUnfolder(&unfolderUserProcessingInit{
		fnInit: *((*userProcessingInitFn)(stunsafe.UnsafeFnPtr(fn))),
	}), nil
}

func checkProcessingUnfolder(fn reflect.Value) error {
	if fn.Kind() != reflect.Func {
		return fmt.Errorf("processing unfolder '%v' is no function", fn)
	}

	t := fn.Type()

	// check input
	if t.NumIn() != 1 {
		return fmt.Errorf("processing unfolder '%v' must accept one target argument", fn)
	}
	in := t.In(0)
	if in.Kind() != reflect.Ptr {
		return fmt.Errorf("processing unfolder '%v' target argument must be a pointer", fn)
	}

	// check returns
	if t.NumOut() != 2 {
		return fmt.Errorf("processing unfolder '%v' must return 2 values", fn)
	}
	if t.Out(0) != tInterface {
		return fmt.Errorf("processing unfolder '%v' must return interface{} as first value", fn)
	}
	proc := t.Out(1)
	if proc.Kind() != reflect.Func {
		return fmt.Errorf("processing unfolder '%v' second return must be a function", fn)
	}

	// check processing function input
	if proc.NumIn() != 2 {
		return fmt.Errorf("processing function of '%v' must accept 2 arguments", fn)
	}
	if proc.In(0) != in {
		return fmt.Errorf("processing function of '%v' must accept the target type '%v'", fn, in)
	}
	if proc.In(1) != tInterface {
		return fmt.Errorf("processing function of '%v' must accept interface{} as second argument", fn)
	}

	// check processing function output
	if proc.NumOut() != 1 {
		return fmt.Errorf("processing function of '%v' must return exactly one value", fn)
	}
	if proc.Out(0) != tError {
		return fmt.Errorf("processing function of '%v' must return an error value", fn)
	}

	return nil
}
