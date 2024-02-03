package parser

import (
	"kol/ast"
	"kol/token"
)

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken, Mutable: false}

	if p.peekTokenIs(token.MUT) {
		stmt.Mutable = true
		p.nextToken()
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
func (p *Parser) parseReassignStatement() ast.Statement {
	stmt := &ast.ReassignStatement{Token: p.curToken}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	p.nextToken()
	tok := p.curToken
	p.nextToken()

	stmt.Value = getValue(*stmt.Name, tok, p.parseExpression(LOWEST))

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
func getValue(ident ast.Identifier, tok token.Token, exp ast.Expression) ast.Expression {
	switch tok.Type {
	case token.ASSIGN:
		return exp
	case token.PLUSASS:
		return &ast.InfixExpression{
			Token:    token.Token{Type: token.PLUS, Literal: "+", Position: tok.Position},
			Left:     &ident,
			Right:    exp,
			Operator: "+",
		}
	case token.MINASS:
		return &ast.InfixExpression{
			Token:    token.Token{Type: token.MINUS, Literal: "+", Position: tok.Position},
			Left:     &ident,
			Right:    exp,
			Operator: "-",
		}
	case token.MULTASS:
		return &ast.InfixExpression{
			Token:    token.Token{Type: token.ASTERISK, Literal: "+", Position: tok.Position},
			Left:     &ident,
			Right:    exp,
			Operator: "*",
		}
	case token.DIVASS:
		return &ast.InfixExpression{
			Token:    token.Token{Type: token.SLASH, Literal: "/", Position: tok.Position},
			Left:     &ident,
			Right:    exp,
			Operator: "/",
		}
	default:
		return nil
	}
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	if p.peekTokenIs(token.SEMICOLON) || p.peekTokenIs(token.RPAREN) || p.peekTokenIs(token.RBRACKET) || p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		stmt.ReturnValue = nil
	} else {
		p.nextToken()
		stmt.ReturnValue = p.parseExpression(LOWEST)
	}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
func (p *Parser) parseBreakStatement() ast.Statement {
	stmt := &ast.BreakStatement{Token: p.curToken}

	p.nextToken()

	stmt.BreakValue = p.parseExpression(LOWEST)

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}
	p.nextToken()
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}
