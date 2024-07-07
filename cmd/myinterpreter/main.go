package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/internal/console"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh <command> <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	itp := console.NewInterpreter()

	fnc, err := itp.GetFnc(command)

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	fnc(os.Args[2:])
}
