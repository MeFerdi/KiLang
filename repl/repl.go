package repl

import (
	"bufio"
	"fmt"
	"io"
	lexer "ksm/Lexer"
)

const PROMT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Println(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		lex := lexer.NewLexer(line)

		for tnk := lex.NextToken(); tnk.Type != lexer.EOF; tnk = lex.NextToken() { // Use lexer.EOF to check for end of file
			fmt.Fprintf(out, "%+v\n", tnk) // Use fmt.Fprintf to write to the provided output stream
		}
	}
}
