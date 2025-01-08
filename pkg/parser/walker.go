package parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Walker interface {
	process(name string, expressions ...Expression) interface{}
}

// BaseWalker
type BaseWalker struct {
	walker Walker
}

func (bw *BaseWalker) VisitBinary(binary *Binary) interface{} {
	return bw.walker.process(binary.Operator.Lexeme, binary.Left, binary.Right)
}

func (bw *BaseWalker) VisitGrouping(grouping *Grouping) interface{} {
	return bw.walker.process("group", grouping.Expression)
}

func (bw *BaseWalker) VisitLiteral(literal *Literal) interface{} {
	if literal.Value == nil {
		return "nil"
	}
	if value, ok := literal.Value.(float64); ok {
		_, fractionalPart := math.Modf(value)
		if fractionalPart == 0 {
			return fmt.Sprintf("%.1f", value)
		} else {
			return strconv.FormatFloat(value, 'f', -1, 64)
		}
	}
	return literal.Value
}

func (bw *BaseWalker) VisitUnary(unary *Unary) interface{} {
	return bw.walker.process(unary.Operator.Lexeme, unary.Right)
}

func (walker *BaseWalker) process(
	name string,
	expressions ...Expression,
) interface{} {
	fmt.Println("walker.process() should be implemented by specific types")
	return nil
}

// ASTPrinter
type ASTPrinter struct {
	BaseWalker
}

func NewASTPrinter() *ASTPrinter {
	printer := &ASTPrinter{}
	printer.walker = printer
	return printer
}

func (printer *ASTPrinter) process(
	name string,
	expressions ...Expression,
) interface{} {
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

func (printer *ASTPrinter) Print(expression Expression) interface{} {
	return expression.Accept(printer)
}

// RPNPrinter
type RPNPrinter struct {
	BaseWalker
}

func NewRPNPrinter() *RPNPrinter {
	printer := &RPNPrinter{}
	printer.walker = printer
	return printer
}

func (printer *RPNPrinter) process(
	name string,
	expressions ...Expression,
) interface{} {
	var builder strings.Builder
	for _, expression := range expressions {
		builder.WriteString(
			fmt.Sprintf("%v ", expression.Accept(printer)),
		)
	}
	builder.WriteString(name)
	return builder.String()
}

func (printer *RPNPrinter) Print(expression Expression) string {
	value := expression.Accept(printer)
	if value == nil {
		return "nil"
	}
	return value.(string)
}
