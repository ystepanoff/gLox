package parser

// ExpressionVisitor
type ExpressionVisitor interface {
	VisitBinary(binary *Binary) interface{}
	VisitGrouping(grouping *Grouping) interface{}
	VisitLiteral(literal *Literal) interface{}
	VisitUnary(unary *Unary) interface{}
}

func (b *Binary) Accept(visitor ExpressionVisitor) interface{} {
	return visitor.VisitBinary(b)
}

func (g *Grouping) Accept(visitor ExpressionVisitor) interface{} {
	return visitor.VisitGrouping(g)
}

func (l *Literal) Accept(visitor ExpressionVisitor) interface{} {
	return visitor.VisitLiteral(l)
}

func (u *Unary) Accept(visitor ExpressionVisitor) interface{} {
	return visitor.VisitUnary(u)
}
