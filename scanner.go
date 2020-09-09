package lox

import (
	"fmt"
)

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func (s *Scanner) ScanTokens() []Token {
	for s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, Token{TypeOf: EOF, Lexeme: "", Literal: nil, Line: s.line})

	return s.tokens
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
		break
	default:
		// lox.Lox.Error(s.line, "Unexpected character.")
		fmt.Println("Todo figure out error reporting here")
		panic("Unexpected character")
	}
}

func (s Scanner) isAtEnd() bool {
	return s.current > len(s.source)
}

func (s *Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) addToken(typeOf TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{TypeOf: typeOf, Lexeme: text, Literal: literal, Line: s.line})

}
