package console

import "fmt"

type fnc func(args []string)

type Interpreter struct {
	fncs map[string]fnc
}

type FunctionNotFoundError struct {
	nameFnc string
}

func (e *FunctionNotFoundError) Error() string {
	return fmt.Sprintf("%s: function not found", e.nameFnc)
}

func NewInterpreter() *Interpreter {

	i := &Interpreter{make(map[string]fnc)}

	i.boot()

	return i
}

func (i *Interpreter) register(nameFnc string, function fnc) {
	i.fncs[nameFnc] = function
}

func (i *Interpreter) boot() {
	i.register("tokenize", tokenize)
}

func (i *Interpreter) exist(nameFnc string) bool {

	_, ok := i.fncs[nameFnc]

	return ok
}

func (i *Interpreter) GetFnc(nameFnc string) (fnc, error) {

	if !i.exist(nameFnc) {
		return nil, &FunctionNotFoundError{nameFnc}
	}

	return i.fncs[nameFnc], nil
}
