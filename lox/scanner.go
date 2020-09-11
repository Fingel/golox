package lox

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	Source  string
	tokens  []Token
	start   int
	current int
	line    int
	errors  []ScanError
}

type ScanError struct {
	Message 	string
	Line 		int
	Where 		string
}

func (e *ScanError) Error() string {
	return fmt.Sprintf("Scan error line %d at %s: %s", e.Line, e.Where, e.Message)
}

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func NewScanner(source string) Scanner {
	scanner := Scanner{Source: source}
	scanner.line = 1
	return scanner
}

func (s *Scanner) ScanTokens() ([]Token, []ScanError) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{TypeOf: EOF, Lexeme: "", Literal: nil, Line: s.line})

	return s.tokens, s.errors
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN, nil)
		break
	case ')':
		s.addToken(RIGHT_PAREN, nil)
		break
	case '{':
		s.addToken(LEFT_BRACE, nil)
		break
	case '}':
		s.addToken(RIGHT_BRACE, nil)
		break
	case ',':
		s.addToken(COMMA, nil)
		break
	case '.':
		s.addToken(DOT, nil)
		break
	case '-':
		s.addToken(MINUS, nil)
		break
	case '+':
		s.addToken(PLUS, nil)
		break
	case ';':
		s.addToken(SEMICOLON, nil)
		break
	case '*':
		s.addToken(STAR, nil)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
		break
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}
		break
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}
		break
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}
		break
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}
		break
	case '"':
		s.getString()
		break
	case ' ':
		break
	case '\r':
		break
	case '\t':
		break
	case '\n':
		s.line++
		break
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			fmt.Println("Appending new error")
			s.errors = append(s.errors, ScanError{"Unexpected character", s.line, string(c)})
		}
	}
}

func (s *Scanner) identifier() {
	for s.isAlphaNumberic(s.peek()) {
		s.advance()
	}
	text := s.Source[s.start:s.current]
	typeOf, ok := keywords[text]
	if ok == false {
		typeOf = INDENTIFIER
	}
	s.addToken(typeOf, nil)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	value, err := strconv.ParseFloat(s.Source[s.start:s.current], 64)
	if err != nil {
		s.errors = append(s.errors, ScanError{"Malformed number", s.line, s.Source[s.start:s.current]})
		return
	}
	s.addToken(NUMBER, value)
}

func (s *Scanner) getString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.errors = append(s.errors, ScanError{"Unterminated line", s.line, ""})
		return
	}

	s.advance()

	value := s.Source[s.start+1 : s.current-1]
	s.addToken(STRING, value)
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) advance() byte {
	s.current++
	return s.Source[s.current-1]
}

func (s *Scanner) addToken(typeOf TokenType, literal interface{}) {
	text := s.Source[s.start:s.current]
	s.tokens = append(s.tokens, Token{TypeOf: typeOf, Lexeme: text, Literal: literal, Line: s.line})
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.Source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s Scanner) peek() byte {
	if s.isAtEnd() {
		return '\x00'
	}
	return s.Source[s.current]
}

func (s Scanner) peekNext() byte {
	if s.current+1 >= len(s.Source) {
		return '\x00'
	}
	return s.Source[s.current+1]
}

func (s Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (s Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s Scanner) isAlphaNumberic(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
}
