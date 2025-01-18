package interpreter

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/pkg/parser"
	"github.com/codecrafters-io/interpreter-starter-go/pkg/scanner"
)

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) VisitBinary(binary *parser.Binary) interface{} {
	left := binary.Left.Accept(i)
	right := binary.Right.Accept(i)
	switch binary.Operator.TokenType {
	case scanner.STAR:
		return left.(float64) * right.(float64)
	case scanner.SLASH:
		return left.(float64) / right.(float64)
	case scanner.PLUS:
		return left.(float64) + right.(float64)
	case scanner.MINUS:
		return left.(float64) - right.(float64)
	}
	return nil
}

func (i *Interpreter) VisitGrouping(grouping *parser.Grouping) interface{} {
	return grouping.Expression.Accept(i)
}

func (i *Interpreter) VisitLiteral(literal *parser.Literal) interface{} {
	return literal.Value
}

func (i *Interpreter) VisitUnary(unary *parser.Unary) interface{} {
	value := unary.Right.Accept(i)
	switch unary.Operator.TokenType {
	case scanner.MINUS:
		return -value.(float64)
	case scanner.BANG:
		if value == nil {
			return true
		}
		if v, ok := value.(bool); ok {
			return !v
		}
		return false
	}
	return value
}

func (i *Interpreter) Interpret(expression parser.Expression) {
	value := expression.Accept(i)
	if value == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(value)
	}
}
