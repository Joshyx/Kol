package token

type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	Position Position
}
type Position struct {
	Line int
	Char int
}

func New(tokenType TokenType, ch byte, pos Position) Token {
	return Token{Type: tokenType, Literal: string(ch), Position: pos}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1343456
	FLOAT  = "FLOAT"
	STRING = "STRING"
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	EQ     = "=="
	NOT_EQ = "!="

	LTEQ = "<="
	GTEQ = ">="
	LT   = "<"
	GT   = ">"
	// Delimiters
	COMMA     = ","
	PERIOD    = "."
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	MUT      = "MUT"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	STRUCT   = "STRUCT"
)

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"let":    LET,
	"mut":    MUT,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"struct": STRUCT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
