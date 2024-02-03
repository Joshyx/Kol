package ast

import (
	"bytes"
	"kol/token"
)

type Node interface {
	TokenLiteral() string
	String() string
	GetPosition() token.Position
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
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
func (p *Program) GetPosition() token.Position { return token.Position{Line: 1, Column: 1} }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()             {}
func (i *Identifier) TokenLiteral() string        { return i.Token.Literal }
func (i *Identifier) String() string              { return i.Value }
func (i *Identifier) GetPosition() token.Position { return i.Token.Position }

var Types = []string{
	"int",
	"float",
	"bool",
	"str",
	"map",
	"array",
	"float",
	"fn",
	"void",
}
