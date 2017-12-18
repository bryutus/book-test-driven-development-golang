package chapter22

import (
	"errors"
	"fmt"
	"reflect"
)

type TestResult struct {
	runCount int
}

func newTestResult() *TestResult {
	return &TestResult{runCount: 0}

}

func (t *TestResult) testStarted() {
	t.runCount++
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
	result := newTestResult()
	result.testStarted()
	w.setUp()
	method := t.name
	fv := reflect.ValueOf(w).MethodByName(method)
	fv.Call(nil)
	w.tearDown()
	return result
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

func (w *WasRun) TestBrokenMethod() error {
	return errors.New("Exception")
}

func (w *WasRun) tearDown() {
	w.log += "tearDown "
}
