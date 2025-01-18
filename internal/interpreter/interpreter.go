package interpreter

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/internal/parser"
	"github.com/codecrafters-io/interpreter-starter-go/internal/scanner"
)

type Interpreter struct {
	hadErrors bool
}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) reportError(operator scanner.Token, message string) {
	fmt.Fprintf(os.Stderr, "%s", message)
}

func (i *Interpreter) HadErrors() bool {
	return i.hadErrors
}

func (i *Interpreter) VisitBinary(binary *parser.Binary) interface{} {
	if i.hadErrors {
		return nil
	}
	left := binary.Left.Accept(i)
	right := binary.Right.Accept(i)
	switch binary.Operator.TokenType {
	case scanner.GREATER:
		if !checkValuesType[float64](left, right) {
			i.reportError(binary.Operator, "Operands must be numbers.")
			i.hadErrors = true
			return nil
		}
		return left.(float64) > right.(float64)
	case scanner.GREATER_EQUAL:
		if !checkValuesType[float64](left, right) {
			i.reportError(binary.Operator, "Operands must be numbers.")
			i.hadErrors = true
			return nil
		}
		return left.(float64) >= right.(float64)
	case scanner.LESS:
		if !checkValuesType[float64](left, right) {
			i.reportError(binary.Operator, "Operands must be numbers.")
			i.hadErrors = true
			return nil
		}
		return left.(float64) < right.(float64)
	case scanner.LESS_EQUAL:
		if !checkValuesType[float64](left, right) {
			i.reportError(binary.Operator, "Operands must be numbers.")
			i.hadErrors = true
			return nil
		}
		return left.(float64) <= right.(float64)
	case scanner.STAR:
		if !checkValuesType[float64](left, right) {
			i.reportError(binary.Operator, "Operands must be numbers.")
			i.hadErrors = true
			return nil
		}
		return left.(float64) * right.(float64)
	case scanner.SLASH:
		if !checkValuesType[float64](left, right) {
			i.reportError(binary.Operator, "Operands must be numbers.")
			i.hadErrors = true
			return nil
		}
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
		i.reportError(
			binary.Operator,
			"Operands must be two numbers or two strings.",
		)
		i.hadErrors = true
		return nil
	case scanner.MINUS:
		if !checkValuesType[float64](left, right) {
			i.reportError(binary.Operator, "Operands must be numbers.")
			i.hadErrors = true
			return nil
		}
		return left.(float64) - right.(float64)
	}
	return nil
}

func (i *Interpreter) VisitGrouping(grouping *parser.Grouping) interface{} {
	if i.hadErrors {
		return nil
	}
	return grouping.Expression.Accept(i)
}

func (i *Interpreter) VisitLiteral(literal *parser.Literal) interface{} {
	if i.hadErrors {
		return nil
	}
	return literal.Value
}

func (i *Interpreter) VisitUnary(unary *parser.Unary) interface{} {
	if i.hadErrors {
		return nil
	}
	right := unary.Right.Accept(i)
	switch unary.Operator.TokenType {
	case scanner.MINUS:
		if !checkValueType[float64](right) {
			i.reportError(unary.Operator, "Operand must be a number.")
			i.hadErrors = true
			return nil
		}
		return -right.(float64)
	case scanner.BANG:
		if right == nil {
			return true
		}
		if v, ok := right.(bool); ok {
			return !v
		}
		return false
	}
	return right
}

func (i *Interpreter) Interpret(expression parser.Expression) {
	value := expression.Accept(i)
	if i.hadErrors {
		return
	}
	if value == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(value)
	}
}
