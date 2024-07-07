package console

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/internal/handlers"
)

func tokenize(args []string) {

	buffer, err := os.ReadFile(args[0])

	if err != nil {
		panic(err)
	}

	s := handlers.NewScanner(string(buffer))

	tokens := s.ScanTokens()

	for _, t := range tokens {
		fmt.Println(t.ToString())
	}
}
