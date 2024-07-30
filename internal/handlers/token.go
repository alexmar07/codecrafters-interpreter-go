package handlers

const (
	// Single tokens
	LEFT_PAREN  = "LEFT_PAREN"
	RIGHT_PAREN = "RIGHT_PAREN"
	LEFT_BRACE  = "LEFT_BRACE"
	RIGHT_BRACE = "RIGHT_BRACE"
	COMMA       = "COMMA"
	DOT         = "DOT"
	MINUS       = "MINUS"
	PLUS        = "PLUS"
	SEMICOLON   = "SEMICOLON"
	SLASH       = "SLASH"
	STAR        = "STAR"

	// One or two tokens
	BANG          = "BANG"
	BANG_EQUAL    = "BANG_EQUAL"
	EQUAL         = "EQUAL"
	EQUAL_EQUAL   = "EQUAL_EQUAL"
	GREATER       = "GREATER"
	GREATER_EQUAL = "GREATER_EQUAL"
	LESS          = "LESS"
	LESS_EQUAL    = "LESS_EQUAL"

	// Literals
	IDENTIFIER = ""
	STRING     = "STRING"
	NUMBER     = "NUMBER"

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

	EOF = "EOF"
)

type Token struct {
	tokenType *TokenType
	lexeme    string
	literal   interface{}
	line      *int
}

func NewToken(tt *TokenType, lexeme string, literal interface{}, line *int) *Token {
	return &Token{
		tokenType: tt,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (t *Token) ToString() string {

	if t.literal != nil {

		str, _ := t.literal.(string)

		return t.tokenType.GetValue() + " " + t.lexeme + " " + str

	}

	return t.tokenType.GetValue() + " " + t.lexeme + " null"
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

func createSimpleToken(tokenType string, lexeme string) *Token {
	return NewToken(NewTokenType(tokenType), lexeme, nil, nil)
}

func createTokenWithValue(tokenType string, lexeme string, literal interface{}) *Token {
	return NewToken(NewTokenType(tokenType), lexeme, literal, nil)
}
