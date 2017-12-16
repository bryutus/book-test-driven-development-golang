package chapter21

import (
	"testing"
)

func TestTemplateMethod(testing *testing.T) {
	testcase := newTestCase("TestMethod")
	test := newWasRun()

	testcase.run(test)

	if got, want := test.log, "setUp testMethod tearDown "; got != want {
		testing.Errorf("got %v want %v", got, want)
	}
}

func TestSummary(testing *testing.T) {
	testcase := newTestCase("TestMethod")
	test := newWasRun()

	result := testcase.run(test)

	if got, want := result.summary(), "1 run, 0 failed"; got != want {
		testing.Errorf("got %v want %v", got, want)
	}
}
