package parser

import (
	"fmt"
	"kol/token"
	"testing"
)

func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) peekError(t token.TokenType, pos token.Position) {
	p.addError("expected next token to be %s, got %s instead", pos, t, p.peekToken.Type)
}
func (p *Parser) noPrefixParseFnError(t token.TokenType, pos token.Position) {
	p.addError("no prefix parse function for %s found", pos, t)
}
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
func (p *Parser) addError(msg string, pos token.Position, a ...interface{}) {
	newMsg := fmt.Sprintf(msg, a...)
	p.errors = append(p.errors, fmt.Sprintf("Parser error at %d:%d: %s", pos.Line, pos.Column, newMsg))
}
