package gotype

import (
	"reflect"
	"testing"
)

func BenchmarkResolveVisitor(b *testing.B) {
	benchmarks := []struct {
		value interface{}
	}{
		{
			int8(23),
		},
		{
			float32(3.14),
		},
		{
			[]uint{1, 2, 3},
		},
		{
			map[string]int{
				"a": 1,
			},
		},
		{
			map[string]map[string]string{
				"a": nil,
			},
		},
		{
			map[string][]int{
				"a": {1, 2, 3},
			},
		},
		{
			map[string]map[string]interface{}{
				"a": nil,
			},
		},
		{
			struct {
				A int
				B uint
				C map[string]bool
				D struct {
					X bool
				}
				int
			}{},
		},
		{
			struct {
				A *int
				B *uint
				C *map[string]bool
				D *struct {
					X bool
				}
			}{},
		},
	}

	for i, run := range benchmarks {
		// reset registry between runs
		visitorRegistry = newTypeVisitorRegistry()

		name := reflect.TypeOf(run.value).String()
		b.Logf("run benchmark (%v): %v", i, name)
		b.Run(name, makeGetVisitorBench(run.value))
	}

}

func makeGetVisitorBench(v interface{}) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if f := getTypeVisitor(v); f != nil {
				continue
			}

			f, err := getVisitorReflect(reflect.ValueOf(v).Type())
			if err != nil {
				b.Error(err)
			}
			if f == nil {
				b.Errorf("no visitor")
				return
			}
		}
	}
}
