package handlers

import (
	"strconv"
	"strings"
	"unicode"
)

type Scanner struct {
	source    string
	tokens    []*Token
	current   int
	start     int
	line      int
	HasErrors bool
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:    source,
		current:   0,
		start:     0,
		line:      1,
		HasErrors: false,
	}
}

func (s *Scanner) ScanTokens() []*Token {

	for !s.isEndAt() {
		s.start = s.current
		s.scanToken()
	}

	s.addToken(NewToken(NewTokenType(EOF), "", nil, &s.line))

	return s.tokens
}

func (s *Scanner) scanToken() {

	c := s.source[s.current]

	s.current++

	switch string(c) {
	case "(":
		s.addToken(createSimpleToken(LEFT_PAREN, "("))
	case ")":
		s.addToken(createSimpleToken(RIGHT_PAREN, ")"))
	case "{":
		s.addToken(createSimpleToken(LEFT_BRACE, "{"))
	case "}":
		s.addToken(createSimpleToken(RIGHT_BRACE, "}"))
	case ",":
		s.addToken(createSimpleToken(COMMA, ","))
	case ".":
		s.addToken(createSimpleToken(DOT, "."))
	case "-":
		s.addToken(createSimpleToken(MINUS, "-"))
	case "+":
		s.addToken(createSimpleToken(PLUS, "+"))
	case ";":
		s.addToken(createSimpleToken(SEMICOLON, ";"))
	case "*":
		s.addToken(createSimpleToken(STAR, "*"))
	case "!":
		s.addTokenWithTwoChars("!", BANG, BANG_EQUAL)
	case "=":
		s.addTokenWithTwoChars("=", EQUAL, EQUAL_EQUAL)
	case "<":
		s.addTokenWithTwoChars("<", LESS, LESS_EQUAL)
	case ">":
		s.addTokenWithTwoChars(">", GREATER, GREATER_EQUAL)
	case "/":
		// Gestisce i commenti
		if s.checkNextChar("/") {
			for s.peek() != "\n" && !s.isEndAt() {
				s.current++
			}
		} else {
			s.addToken(createSimpleToken(SLASH, "/"))
		}
	case " ":
	case "\t":
	case "\r":
	case "\n":
		s.line++
	case "\"":
		s.detectString()
	default:

		if isDigit(string(c)) {
			s.detectNumber()
		} else {
			ErrorUnexpectedCharacter(s.line, string(c))
			s.HasErrors = true
		}

	}

}

func (s *Scanner) peek() string {

	if s.isEndAt() {
		return "\\0"
	}

	return string(s.source[s.current])

}

func (s *Scanner) detectString() {

	for s.peek() != "\"" && !s.isEndAt() {
		if s.peek() == "\n" {
			s.line++
		}

		s.current++
	}

	if s.isEndAt() {
		UnterminatedString(s.line)
		s.HasErrors = true
		return
	}

	s.current++

	str := s.source[s.start:s.current]
	value := s.source[s.start+1 : s.current-1]

	s.addToken(createTokenWithValue(STRING, str, value))
}

func (s *Scanner) detectNumber() {

	for isDigit(s.peek()) {
		s.current++
	}

	if s.peek() == "." && isDigit(s.peekNext()) {
		s.current++
	}

	for isDigit(s.peek()) {
		s.current++
	}

	num := s.source[s.start:s.current]

	var value string

	fp, _ := strconv.ParseFloat(num, 64)

	values := strings.Split(num, ".")

	if len(values) == 1 {
		value = strconv.FormatFloat(fp, 'f', 1, 64)
	} else {
		value = num
	}

	s.addToken(createTokenWithValue(NUMBER, num, value))
}

func (s *Scanner) peekNext() string {

	if s.current+1 >= len(s.source) {
		return "\\0"
	}

	return string(s.source[s.current+1])
}

func (s *Scanner) addTokenWithTwoChars(c string, oneCharOption string, twoCharsOption string) {

	if s.checkNextChar("=") {
		s.addToken(createSimpleToken(twoCharsOption, c+"="))
	} else {
		s.addToken(createSimpleToken(oneCharOption, c))
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

func (s *Scanner) addToken(t *Token) {
	s.tokens = append(s.tokens, t)
}

func (s Scanner) isEndAt() bool {
	return s.current >= len(s.source)
}

func isDigit(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
