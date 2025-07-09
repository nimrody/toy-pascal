package token

// TokenType is a string that represents the type of a token.
type TokenType string

// Token represents a single lexical token.
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Constants for all token types.
const (
	// Special tokens
	ILLEGAL = "ILLEGAL" // A token/character we don't know about
	EOF     = "EOF"     // "End of File"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = ":="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	CARET    = "^"
	DOT      = "."

	LT     = "<"
	GT     = ">"
	EQ     = "="
	NOT_EQ = "<>"
	LTE    = "<="
	GTE    = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"

	// Keywords
	FUNCTION  = "FUNCTION"
	PROCEDURE = "PROCEDURE"
	TYPE      = "TYPE"
	VAR       = "VAR"
	RECORD    = "RECORD"
	IF        = "IF"
	ELSE      = "ELSE"
	THEN      = "THEN"
	WHILE     = "WHILE"
	DO        = "DO"
	BEGIN     = "BEGIN"
	END       = "END"
	NIL       = "NIL"
)

// keywords maps keyword strings to their corresponding TokenType.
var keywords = map[string]TokenType{
	"function":  FUNCTION,
	"procedure": PROCEDURE,
	"type":      TYPE,
	"var":       VAR,
	"record":    RECORD,
	"if":        IF,
	"else":      ELSE,
	"then":      THEN,
	"while":     WHILE,
	"do":        DO,
	"begin":     BEGIN,
	"end":       END,
	"nil":       NIL,
}

// LookupIdent checks if an identifier is a keyword.
// If it is, the keyword's TokenType is returned. Otherwise, IDENT is returned.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
