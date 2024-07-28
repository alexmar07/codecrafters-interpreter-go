package handlers

import (
	"fmt"
	"os"
)

func ErrorUnexpectedCharacter(line int, c string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", line, c)
}
