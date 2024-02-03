package ast

import (
	"bytes"
	"kol/token"
	"strings"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()             {}
func (il *IntegerLiteral) TokenLiteral() string        { return il.Token.Literal }
func (il *IntegerLiteral) String() string              { return il.Token.Literal }
func (il *IntegerLiteral) GetPosition() token.Position { return il.Token.Position }

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode()             {}
func (fl *FloatLiteral) TokenLiteral() string        { return fl.Token.Literal }
func (fl *FloatLiteral) String() string              { return fl.Token.Literal }
func (fl *FloatLiteral) GetPosition() token.Position { return fl.Token.Position }

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (b *BooleanLiteral) expressionNode()             {}
func (b *BooleanLiteral) TokenLiteral() string        { return b.Token.Literal }
func (b *BooleanLiteral) String() string              { return b.Token.Literal }
func (b *BooleanLiteral) GetPosition() token.Position { return b.Token.Position }

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*FunctionParameter
	Body       *BlockStatement
	ReturnType *Identifier
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.ReturnType.String() + " ")
	out.WriteString(fl.Body.String())
	return out.String()
}
func (fl *FunctionLiteral) GetPosition() token.Position { return fl.Token.Position }

type FunctionParameter struct {
	Token token.Token
	Ident Identifier
	Type  Identifier
}

func (fp *FunctionParameter) expressionNode()      {}
func (fp *FunctionParameter) TokenLiteral() string { return fp.Token.Literal }
func (fp *FunctionParameter) String() string {
	return fp.Ident.String() + " " + string(fp.Type.Value)
}
func (fp *FunctionParameter) GetPosition() token.Position { return fp.Token.Position }

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()             {}
func (sl *StringLiteral) TokenLiteral() string        { return sl.Token.Literal }
func (sl *StringLiteral) String() string              { return sl.Token.Literal }
func (sl *StringLiteral) GetPosition() token.Position { return sl.Token.Position }

type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
func (al *ArrayLiteral) GetPosition() token.Position { return al.Token.Position }

type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
func (hl *HashLiteral) GetPosition() token.Position { return hl.Token.Position }
