package main

import (
	"fmt"
	"ksm/lexer"
)

func main() {
	lex := lexer.NewLexer("malika")
	
	fmt.Printf("lexer: %v\n", lex)
}

// fmt.Println((lex))
// file, _ := os.ReadFile("./tests/file.ksm")
// code := string(file)