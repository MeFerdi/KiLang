package parser

import (
	ast "ksm/Ast"
	lexer "ksm/Lexer"
	"strconv"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken lexer.Token
	peekToken    lexer.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l} //instanct of passer is creaded to bind lexer to parser to enable parser fetch token from lexer

	p.NextToken() //newtoken is called twice to be populate currenttoken and peek token
	p.NextToken() //ensures parser always has 2 tokens, the current one its processing and the next one for lookahead

	return p
}
func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()

}

// ParseProgram handles the entire program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	p.NextToken() // Start fetching tokens

	// Loop through all tokens until EOF
	for p.currentToken.Type != lexer.EOF {
		statement := p.parseStatement() // Parse different types of statements
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.NextToken() // Fetch next token
	}
	return program
}

// ParseStatement decides which specific statement parser to call
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case lexer.LET:
		return p.parseLetStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// ParseLetStatement parses let statements like `let x = 5;`
func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}
	if !p.expectPeek(lexer.IDENT) {
		return nil
	}
	statement.Name = p.parseIdentifier()

	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}

	p.NextToken() // Move to the value of the let statement
	statement.Value = p.parseExpression()
	return statement
}

func (p *Parser) parseExpression() ast.Expression {
	switch p.currentToken.Type {
	case lexer.INT:
		return p.parseIntegerLiteral()
	case lexer.PLUS:
		return p.parseOperatorExpression()
	default:
		return nil
	}
}

func (p *Parser) parseIdentifier() *ast.Identifiers {
	return &ast.Identifiers{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseOperatorExpression() ast.Expression {
	expression := &ast.OperatorExpression{
		Left:     p.parseIntegerLiteral(),
		Operator: p.currentToken,
	}
	p.NextToken() // Move to next part of expression (right side)
	expression.Right = p.parseExpression()
	return expression
}
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.NextToken()

	stmt.ReturnValue = p.parseExpression()

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.NextToken()
	}

	return stmt
}
func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekToken.Type == t {
		p.NextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		return nil
	}

	lit.Value = value
	return lit
}
