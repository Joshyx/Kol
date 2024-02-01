package lexer

import (
	"kol/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int  // current pos in input (points to current char)
	readPosition int  // current reading pos in input (after current char)
	ch           byte // current char under examination

	curLine int
	curChar int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, curLine: 1, curChar: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.ch == '\n' {
		l.curLine += 1
		l.curChar = 1
	} else {
		l.curChar += 1
	}
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	l.skipComment()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:     token.EQ,
				Literal:  string(ch) + string(l.ch),
				Position: token.Position{Line: l.curLine, Char: l.curChar - 1},
			}
		} else {
			tok = l.getToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = l.getToken(token.SEMICOLON, l.ch)
	case ':':
		tok = l.getToken(token.COLON, l.ch)
	case '(':
		tok = l.getToken(token.LPAREN, l.ch)
	case ')':
		tok = l.getToken(token.RPAREN, l.ch)
	case '+':
		tok = l.getToken(token.PLUS, l.ch)
	case '-':
		tok = l.getToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:     token.NOT_EQ,
				Literal:  string(ch) + string(l.ch),
				Position: token.Position{Line: l.curLine, Char: l.curChar - 1},
			}
		} else {
			tok = l.getToken(token.BANG, l.ch)
		}
	case '/':
		tok = l.getToken(token.SLASH, l.ch)
	case '*':
		tok = l.getToken(token.ASTERISK, l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:     token.LTEQ,
				Literal:  string(ch) + string(l.ch),
				Position: token.Position{Line: l.curLine, Char: l.curChar - 1},
			}
		} else {
			tok = l.getToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:     token.GTEQ,
				Literal:  string(ch) + string(l.ch),
				Position: token.Position{Line: l.curLine, Char: l.curChar - 1},
			}
		} else {
			tok = l.getToken(token.GT, l.ch)
		}
	case ',':
		tok = l.getToken(token.COMMA, l.ch)
	case '.':
		tok = l.getToken(token.PERIOD, l.ch)
	case '{':
		tok = l.getToken(token.LBRACE, l.ch)
	case '}':
		tok = l.getToken(token.RBRACE, l.ch)
	case '[':
		tok = l.getToken(token.LBRACKET, l.ch)
	case ']':
		tok = l.getToken(token.RBRACKET, l.ch)
	case '"':
		tok.Position = token.Position{Line: l.curLine, Char: l.curChar}
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Position = token.Position{Line: l.curLine, Char: l.curChar}
	default:
		tok.Position = token.Position{Line: l.curLine, Char: l.curChar}
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			if strings.Contains(tok.Literal, ".") {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}
			return tok
		} else {
			tok = l.getToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) getToken(tok token.TokenType, ch byte) token.Token {
	return token.New(tok, ch, token.Position{Line: l.curLine, Char: l.curChar})
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipComment() {
	if l.ch != '/' {
		return
	}

	if l.peekChar() == '/' {
		for l.ch != '\n' {
			l.readChar()
		}
		l.skipWhitespace()
	} else if l.peekChar() == '*' {
		for !(l.ch == '*' && l.peekChar() == '/') {
			l.readChar()
		}
		l.readChar()
		l.readChar()
		l.skipWhitespace()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	foundDecPoint := false
	for {
		if isDigit(l.ch) {
			l.readChar()
			continue
		}
		if !foundDecPoint && l.ch == '.' {
			foundDecPoint = true
			l.readChar()
			continue
		}
		break
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
