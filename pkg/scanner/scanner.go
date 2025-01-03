package scanner

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
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
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, &Token{
		TokenType: EOF,
		Lexeme:    "",
		Literal:   nil,
		Line:      s.line,
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
			s.advance()
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '<':
		if s.peek() == '=' {
			s.advance()
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.peek() == '=' {
			s.advance()
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.peek() == '/' {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ', '\r', '\t':
		// Ignore
	case '\n':
		s.line++
	case '"':
		s.scanString()
	default:
		if unicode.IsDigit(c) {
			s.scanNumber()
		} else if unicode.IsLetter(c) || c == '_' {
			s.scanIdentifier()
		} else {
			s.reportError(s.line, fmt.Sprintf("Unexpected character: %c", c))
		}
	}
}

func (s *Scanner) advance() rune {
	s.current++
	return rune(s.source[s.current-1])
}

func (s *Scanner) peek(args ...int) rune {
	step := 0
	if len(args) > 0 {
		step = args[0]
	}
	if s.current+step >= len(s.source) {
		return '\000'
	}
	return rune(s.source[s.current+step])
}

func (s *Scanner) scanString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		s.reportError(s.line, "Unterminated string.")
		return
	}
	s.advance()
	s.addTokenLiteral(STRING, s.source[s.start+1:s.current-1])
}

func (s *Scanner) scanNumber() {
	for unicode.IsDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && unicode.IsDigit(s.peek(1)) {
		s.advance()
		for unicode.IsDigit(s.peek()) {
			s.advance()
		}
	}
	value, _ := strconv.ParseFloat(s.source[s.start:s.current], 64)
	s.addTokenLiteral(NUMBER, value)
}

func (s *Scanner) scanIdentifier() {
	for unicode.IsLetter(s.peek()) || unicode.IsDigit(s.peek()) || s.peek() == '_' {
		s.advance()
	}
	s.addToken(IDENTIFIER)
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenLiteral(tokenType, nil)
}

func (s *Scanner) addTokenLiteral(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	if keywordType, exists := keywords[text]; tokenType == IDENTIFIER && exists {
		tokenType = keywordType
	}
	s.tokens = append(s.tokens, &Token{
		TokenType: tokenType,
		Lexeme:    text,
		Literal:   literal,
		Line:      s.line,
	})
}

func (s *Scanner) reportError(line int, message string) {
	s.hadErrors = true
	fmt.Fprintf(os.Stderr, "[Line %d] Error: %s\n", line, message)
}
