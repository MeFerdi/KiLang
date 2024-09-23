package parser

import lexer "ksm/Lexer"

type Parser struct {
	lexer        *lexer.Lexer
	currentToken lexer.Token
}

// NewParser initializes a new Parser with the given lexer
func NewParser(lex *lexer.Lexer) *Parser {
	p := &Parser{lexer: lex}
	p.nextToken() // advance to the first token
	return p
}

// nextToken advances the parser to the next token
func (p *Parser) nextToken() {
	p.currentToken = p.lexer.NextToken()
}

func (p *Parser) parseExpression() Expression {
	switch p.currentToken.Type {
	case lexer.INT, lexer.FLOAT:
		return p.parseNumber()
	case lexer.IDENT:
		return p.parseIdentifier()
	default:
		return nil // handle errors appropriately
	}
}

func (p *Parser) parseNumber() Expression {
	// Create a Number node
	// define a Number type
	node := &NumberLiteral{Value: p.currentToken.Literal}
	p.nextToken() // advance to the next token
	return node
}

func (p *Parser) parseIdentifier() Expression {
	// Create an Identifier node
	// define an Identifier type
	node := &Identifier{Name: p.currentToken.Literal}
	p.nextToken() // advance to the next token
	return node
}

func (p *Parser) parseStatement() Statement {
	switch p.currentToken.Type {
	case lexer.FUNCTION:
		return p.parseFunctionStatement()
	case lexer.IDENT:
		return p.parseAssignmentStatement()
	default:
		return nil
	}
}

func (p *Parser) parseAssignmentStatement() Statement {
	//Assignment nodes
	node := &AssignmentStatement{Name: p.currentToken.Literal}
	p.nextToken()
	if p.currentToken.Type != lexer.ASSIGN {
		return nil // handle error for missing assignment operator
	}
	p.nextToken() // advance to the value being assigned
	node.Value = p.parseExpression()
	return node
}

func (p *Parser) parseFunctionStatement() Statement {
	// This typically includes parsing the function name, parameters, and body
	return nil
}

// Resulting AST Implementations

type NumberLiteral struct {
	Value string // or can be float64/int
}

type AssignmentStatement struct {
	Name  string
	Value Expression
}

func (p *Parser) Parse() []Statement {
	var statements []Statement
	for p.currentToken.Type != lexer.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			statements = append(statements, stmt)
		}
		p.nextToken() // move to the next token
	}
	return statements
}
