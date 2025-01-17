package scanner

import (
	"fmt"
	"math"
	"strconv"
)

// TokenType
type TokenType int

const (
	LEFT_PAREN TokenType = iota
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

	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	IDENTIFIER
	STRING
	NUMBER

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

var tokenTypes = []string{
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"COMMA",
	"DOT",
	"MINUS",
	"PLUS",
	"SEMICOLON",
	"SLASH",
	"STAR",

	"BANG",
	"BANG_EQUAL",
	"EQUAL",
	"EQUAL_EQUAL",
	"GREATER",
	"GREATER_EQUAL",
	"LESS",
	"LESS_EQUAL",

	"IDENTIFIER",
	"STRING",
	"NUMBER",

	"AND",
	"CLASS",
	"ELSE",
	"FALSE",
	"FUN",
	"FOR",
	"IF",
	"NIL",
	"OR",
	"PRINT",
	"RETURN",
	"SUPER",
	"THIS",
	"TRUE",
	"VAR",
	"WHILE",

	"EOF",
}

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func (t TokenType) String() string {
	return tokenTypes[t]
}

// Token
type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   interface{}
	Line      int
}

func (t *Token) String() string {
	literal := t.Literal
	if literal == nil {
		literal = "null"
	} else if t.TokenType == NUMBER {
		_, fractionalPart := math.Modf(literal.(float64))
		if fractionalPart == 0 {
			literal = fmt.Sprintf("%.1f", literal.(float64))
		} else {
			literal = strconv.FormatFloat(literal.(float64), 'f', -1, 64)
		}
	}
	return fmt.Sprintf("%s %s %v", t.TokenType, t.Lexeme, literal)
}
