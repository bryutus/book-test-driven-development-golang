package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := newWasRun("TestMethod")
	fmt.Println(test.wasRun)
	test.run()
	fmt.Println(test.wasRun)
}

type TestCase struct {
	name string
}

func newTestCase(name string) TestCase {
	return TestCase{name: name}
}

type WasRun struct {
	wasRun bool
	TestCase
}

func newWasRun(name string) WasRun {
	return WasRun{wasRun: false, TestCase: newTestCase(name)}
}

func (w *WasRun) run() {
	method := w.name
	fv := reflect.ValueOf(w).MethodByName(method)
	fv.Call(nil)
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}
