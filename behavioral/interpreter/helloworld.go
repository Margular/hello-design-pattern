package interpreter

import (
	"fmt"
	"strings"
)

// STEP 1: Define an interface
type Func interface {
	call(string)
}

// STEP 2: Implements the interface
type MyPrintln struct {}

func (pl MyPrintln) call(param string) {
	fmt.Println(param)
}

// STEP 3: Define an interpreter object that have ability to invoke registered functions
type HelloInterpreter struct {
	funcs map[string]Func
}

func (hw *HelloInterpreter) RegFunc(name string, f Func) {
	if hw.funcs == nil {
		hw.funcs = map[string]Func{}
	}

	hw.funcs[name] = f
}

func (hw *HelloInterpreter) Interpret(expr string) {
	funcName := strings.TrimSpace(expr[:strings.Index(expr, "(")])
	f := hw.funcs[funcName]

	after := expr[strings.Index(expr, "(") + 1:]
	param := strings.Trim(after[:strings.LastIndex(after, ")")], "'")

	f.call(param)
}