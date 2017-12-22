package chapter23

import (
	"errors"
	"fmt"
	"reflect"
)

type TestResult struct {
	runCount   int
	errorCount int
}

func newTestResult() *TestResult {
	return &TestResult{runCount: 0, errorCount: 0}
}

func (t *TestResult) testStarted() {
	t.runCount++
}

func (t *TestResult) testFailed() {
	t.errorCount++
}

func (t *TestResult) summary() string {
	return fmt.Sprintf("%v run, %v failed", t.runCount, t.errorCount)
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
	res := fv.Call(nil)
	err := res[0].Interface()
	if err != nil {
		result.testFailed()
	}
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

func (w *WasRun) TestMethod() error {
	w.log += "testMethod "
	return nil
}

func (w *WasRun) TestBrokenMethod() error {
	return errors.New("Exception")
}

func (w *WasRun) tearDown() {
	w.log += "tearDown "
}
