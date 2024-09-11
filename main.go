package main

import (
	"fmt"
	"ksm/lexer"
)

func main() {
	lex := lexer.NewLexer("hello")
	fmt.Printf("lexer: %v\n", lex)
}
