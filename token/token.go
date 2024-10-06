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
	EQ    = "=="
	NOTEQ = "!="

	//Delimitors
	SEMICOLON = ";"
	COLON     = ":"
	COMMA     = ","
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	BANG      = "!"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
