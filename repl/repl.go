package repl

import (
	"bufio"
	"fmt"
	"io"
	lexer "ksm/Lexer"
)

const Type = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(Type)
		inputText := scanner.Scan()
		if !inputText {
			return
		}
		line := scanner.Text()

		lex := lexer.NewLexer(line) //creates lexer objects by breaking the users inmput into serirs of tokens

		for tnk := lex.NextToken(); tnk.Type != lexer.EOF; tnk = lex.NextToken() { //the loop get tokens one by one using len.nextToken() and runs until it reaches EOF
			fmt.Fprintf(out, "%+v\n", tnk) //"%+v" format is used to show the detailed information about the token, including its type and value
		}
	}
}
