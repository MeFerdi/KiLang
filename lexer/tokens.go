package lexer

import (
	"fmt"
)

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT      TokenType = "IDENT"
	INT        TokenType = "INT"
	STRING     TokenType = "STRING"
	FLOAT      TokenType = "FLOAT"
	SIGNED_INT TokenType = "SIGNED_INT"

	// Operators
	ASSIGN      TokenType = "ASSIGN"
	PLUS        TokenType = "PLUS"
	MINUS       TokenType = "MINUS"
	BANG        TokenType = "BANG"
	ASTERISK    TokenType = "ASTERIK"
	SLASH       TokenType = "SLASH"
	LESS        TokenType = "LESS"
	LESS_EQUAL  TokenType = "LESS_EQUAL"
	GREAT       TokenType = "GREAT"
	GREAT_EQUAL TokenType = "GREAT_EQUAL"
	AND         TokenType = "AND"

	// Delimiters
	COMMA        TokenType = "COMMA"
	SEMICOLON    TokenType = "SEMICOLON"
	LPAREN       TokenType = "LPAREN"
	RPAREN       TokenType = "RPAREN"
	LBRACE       TokenType = "LBRACE"
	RBRACE       TokenType = "RBRACE"
	PERIOD       TokenType = "PERIOD"
	LSQUAREBRACE TokenType = "LSQUAREBRACE"
	RSQUAREBRACE TokenType = "RSQUAREBRACE"

	// Keywords
	LET      TokenType = "LET"
	FUNCTION TokenType = "FUNCTION"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	NULL     TokenType = "NULL"
)

type Token struct {
	Type    TokenType
	Literal string
}

// IsToken checks if the token type matches any in the provided list.
func (t Token) IsToken(expectedTokens ...TokenType) bool {
	for _, expected := range expectedTokens {
		if expected == t.Type {
			return true
		}
	}
	return false
}

// Debug prints a debug string representation of the token.
func (t Token) Debug() {
	if t.IsToken(IDENT, INT, STRING) {
		fmt.Printf("%s (%s)\n", TokenTypeString(t.Type), t.Literal)
	} else {
		fmt.Printf("%s ()\n", TokenTypeString(t.Type))
	}
}

// TokenTypeString converts a TokenType to its string representation.
func TokenTypeString(types TokenType) string {
	switch types {
	case EOF:
		return "eof"
	case STRING:
		return "string"
	case INT:
		return "int"
	case IDENT:
		return "ident"
	case ASSIGN:
		return "assign"
	case PLUS:
		return "plus"
	case MINUS:
		return "minus"
	case ASTERISK:
		return "asterisk"
	case SLASH:
		return "slash"
	case LESS:
		return "less"
	case LESS_EQUAL:
		return "less_equal"
	case GREAT:
		return "great"
	case GREAT_EQUAL:
		return "great_equal"
	case AND:
		return "and"
	case COMMA:
		return "comma"
	case SEMICOLON:
		return "semicolon"
	case LPAREN:
		return "lparen"
	case RPAREN:
		return "rparen"
	case LBRACE:
		return "lbrace"
	case RBRACE:
		return "rbrace"
	case FUNCTION:
		return "function"
	case IF:
		return "if"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case RETURN:
		return "return"
	case LSQUAREBRACE:
		return "lsquarebrace"
	case RSQUAREBRACE:
		return "rsquarebrace"
	case PERIOD:
		return "period"
	case NULL:
		return "null"
	default:
		return string(types)
	}
}
