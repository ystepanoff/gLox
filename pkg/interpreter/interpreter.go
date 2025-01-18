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
	case scanner.GREATER:
		return left.(float64) > right.(float64)
	case scanner.GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case scanner.LESS:
		return left.(float64) < right.(float64)
	case scanner.LESS_EQUAL:
		return left.(float64) <= right.(float64)
	case scanner.STAR:
		return left.(float64) * right.(float64)
	case scanner.SLASH:
		return left.(float64) / right.(float64)
	case scanner.EQUAL_EQUAL:
		if left == nil && right == nil {
			return true
		}
		if left == nil {
			return false
		}
		return left == right
	case scanner.BANG_EQUAL:
		if left == nil && right == nil {
			return false
		}
		if left == nil {
			return true
		}
		return left != right
	case scanner.PLUS:
		if l, okL := left.(float64); okL {
			if r, okR := right.(float64); okR {
				return l + r
			}
		}
		if l, okL := left.(string); okL {
			if r, okR := right.(string); okR {
				return l + r
			}
		}
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
