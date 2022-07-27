package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const(
	ILLEGAL		= "ILLEGAL"
	EOF			= "EOF"

	//식별자 + 리터럴
	INDENT		= "INDENT"
	INT			= "INT"

	//연산자
	ASSIGN		= "="
	PLUS		= "+"

	//구분자
	COMMA		= ","
	SEMICOLON	= ";"
	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"

	//예약어
	FUNCTION	= "FUNCTION"
	LET			= "LET"
)