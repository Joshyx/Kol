package parser

import (
	"fmt"
	"kol/ast"
	"kol/token"
	"slices"
	"strconv"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}
func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.curToken}
	value, err := strconv.ParseFloat(p.curToken.Literal, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.BooleanLiteral{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	if p.peekTokenIs(token.IDENT) {
		return nil
	}
	lit := &ast.FunctionLiteral{Token: p.curToken}
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	lit.Parameters = p.parseFunctionParameters()
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		if !slices.Contains(ast.Types, p.curToken.Literal) {
			panic(fmt.Sprintf("Can't find type with name %s", p.curToken.Literal))
		}

		lit.ReturnType = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else {
		lit.ReturnType = &ast.Identifier{Token: p.curToken, Value: "void"}
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	lit.Body = p.parseBlockStatement()
	return lit
}
func (p *Parser) parseFunction() ast.Statement {
	lit := &ast.FunctionLiteral{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	ident := p.curToken
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	lit.Parameters = p.parseFunctionParameters()
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		if !slices.Contains(ast.Types, p.curToken.Literal) {
			panic(fmt.Sprintf("Can't find type with name %s", p.curToken.Literal))
		}

		lit.ReturnType = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else {
		lit.ReturnType = &ast.Identifier{Token: p.curToken, Value: "void"}
	}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	lit.Body = p.parseBlockStatement()
	return &ast.LetStatement{
		Token:   token.Token{Type: token.LET, Literal: "let", Position: lit.GetPosition()},
		Name:    &ast.Identifier{Token: ident, Value: ident.Literal},
		Value:   lit,
		Mutable: false,
	}
}

func (p *Parser) parseFunctionParameters() []*ast.FunctionParameter {
	arguments := []*ast.FunctionParameter{}
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return arguments
	}
	p.nextToken()
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	p.nextToken()
	if !slices.Contains(ast.Types, p.curToken.Literal) {
		panic(fmt.Sprintf("Can't find type with name %s", p.curToken.Literal))
	}
	expType := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	arguments = append(arguments, &ast.FunctionParameter{Ident: *ident, Type: *expType})
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		p.nextToken()
		if !slices.Contains(ast.Types, p.curToken.Literal) {
			panic(fmt.Sprintf("Can't find type with name %s", p.curToken.Literal))
		}
		expType := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		arguments = append(arguments, &ast.FunctionParameter{Ident: *ident, Type: *expType})
	}
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return arguments
}
func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.curToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)
	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)
		if !p.expectPeek(token.COLON) {
			return nil
		}
		p.nextToken()
		value := p.parseExpression(LOWEST)
		hash.Pairs[key] = value
		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}
	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return hash
}
