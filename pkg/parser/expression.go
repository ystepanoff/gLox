package parser

import "github.com/codecrafters-io/interpreter-starter-go/pkg/scanner"

type Expression interface {
	Accept(visitor ExpressionVisitor) interface{}
}

type Binary struct {
	Left     Expression
	Operator scanner.Token
	Right    Expression
}

type Grouping struct {
	Expression Expression
}

type Literal struct {
	Value interface{}
}

type Unary struct {
	Operator scanner.Token
	Right    Expression
}
