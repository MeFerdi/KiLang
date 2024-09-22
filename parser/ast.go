package parser

// Statement interface - all Statements should implement this
type Statement interface{}

// Expression interface - all Expressions should implement this
type Expression interface{}

// FunctionStatement represents a function declaration
type FunctionStatement struct {
	Name       string      // Name of the function
	Parameters []string    // List of parameter names
	Body       []Statement // List of statements in the function body
}

// ReturnStatement represents a return statement
type ReturnStatement struct {
	Value Expression // The expression to return
}

// BlockStatement represents a block of statements
type BlockStatement struct {
	Statements []Statement // A list of statements within the block
}

// Identifier represents a variable or function name
type Identifier struct {
	Name string // The name of the identifier
}

// BinaryExpression represents a binary operation like a + b
type BinaryExpression struct {
	Left     Expression // Left-hand side expression
	Operator string     // Operator (e.g., "+", "-", etc.)
	Right    Expression // Right-hand side expression
}

type Number struct {
	Value float64 // The numeric value
}
