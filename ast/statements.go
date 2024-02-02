package ast

import (
	"bytes"
	"kol/token"
)

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token   token.Token
	Name    *Identifier
	Value   Expression
	Mutable bool
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	if ls.Mutable {
		out.WriteString("mut ")
	}
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
func (ls *LetStatement) GetPosition() token.Position { return ls.Token.Position }

type ReassignStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (rs *ReassignStatement) statementNode()       {}
func (rs *ReassignStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReassignStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.Name.String())
	out.WriteString(" = ")
	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
func (rs *ReassignStatement) GetPosition() token.Position { return rs.Token.Position }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
func (rs *ReturnStatement) GetPosition() token.Position { return rs.Token.Position }

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) GetPosition() token.Position { return es.Token.Position }

type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
func (bs *BlockStatement) GetPosition() token.Position { return bs.Token.Position }
