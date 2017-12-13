package chapter18

import (
	"reflect"
	"testing"
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

func TestRunning(testing *testing.T) {
	testcase := newTestCase("TestMethod")
	test := newWasRun()

	if got, want := test.wasRun, false; got != want {
		testing.Errorf("got %v want %v", got, want)
	}

	testcase.run(test)

	if got, want := test.wasRun, true; got != want {
		testing.Errorf("got %v want %v", got, want)
	}
}
