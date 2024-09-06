package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	//operators
	PLUS  = "+"
	MINUS = "-"
	EQUAL = "="

	//Delimitors
	SEMICOLON = ";"
	COLON     = ":"
	COMMA     = ","
	LPARAN    = "("
	RPARAN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
