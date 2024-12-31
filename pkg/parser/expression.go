package parser

import "github.com/codecrafters-io/interpreter-starter-go/pkg/scanner"

type Expression struct {
}

type Binary struct {
	Expression
	Left     Expression
	Operator scanner.Token
	Right    Expression
}

func NewBinary(left Expression, operator scanner.Token, right Expression) *Binary {
	return &Binary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}
