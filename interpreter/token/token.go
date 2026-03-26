package token

const (
	ILLEGAL = "ILLEGAL" // token we don't know about
	EOF     = "EOF"     // end of file

	// identifier + literals
	IDENT = "IDENT" // foobar , add , b ,x
	INT   = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType //EQUAL_SIGN , IDENTIFIER , SEMICOLON
	Literal string    // stores the value of the token
}

// source code will be a string for simplicity but in production we should be using io.Reader to read the file and attach file name , line numbers in the lexer output
