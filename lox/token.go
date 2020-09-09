package lox

type Token struct {
	TypeOf  TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (token Token) String() string {
	return "TOKEN> typeOf: " + token.TypeOf.String() + " lexeme: " + token.Lexeme
}
