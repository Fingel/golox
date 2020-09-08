package token

type Token struct {
	TypeOf  TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (token Token) String() string {
	return string(token.TypeOf) + " " + string(token.Lexeme)
}
