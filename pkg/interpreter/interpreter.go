package interpreter

import (
	"fmt"

	"github.com/codecrafters-io/interpreter-starter-go/pkg/parser"
)

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) VisitBinary(binary *parser.Binary) interface{} {
	return nil
}

func (i *Interpreter) VisitGrouping(grouping *parser.Grouping) interface{} {
	return grouping.Expression.Accept(i)
}

func (i *Interpreter) VisitLiteral(literal *parser.Literal) interface{} {
	return literal.Value
}

func (i *Interpreter) VisitUnary(unary *parser.Unary) interface{} {
	return nil
}

func (i *Interpreter) Interpret(expression parser.Expression) {
	value := expression.Accept(i)
	if value == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(value)
	}
}