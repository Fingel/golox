package scanner

import (
	"fmt"

	"github.com/Fingel/golox/token"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

func (scanner *Scanner) ScanTokens() []token.Token {
	for scanner.isAtEnd() {
		scanner.start = scanner.current
		scanner.scanToken()
	}
	scanner.tokens = append(scanner.tokens, token.Token{TypeOf: token.EOF, Lexeme: "", Literal: nil, Line: scanner.line})

	return scanner.tokens
}

func (scanner *Scanner) scanToken() {
	c := scanner.advance()
	switch c {
	case '(':
		scanner.addToken(token.LEFT_PAREN, nil)
		break
	case ')':
		scanner.addToken(token.RIGHT_PAREN, nil)
		break
	case '{':
		scanner.addToken(token.LEFT_BRACE, nil)
		break
	case '}':
		scanner.addToken(token.RIGHT_BRACE, nil)
		break
	case ',':
		scanner.addToken(token.COMMA, nil)
		break
	case '.':
		scanner.addToken(token.DOT, nil)
		break
	case '-':
		scanner.addToken(token.MINUS, nil)
		break
	case '+':
		scanner.addToken(token.PLUS, nil)
		break
	case ';':
		scanner.addToken(token.SEMICOLON, nil)
		break
	case '*':
		scanner.addToken(token.STAR, nil)
		break
	default:
		// lox.Lox.Error(scanner.line, "Unexpected character.")
		fmt.Println("Todo figure out error reporting here")
		panic("Unexpected character")
	}
}

func (scanner Scanner) isAtEnd() bool {
	return scanner.current > len(scanner.source)
}

func (scanner *Scanner) advance() byte {
	scanner.current++
	return scanner.source[scanner.current-1]
}

func (scanner *Scanner) addToken(typeOf token.TokenType, literal interface{}) {
	text := scanner.source[scanner.start:scanner.current]
	scanner.tokens = append(scanner.tokens, token.Token{TypeOf: typeOf, Lexeme: text, Literal: literal, Line: scanner.line})

}
