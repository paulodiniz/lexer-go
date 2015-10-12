package lexer

type TokenType int

const (
	// Special tokens
	ILLEGAL = iota
	EOF
	WS
	// Literals
	IDENT // fields, table_name
	// Misc characters
	ASTERISK // *
	COMMA    // ,
	// Keywords
	SELECT
	FROM
)

type Token struct {
  Type  TokenType
  Value string
}
