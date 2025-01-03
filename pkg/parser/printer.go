package parser

import (
	"fmt"
	"strings"
)

type ASTPrinter struct{}

func NewASTPrinter() *ASTPrinter {
	return &ASTPrinter{}
}

func (printer *ASTPrinter) VisitBinary(binary *Binary) interface{} {
	return printer.process(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (printer *ASTPrinter) VisitGrouping(grouping *Grouping) interface{} {
	return printer.process("group", grouping.Expression)
}

func (printer *ASTPrinter) VisitLiteral(literal *Literal) interface{} {
	if literal.Value == nil {
		return "nil"
	}
	return literal.Value
}

func (printer *ASTPrinter) VisitUnary(unary *Unary) interface{} {
	return printer.process(unary.Operator.Lexeme, unary.Right)
}

func (printer *ASTPrinter) process(name string, expressions ...Expression) string {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)
	for _, expression := range expressions {
		builder.WriteString(" ")
		builder.WriteString(fmt.Sprintf("%v", expression.Accept(printer)))
	}
	builder.WriteString(")")
	return builder.String()
}

func (printer *ASTPrinter) Print(expression Expression) string {
	return expression.Accept(printer).(string)
}

type RPNPrinter struct{}

func NewRPNPrinter() *RPNPrinter {
	return &RPNPrinter{}
}

func (printer *RPNPrinter) VisitBinary(binary *Binary) interface{} {
	return printer.process(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (printer *RPNPrinter) VisitGrouping(grouping *Grouping) interface{} {
	return grouping.Expression.Accept(printer)
}

func (printer *RPNPrinter) VisitLiteral(literal *Literal) interface{} {
	return literal.Value
}

func (printer *RPNPrinter) VisitUnary(unary *Unary) interface{} {
	return printer.process(unary.Operator.Lexeme, unary.Right)
}

func (printer *RPNPrinter) process(name string, expressions ...Expression) string {
	var builder strings.Builder
	for _, expression := range expressions {
		builder.WriteString(fmt.Sprintf("%v ", expression.Accept(printer)))
	}
	builder.WriteString(name)
	return builder.String()
}

func (printer *RPNPrinter) Print(expression Expression) string {
	return expression.Accept(printer).(string)
}
