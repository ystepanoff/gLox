package parser

import "github.com/codecrafters-io/interpreter-starter-go/pkg/scanner"

type Expression struct {
	Accept func(ExpressionVisitor) interface{}
}

type Binary struct {
	Left     *Expression
	Operator scanner.Token
	Right    *Expression
}

func NewBinary(left *Expression, operator scanner.Token, right *Expression) *Binary {
	return &Binary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

type Grouping struct {
	Expression *Expression
}

func NewGrouping(expression *Expression) *Grouping {
	return &Grouping{
		Expression: expression,
	}
}

type Literal struct {
	Value *interface{}
}

func NewLiteral(value *interface{}) *Literal {
	return &Literal{
		Value: value,
	}
}

type Unary struct {
	Operator scanner.Token
	Right    *Expression
}

func NewUnary(operator scanner.Token, right *Expression) *Unary {
	return &Unary{
		Operator: operator,
		Right:    right,
	}
}
