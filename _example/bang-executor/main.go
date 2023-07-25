package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/elk-language/go-prompt"
)

func main() {
	p := prompt.New(
		executor,
		prompt.WithPrefix(">>> "),
		prompt.WithExecuteOnEnterCallback(ExecuteOnEnter),
	)

	p.Run()
}

func ExecuteOnEnter(input string, indentSize int) (int, bool) {
	char, _ := utf8.DecodeLastRuneInString(input)
	return 0, char == '!'
}

func executor(s string) {
	fmt.Println("Your input: " + s)
}
