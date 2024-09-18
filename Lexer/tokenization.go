package lexer

type Lexer struct {
	input               string //the source code lexer reading and a string containing the entire code to be broken into token
	position            int    // keeps track of current position in the input and points to the character
	readCurrentPosition int    //current reading position in the input/is the index in the input where the lexer will read the next character.
	ch                  byte   // current character to be tokenized
}

// initializes the lexer with the input string and sets the initial state.
func NewLexer(input string) *Lexer { //instance
	lex := &Lexer{input: input}
	lex.readChar() // It calls readChar to load the first character into ch.
	return lex
}

func (lex *Lexer) readChar() { //This function advances the position and readCurrentPosition pointers in the input string and updates ch to the current character.
	if lex.readCurrentPosition >= len(lex.input) {
		lex.ch = 0 // sets ch to 0 if readposition go beyond len of input indicating EOL
	} else {
		lex.ch = lex.input[lex.readCurrentPosition]
	}
	lex.position = lex.readCurrentPosition
	lex.readCurrentPosition++
}

func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
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

func newToken(tokenType TokenType, ch byte) Token { //core method
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
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
			tok = Token{Type: GREAT_EQUALS, Literal: string(ch) + string(lex.ch)}
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
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.readIdentifier()
			tok.Type = lex.lookupIdent(tok.Literal)
			return tok
		} else if isDigit(lex.ch) || (lex.ch == '-' || lex.ch == '+') {
			if lex.peekChar() == '.' {
				tok.Type = FLOAT
				tok.Literal = lex.readFloat(lex.position)
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

	if lex.ch == '-' || lex.ch == '+' { // Check for optional sign
		lex.readChar() // consume the sign
	}

	for isDigit(lex.ch) {
		lex.readChar()
	}

	return lex.input[position:lex.position]
}

// New function to read floating-point numbers
func (lex *Lexer) readFloat(startPosition int) string {
	position := startPosition
	// Consume the decimal point.
	if lex.ch == '.' {
		lex.readChar()
	}

	// Read digits after the decimal point.
	for isDigit(lex.ch) {
		lex.readChar()
	}

	return lex.input[position:lex.position]
}

func (lex *Lexer) lookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"fn":     FUNCTION,
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
