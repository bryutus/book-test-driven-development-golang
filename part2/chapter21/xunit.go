package chapter21

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
	w.tearDown()
}

type WasRun struct {
	log string
}

func newWasRun() *WasRun {
	return &WasRun{}
}

func (w *WasRun) setUp() {
	w.log += "setUp "
}

func (w *WasRun) TestMethod() {
	w.log += "testMethod "
}

func (w *WasRun) tearDown() {
	w.log += "tearDown "
}
