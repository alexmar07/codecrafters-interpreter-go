package handlers

import (
	"fmt"
	"os"
)

func ErrorUnexpectedCharacter(line int, c string) {
	genericError(line, fmt.Sprintf("Unexpected character: %s", c))
}

func UnterminatedString(line int) {
	genericError(line, "Unterminated string.")
}

func genericError(line int, err string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", line, err)
}
