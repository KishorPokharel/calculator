package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL    TokenType = "ILLEGAL"
	EOF        TokenType = "EOF"
	PLUS       TokenType = "+"
	SUBTRACT   TokenType = "-"
	DIVIDE     TokenType = "/"
	MULTIPLY   TokenType = "*"
	LPAREN     TokenType = "("
	RPAREN     TokenType = ")"
	BAR        TokenType = "|"
	NUMBER     TokenType = "NUMBER"
	ASSIGN     TokenType = "="
	IDENTIFIER TokenType = "IDENTIFIER"
)
