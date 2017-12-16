package chapter21

import (
	"fmt"
	"reflect"
)

type TestResult struct {
	runCount int
}

func newTestResult() *TestResult {
	return &TestResult{runCount: 1}

}

func (t *TestResult) summary() string {
	return fmt.Sprintf("%v run, 0 failed", t.runCount)
}

type TestCase struct {
	name string
}

func newTestCase(name string) *TestCase {
	return &TestCase{name: name}
}

func (t *TestCase) run(w *WasRun) *TestResult {
	w.setUp()
	method := t.name
	fv := reflect.ValueOf(w).MethodByName(method)
	fv.Call(nil)
	w.tearDown()
	return newTestResult()
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
