package lexer

import (
	"strings"
	"unicode"
)

type Lexer struct {
	input               string // the source code lexer reading
	position            int    // keeps track of current position in the input
	readCurrentPosition int    // current reading position in the input
	ch                  byte   // current character to be tokenized
}

// initializes the lexer with the input string and sets the initial state.
func NewLexer(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readCurrentPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readCurrentPosition]
	}
	lex.position = lex.readCurrentPosition
	lex.readCurrentPosition++
}

func (lex *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(lex.ch)) {
		lex.readChar()
	}
}

func (lex *Lexer) peekChar() byte {
	if lex.readCurrentPosition >= len(lex.input) {
		return 0
	} else {
		return lex.input[lex.readCurrentPosition]
	}
}

func (lex *Lexer) readIdentifier() string {
	position := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

func (lex *Lexer) NextToken() Token {
	var tok Token

	lex.skipWhitespace()

	switch lex.ch {
	case '=':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			tok = Token{Type: ASSIGN, Literal: string(ch) + string(lex.ch)}
		} else {
			tok = newToken(ASSIGN, lex.ch)
		}
	case '+':
		tok = newToken(PLUS, lex.ch)
	case '-':
		tok = newToken(MINUS, lex.ch)
	case '!':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			tok = Token{Type: BANG, Literal: string(ch) + string(lex.ch)}
		} else {
			tok = newToken(BANG, lex.ch)
		}
	case '/':
		tok = newToken(SLASH, lex.ch)
	case '*':
		tok = newToken(ASTERISK, lex.ch)
	case '<':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			tok = Token{Type: LESS_EQUAL, Literal: string(ch) + string(lex.ch)}
		} else {
			tok = newToken(LESS, lex.ch)
		}
	case '>':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			tok = Token{Type: GREAT_EQUAL, Literal: string(ch) + string(lex.ch)}
		} else {
			tok = newToken(GREAT, lex.ch)
		}
	case '&':
		tok = newToken(AND, lex.ch)
	case ';':
		tok = newToken(SEMICOLON, lex.ch)
	case '(':
		tok = newToken(LPAREN, lex.ch)
	case ')':
		tok = newToken(RPAREN, lex.ch)
	case ',':
		tok = newToken(COMMA, lex.ch)
	case '{':
		tok = newToken(LBRACE, lex.ch)
	case '}':
		tok = newToken(RBRACE, lex.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	case '[':
		tok = newToken(LSQUAREBRACE, lex.ch)
	case ']':
		tok = newToken(RSQUAREBRACE, lex.ch)
	case '.':
		// Handle case where a float might start with a decimal point
		if isDigit(lex.peekChar()) {
			tok.Type = FLOAT
			tok.Literal = lex.readFloat()
		} else {
			tok = newToken(PERIOD, lex.ch)
		}
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifier()
			tok.Type = lex.lookupIdent(tok.Literal)
			return tok
		} else if isDigit(lex.ch) {
			if lex.peekChar() == '.' || strings.ContainsAny(string(lex.ch), "eE") {
				tok.Type = FLOAT
				tok.Literal = lex.readFloat()
			} else {
				tok.Type = SIGNED_INT
				tok.Literal = lex.readNumber()
			}
			return tok
		} else {
			tok = newToken(ILLEGAL, lex.ch)
		}
	}

	lex.readChar()
	return tok
}

func (lex *Lexer) readNumber() string {
	position := lex.position

	if lex.ch == '-' || lex.ch == '+' {
		lex.readChar()
	}

	for isDigit(lex.ch) {
		lex.readChar()
	}

	return lex.input[position:lex.position]
}

// Updated function to read floating-point numbers
func (lex *Lexer) readFloat() string {
	position := lex.position

	if lex.ch == '-' || lex.ch == '+' {
		lex.readChar()
	}

	if lex.ch == '.' {
		lex.readChar()
	}

	for isDigit(lex.ch) {
		lex.readChar()
	}

	if lex.ch == 'e' || lex.ch == 'E' {
		lex.readChar()
		if lex.ch == '+' || lex.ch == '-' {
			lex.readChar()
		}
		for isDigit(lex.ch) {
			lex.readChar()
		}
	}

	return lex.input[position:lex.position]
}

func (lex *Lexer) lookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"func":   FUNCTION,
		"true":   TRUE,
		"false":  FALSE,
		"if":     IF,
		"else":   ELSE,
		"return": RETURN,
		"null":   NULL,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
