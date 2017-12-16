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
