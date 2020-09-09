package lox

type TokenType int

const (

	// Single-character tokens
	LEFT_PAREN TokenType = iota + 1
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	INDENTIFIER
	STRING
	NUMBER

	// KEYWORDS
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

func (t TokenType) String() string {
	switch t {
	case LEFT_PAREN:
		return "Left Paren"
	case RIGHT_PAREN:
		return "Right Paren"
	case LEFT_BRACE:
		return "Left Brace"
	case RIGHT_BRACE:
		return "Right Brace"
	case COMMA:
		return "Comma"
	case DOT:
		return "Dot"
	case MINUS:
		return "Minus"
	case PLUS:
		return "Plus"
	case SEMICOLON:
		return "Semicolon"
	case SLASH:
		return "Slash"
	case STAR:
		return "Star"
	case BANG:
		return "Bang"
	case BANG_EQUAL:
		return "Bang Equal"
	case EQUAL:
		return "Equal"
	case EQUAL_EQUAL:
		return "Equal Equal"
	case GREATER:
		return "Greater"
	case GREATER_EQUAL:
		return "Greater_equal"
	case LESS:
		return "Less"
	case LESS_EQUAL:
		return "Less_equal"
	case INDENTIFIER:
		return "Identifier"
	case STRING:
		return "String"
	case NUMBER:
		return "Number"
	case AND:
		return "And"
	case CLASS:
		return "Class"
	case ELSE:
		return "Else"
	case FALSE:
		return "False"
	case FUN:
		return "Fun"
	case FOR:
		return "For"
	case IF:
		return "If"
	case NIL:
		return "Nil"
	case OR:
		return "Or"
	case PRINT:
		return "Print"
	case RETURN:
		return "Return"
	case SUPER:
		return "Super"
	case THIS:
		return "This"
	case TRUE:
		return "True"
	case VAR:
		return "Var"
	case WHILE:
		return "While"
	case EOF:
		return "EOF"
	default:
		return "Unknown"
	}
}
