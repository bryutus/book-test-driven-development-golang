package chapter18

import (
	"testing"
)

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
