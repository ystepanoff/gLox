package scanner

import (
	"fmt"
	"os"
)

type Scanner struct {
	source    string
	tokens    []*Token
	hadErrors bool

	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		line:   1,
	}
}

func (s *Scanner) ScanTokens() {
	for !s.isAtEnd() && s.peek() != '\n' {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, &Token{
		tokenType: EOF,
		lexeme:    "",
		literal:   nil,
		line:      s.line,
	})
}

func (s *Scanner) GetTokens() []*Token {
	return s.tokens
}

func (s *Scanner) HadErrors() bool {
	return s.hadErrors
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '=':
		if s.peek() == '=' {
			s.current++
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '!':
		if s.peek() == '=' {
			s.current++
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	default:
		s.reportError(s.line, fmt.Sprintf("Unexpected character: %c", c))
	}
}

func (s *Scanner) advance() rune {
	s.current++
	return rune(s.source[s.current-1])
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}
	return rune(s.source[s.current])
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenLiteral(tokenType, nil)
}

func (s *Scanner) addTokenLiteral(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}

func (s *Scanner) reportError(line int, message string) {
	s.hadErrors = true
	fmt.Fprintf(os.Stderr, "[line %d] Error: %s\n", line, message)
}
