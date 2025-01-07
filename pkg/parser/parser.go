package parser

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/pkg/scanner"
)

// Parser
type Parser struct {
	tokens  []scanner.Token
	current int
}

func NewParser(tokens []scanner.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (parser *Parser) peek() scanner.Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) previous() scanner.Token {
	return parser.tokens[parser.current-1]
}

func (parser *Parser) isAtEnd() bool {
	return parser.peek().TokenType == scanner.EOF
}

func (parser *Parser) advance() scanner.Token {
	if !parser.isAtEnd() {
		parser.current++
	}
	return parser.previous()
}

func (parser *Parser) match(tokenTypes ...scanner.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if !parser.isAtEnd() && parser.peek().TokenType == tokenType {
			parser.advance()
			return true
		}
	}
	return false
}

func (parser *Parser) consume(
	tokenType scanner.TokenType,
	message string,
) (scanner.Token, error) {
	if !parser.isAtEnd() && parser.peek().TokenType == tokenType {
		return parser.advance(), nil
	}
	return parser.peek(), fmt.Errorf("%v: %s", parser.peek(), message)
}

func (parser *Parser) reportError(token scanner.Token, message string) {
	if token.TokenType == scanner.EOF {
		fmt.Fprintf(
			os.Stderr,
			"[Line %d] Error: at end %s",
			token.Line,
			message,
		)
	} else {
		fmt.Fprintf(os.Stderr, "[Line %d] Error: at '%s' %s", token.Line, token.Lexeme, message)
	}
}

func (parser *Parser) Parse() Expression {
	return parser.expression()
}

// Grammar definition
func (parser *Parser) expression() Expression {
	return parser.equality()
}

func (parser *Parser) equality() Expression {
	equality := parser.comparison()
	for parser.match(scanner.BANG_EQUAL, scanner.EQUAL_EQUAL) {
		equality = &Binary{
			Left:     equality,
			Operator: parser.previous(),
			Right:    parser.comparison(),
		}
	}
	return equality
}

func (parser *Parser) comparison() Expression {
	comparison := parser.term()
	for parser.match(scanner.GREATER, scanner.GREATER_EQUAL, scanner.LESS, scanner.LESS_EQUAL) {
		comparison = &Binary{
			Left:     comparison,
			Operator: parser.previous(),
			Right:    parser.term(),
		}
	}
	return comparison
}

func (parser *Parser) term() Expression {
	term := parser.factor()
	for parser.match(scanner.MINUS, scanner.PLUS) {
		term = &Binary{
			Left:     term,
			Operator: parser.previous(),
			Right:    parser.factor(),
		}
	}
	return term
}

func (parser *Parser) factor() Expression {
	factor := parser.unary()
	for parser.match(scanner.STAR, scanner.SLASH) {
		factor = &Binary{
			Left:     factor,
			Operator: parser.previous(),
			Right:    parser.unary(),
		}
	}
	return factor
}

func (parser *Parser) unary() Expression {
	if parser.match(scanner.PLUS, scanner.MINUS, scanner.BANG) {
		return &Unary{
			Operator: parser.previous(),
			Right:    parser.unary(),
		}
	}
	return parser.primary()
}

func (parser *Parser) primary() Expression {
	if parser.match(scanner.TRUE) {
		return &Literal{Value: true}
	}
	if parser.match(scanner.FALSE) {
		return &Literal{Value: false}
	}
	if parser.match(scanner.NIL) {
		return &Literal{Value: nil}
	}
	if parser.match(scanner.NUMBER, scanner.STRING) {
		return &Literal{Value: parser.previous().Literal}
	}
	if parser.match(scanner.LEFT_PAREN) {
		expression := parser.expression()
		parser.consume(scanner.RIGHT_PAREN, "Expect ')' after expression.")
		return &Grouping{Expression: expression}
	}
	return nil
}
