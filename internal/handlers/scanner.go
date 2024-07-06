package handlers

type Scanner struct {
	source  string
	tokens  []*Token
	current int
	start   int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		current: 0,
		start:   0,
		line:    1,
	}
}

func (s *Scanner) scanTokens() []*Token {

	for !s.isEndAt() {
		s.start = s.current

	}

	s.addToken(NewToken(NewTokenType(EOF), nil, nil, &s.line))

	return s.tokens
}

func (s *Scanner) scanToken() {

	c := s.source[s.current]

	s.current++

	switch string(c) {
	case "(":
		s.addToken(createSimpleToken(LEFT_PAREN))
	case ")":
		s.addToken(createSimpleToken(RIGHT_PAREN))
	case "{":
		s.addToken(createSimpleToken(LEFT_BRACE))
	case "}":
		s.addToken(createSimpleToken(RIGHT_BRACE))
	case ",":
		s.addToken(createSimpleToken(COMMA))
	case ".":
		s.addToken(createSimpleToken(DOT))
	case "-":
		s.addToken(createSimpleToken(MINUS))
	case "+":
		s.addToken(createSimpleToken(PLUS))
	case ";":
		s.addToken(createSimpleToken(SEMICOLON))
	case "/":
		s.addToken(createSimpleToken(SLASH))
	case "*":
		s.addToken(createSimpleToken(STAR))
	case "!":
		s.addTokenWithTwoChars(BANG, BANG_EQUAL)
	case "=":
		s.addTokenWithTwoChars(EQUAL, EQUAL_EQUAL)
	case "<":
		s.addTokenWithTwoChars(LESS, LESS_EQUAL)
	case ">":
		s.addTokenWithTwoChars(GREATER, GREATER_EQUAL)
	}

}

func (s *Scanner) addTokenWithTwoChars(oneCharOption string, twoCharsOption string) {

	if s.checkNextChar("=") {
		s.addToken(createSimpleToken(twoCharsOption))
	} else {
		s.addToken(createSimpleToken(oneCharOption))
	}

}

func (s *Scanner) checkNextChar(c string) bool {

	if s.isEndAt() {
		return false
	}

	if string(s.source[s.current]) != c {
		return false
	}

	s.current++

	return true
}

func createSimpleToken(tokenType string) *Token {
	return NewToken(NewTokenType(tokenType), nil, nil, nil)
}

func (s *Scanner) addToken(t *Token) {
	s.tokens = append(s.tokens, t)
}

func (s Scanner) isEndAt() bool {
	return s.current >= len(s.source)
}
