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

type WasRun struct {
	name   string
	wasRun bool
}

func newWasRun(name string) WasRun {
	return WasRun{name: name, wasRun: false}
}

func (w *WasRun) run() {
	method := w.name
	fv := reflect.ValueOf(w).MethodByName(method)
	fv.Call(nil)
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}
