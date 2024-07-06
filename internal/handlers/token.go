package handlers

const (
	// Single tokens
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"
	COMMA       = ","
	DOT         = "."
	MINUS       = "-"
	PLUS        = "+"
	SEMICOLON   = ";"
	SLASH       = "/"
	STAR        = "*"

	// One or two tokens
	BANG          = "!"
	BANG_EQUAL    = "!="
	EQUAL         = "="
	EQUAL_EQUAL   = "=="
	GREATER       = ">"
	GREATER_EQUAL = ">="
	LESS          = "<"
	LESS_EQUAL    = "<="

	// Literals
	IDENTIFIER = ""
	STRING     = ""
	NUMBER     = ""

	// Keywords
	AND    = ""
	CLASS  = ""
	ELSE   = ""
	FALSE  = ""
	FUN    = ""
	FOR    = ""
	IF     = ""
	NIL    = ""
	OR     = ""
	PRINT  = ""
	RETURN = ""
	SUPER  = ""
	THIS   = ""
	TRUE   = ""
	VAR    = ""
	WHILE  = ""

	EOF = ""
)

type Token struct {
	tokenType *TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func NewToken(tt *TokenType, lexeme *string, literal interface{}, line *int) *Token {
	return &Token{
		tokenType: tt,
		lexeme:    *lexeme,
		literal:   literal,
		line:      *line,
	}
}

func (t *Token) ToString() string {
	return ""
}

type TokenType struct {
	value string
}

func NewTokenType(value string) *TokenType {
	return &TokenType{value: value}
}

func (tt *TokenType) GetValue() string {
	return tt.value
}
