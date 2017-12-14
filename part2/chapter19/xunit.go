package chapter19

import (
	"reflect"
)

type TestCase struct {
	name string
}

func newTestCase(name string) *TestCase {
	return &TestCase{name: name}
}

func (t *TestCase) run(w *WasRun) {
	method := t.name
	fv := reflect.ValueOf(w).MethodByName(method)
	fv.Call(nil)
}

type WasRun struct {
	wasRun bool
}

func newWasRun() *WasRun {
	return &WasRun{wasRun: false}
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}
