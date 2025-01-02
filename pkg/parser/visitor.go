package parser

type ExpressionVisitor interface {
	VisitBinary(binary *Binary) interface{}
	VisitGrouping(grouping *Grouping) interface{}
	VisitLiteral(literal *Literal) interface{}
	VisitUnary(unary *Unary) interface{}
}

func (binary *Binary) Accept(visitor ExpressionVisitor) interface{} {
	return visitor.VisitBinary(binary)
}

func (grouping *Grouping) Accept(visitor ExpressionVisitor) interface{} {
	return visitor.VisitGrouping(grouping)
}

func (literal *Literal) Accept(visitor ExpressionVisitor) interface{} {
	return visitor.VisitLiteral(literal)
}

func (unary *Unary) Accept(visitor ExpressionVisitor) interface{} {
	return visitor.VisitUnary(unary)
}
