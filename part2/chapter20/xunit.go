package chapter20

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
	w.setUp()
	method := t.name
	fv := reflect.ValueOf(w).MethodByName(method)
	fv.Call(nil)
}

type WasRun struct {
	wasRun bool
	log    string
}

func newWasRun() *WasRun {
	return &WasRun{}
}

func (w *WasRun) setUp() {
	w.wasRun = false
	w.log += "setUp "
}

func (w *WasRun) TestMethod() {
	w.wasRun = true
}
