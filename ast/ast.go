package ast

import (
	lexer "ksm/Lexer"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}
type Expression interface {
	Node
	ExpressionNode()
}
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token lexer.Token
	Name  *Identifiers
	Value Expression
}

func (ls *LetStatement) StatementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifiers struct {
	Token lexer.Token
	Value string
}

func (i *Identifiers) ExpressionNode()      {}
func (i *Identifiers) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       lexer.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) StatementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type IntegerLiteral struct {
	Token lexer.Token 
	Value int64       
}

func (il *IntegerLiteral) ExpressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

type OperatorExpression struct {
	Left     Expression 
	Operator lexer.Token 
	Right    Expression  
}

func (oe *OperatorExpression) ExpressionNode()      {}
func (oe *OperatorExpression) TokenLiteral() string { return oe.Operator.Literal }
