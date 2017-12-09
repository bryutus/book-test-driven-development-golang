package main

import "fmt"

func main() {
	test := newWasRun("testMethod")
	fmt.Println(test.wasRun)
	test.run()
	fmt.Println(test.wasRun)
}

type WasRun struct {
	name   string
	wasRun bool
}

func newWasRun(name string) WasRun {
	return WasRun{name: name, wasRun: false}
}

func (w *WasRun) run() {
	w.testMethod()
}

func (w *WasRun) testMethod() {
	w.wasRun = true
}
