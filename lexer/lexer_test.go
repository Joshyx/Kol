package lexer

import (
	"testing"

	"kol/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;//asdasdasdasd                        

let add = fun(x, y) {
  x /*asd*/+ y;
};

let result = add(five, ten);
!-/ *5;
5 < 10 > 5;

if (5 < 10) {
    return true;
} else {
    /*Is this comment being ignored?*/ return false;
}

10 /*asd*/ == 10;
10 != 9;
<= >=
"foobar"
"foo bar"
[1, 2]; // This comment should be ignored
{"foo": "bar"}
123.2343453
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedChar    int
	}{
		{token.LET, "let", 1, 1},
		{token.IDENT, "five", 1, 5},
		{token.ASSIGN, "=", 1, 10},
		{token.INT, "5", 1, 12},
		{token.SEMICOLON, ";", 1, 13},
		{token.LET, "let", 2, 1},
		{token.IDENT, "ten", 2, 5},
		{token.ASSIGN, "=", 2, 9},
		{token.INT, "10", 2, 11},
		{token.SEMICOLON, ";", 2, 13},
		{token.LET, "let", 4, 1},
		{token.IDENT, "add", 4, 5},
		{token.ASSIGN, "=", 4, 9},
		{token.FUNCTION, "fun", 4, 11},
		{token.LPAREN, "(", 4, 14},
		{token.IDENT, "x", 4, 15},
		{token.COMMA, ",", 4, 16},
		{token.IDENT, "y", 4, 18},
		{token.RPAREN, ")", 4, 19},
		{token.LBRACE, "{", 4, 21},
		{token.IDENT, "x", 5, 3},
		{token.PLUS, "+", 5, 12},
		{token.IDENT, "y", 5, 14},
		{token.SEMICOLON, ";", 5, 15},
		{token.RBRACE, "}", 6, 1},
		{token.SEMICOLON, ";", 6, 2},
		{token.LET, "let", 8, 1},
		{token.IDENT, "result", 8, 5},
		{token.ASSIGN, "=", 8, 12},
		{token.IDENT, "add", 8, 14},
		{token.LPAREN, "(", 8, 17},
		{token.IDENT, "five", 8, 18},
		{token.COMMA, ",", 8, 22},
		{token.IDENT, "ten", 8, 24},
		{token.RPAREN, ")", 8, 27},
		{token.SEMICOLON, ";", 8, 28},
		{token.BANG, "!", 9, 1},
		{token.MINUS, "-", 9, 2},
		{token.SLASH, "/", 9, 3},
		{token.ASTERISK, "*", 9, 5},
		{token.INT, "5", 9, 6},
		{token.SEMICOLON, ";", 9, 7},
		{token.INT, "5", 10, 1},
		{token.LT, "<", 10, 3},
		{token.INT, "10", 10, 5},
		{token.GT, ">", 10, 8},
		{token.INT, "5", 10, 10},
		{token.SEMICOLON, ";", 10, 11},
		{token.IF, "if", 12, 1},
		{token.LPAREN, "(", 12, 4},
		{token.INT, "5", 12, 5},
		{token.LT, "<", 12, 7},
		{token.INT, "10", 12, 9},
		{token.RPAREN, ")", 12, 11},
		{token.LBRACE, "{", 12, 13},
		{token.RETURN, "return", 13, 5},
		{token.TRUE, "true", 13, 12},
		{token.SEMICOLON, ";", 13, 16},
		{token.RBRACE, "}", 14, 1},
		{token.ELSE, "else", 14, 3},
		{token.LBRACE, "{", 14, 8},
		{token.RETURN, "return", 15, 40},
		{token.FALSE, "false", 15, 47},
		{token.SEMICOLON, ";", 15, 52},
		{token.RBRACE, "}", 16, 1},
		{token.INT, "10", 18, 1},
		{token.EQ, "==", 18, 12},
		{token.INT, "10", 18, 15},
		{token.SEMICOLON, ";", 18, 17},
		{token.INT, "10", 19, 1},
		{token.NOT_EQ, "!=", 19, 4},
		{token.INT, "9", 19, 7},
		{token.SEMICOLON, ";", 19, 8},
		{token.LTEQ, "<=", 20, 1},
		{token.GTEQ, ">=", 20, 4},
		{token.STRING, "foobar", 21, 1},
		{token.STRING, "foo bar", 22, 1},
		{token.LBRACKET, "[", 23, 1},
		{token.INT, "1", 23, 2},
		{token.COMMA, ",", 23, 3},
		{token.INT, "2", 23, 5},
		{token.RBRACKET, "]", 23, 6},
		{token.SEMICOLON, ";", 23, 7},
		{token.LBRACE, "{", 24, 1},
		{token.STRING, "foo", 24, 2},
		{token.COLON, ":", 24, 7},
		{token.STRING, "bar", 24, 9},
		{token.RBRACE, "}", 24, 14},
		{token.FLOAT, "123.2343453", 25, 1},
		{token.EOF, "", 26, 1},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q (value=%q)",
				i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
		if tok.Position.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line number wrong. expected=%d, got=%d",
				i, tt.expectedLine, tok.Position.Line)
		}
		if tok.Position.Column != tt.expectedChar {
			t.Fatalf("tests[%d] - char index wrong. expected=%d, got=%d",
				i, tt.expectedChar, tok.Position.Column)
		}
	}
}
