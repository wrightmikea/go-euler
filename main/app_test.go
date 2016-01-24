// euler1 tests
package main

import (
	"reflect"
	"testing"
)

var result int

func TestMain(t *testing.T) {
	main()
}

func TestEuler1a(t *testing.T) {
	runTestFuncWith1000(t, Euler1a)
}
func TestEuler1b(t *testing.T) {
	runTestFuncWith1000(t, Euler1b)
}
func TestEuler1c(t *testing.T) {
	runTestFuncWith1000(t, Euler1c)
}
func TestEuler1d(t *testing.T) {
	runTestFuncWith1000(t, Euler1d)
}
func TestEuler1e(t *testing.T) {
	runTestFuncWith1000(t, Euler1e)
}
func TestEuler1f(t *testing.T) {
	runTestFuncWith1000(t, Euler1f)
}
func TestEuler1g(t *testing.T) {
	runTestFuncWith1000(t, Euler1g)
}
func TestEuler1h(t *testing.T) {
	runTestFuncWith1000(t, Euler1h)
}

func runTestFuncWith1000(t *testing.T, f func(int) int) {
	actual := f(1000)
	expected := 233168
	if !compareEq(expected, actual) {
		t.Fatalf("Expected %v (%v), got %v (%v)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
	}
}
func compareEq(expected interface{}, actual interface{}) bool {
	return expected == actual
}
func BenchmarkEuler1a1000x(b *testing.B) {
	for n := 0; n < 1000; n++ {
		result = Euler1a(100000)
	}
}
func BenchmarkEuler1b1000x(b *testing.B) {
	for n := 0; n < 1000; n++ {
		result = Euler1b(100000)
	}
}
func BenchmarkEuler1c1000x(b *testing.B) {
	for n := 0; n < 1000; n++ {
		result = Euler1c(100000)
	}
}
func BenchmarkEuler1d1000x(b *testing.B) {
	for n := 0; n < 1000; n++ {
		result = Euler1d(100000)
	}
}
func BenchmarkEuler1e1000x(b *testing.B) {
	for n := 0; n < 1000; n++ {
		result = Euler1e(100000)
	}
}
func BenchmarkEuler1f1000x(b *testing.B) {
	for n := 0; n < 1000; n++ {
		result = Euler1f(100000)
	}
}
func BenchmarkEuler1g1000x(b *testing.B) {
	for n := 0; n < 1000; n++ {
		result = Euler1g(100000)
	}
}
func BenchmarkEuler1h1000x(b *testing.B) {
	for n := 0; n < 1000; n++ {
		result = Euler1h(100000)
	}
}
