package parser

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/internal/scanner"
)

// Parser
type Parser struct {
	tokens  []scanner.Token
	current int

	parsedExpression Expression

	hadErrors bool
}

func NewParser(tokens []scanner.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) peek() scanner.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() scanner.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == scanner.EOF
}

func (p *Parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) match(tokenTypes ...scanner.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if !p.isAtEnd() && p.peek().TokenType == tokenType {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) consume(
	tokenType scanner.TokenType,
	message string,
) {
	if p.isAtEnd() || p.peek().TokenType != tokenType {
		p.reportError(p.peek(), message)
	}
	p.advance()
}

func (p *Parser) reportError(token scanner.Token, message string) {
	if token.TokenType == scanner.EOF {
		fmt.Fprintf(
			os.Stderr,
			"[line %d] Error at end: %s\n",
			token.Line,
			message,
		)
	} else {
		fmt.Fprintf(os.Stderr, "[line %d] Error at '%s': %s\n", token.Line, token.Lexeme, message)
	}
	p.hadErrors = true
}

func (p *Parser) HadErrors() bool {
	return p.hadErrors
}

func (p *Parser) Parse() {
	p.parsedExpression = p.expression()
}

func (p *Parser) GetParsedExpression() Expression {
	return p.parsedExpression
}

// Grammar definition
func (p *Parser) expression() Expression {
	return p.equality()
}

func (p *Parser) equality() Expression {
	equality := p.comparison()
	for p.match(scanner.BANG_EQUAL, scanner.EQUAL_EQUAL) {
		equality = &Binary{
			Left:     equality,
			Operator: p.previous(),
			Right:    p.comparison(),
		}
	}
	return equality
}

func (p *Parser) comparison() Expression {
	comparison := p.term()
	for p.match(scanner.GREATER, scanner.GREATER_EQUAL, scanner.LESS, scanner.LESS_EQUAL) {
		comparison = &Binary{
			Left:     comparison,
			Operator: p.previous(),
			Right:    p.term(),
		}
	}
	return comparison
}

func (p *Parser) term() Expression {
	term := p.factor()
	for p.match(scanner.MINUS, scanner.PLUS) {
		term = &Binary{
			Left:     term,
			Operator: p.previous(),
			Right:    p.factor(),
		}
	}
	return term
}

func (p *Parser) factor() Expression {
	factor := p.unary()
	for p.match(scanner.STAR, scanner.SLASH) {
		factor = &Binary{
			Left:     factor,
			Operator: p.previous(),
			Right:    p.unary(),
		}
	}
	return factor
}

func (p *Parser) unary() Expression {
	if p.match(scanner.PLUS, scanner.MINUS, scanner.BANG) {
		return &Unary{
			Operator: p.previous(),
			Right:    p.unary(),
		}
	}
	return p.primary()
}

func (p *Parser) primary() Expression {
	if p.match(scanner.TRUE) {
		return &Literal{Value: true}
	}
	if p.match(scanner.FALSE) {
		return &Literal{Value: false}
	}
	if p.match(scanner.NIL) {
		return &Literal{Value: nil}
	}
	if p.match(scanner.NUMBER, scanner.STRING) {
		return &Literal{Value: p.previous().Literal}
	}
	if p.match(scanner.LEFT_PAREN) {
		expression := p.expression()
		p.consume(scanner.RIGHT_PAREN, "Expect ')' after expression.")
		return &Grouping{Expression: expression}
	}
	p.reportError(p.peek(), "Expect expression.")
	return nil
}
