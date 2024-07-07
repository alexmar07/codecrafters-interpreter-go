package handlers

import "fmt"

func ErrorUnexpectedCharacter(line int, c string) {
	fmt.Printf("[line %d] Error: Unexpected character: %s\n", line, c)
}
