package parser

type Parser struct {
	l            *lexer.Lexer
	currentToken lexer.Token
	peekToken    lexer.Token
}
func (l *lexer.Lexer) *Parser {
	p := &Parser{l: l} //instanct of passer is creaded to bind lexer to parser to enable parser fetch token from lexer

	p := NextToken() //newtoken is called twise to be porpulate currenttoken and peek token
	p := NextToken() //ensures parser always has 2 tokens, the current one its processing and the next one for lookahead

	return p
}
func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()

}
