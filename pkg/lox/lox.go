package lox

import (
	"github.com/codecrafters-io/interpreter-starter-go/internal/interpreter"
	"github.com/codecrafters-io/interpreter-starter-go/internal/parser"
	"github.com/codecrafters-io/interpreter-starter-go/internal/scanner"
)

type Lox struct {
	loxCode   string
	hadErrors bool

	Scanner     *scanner.Scanner
	Parser      *parser.Parser
	Interpreter *interpreter.Interpreter

	ASTPrinter *parser.ASTPrinter
}

func NewLox(loxCode string) *Lox {
	return &Lox{loxCode: loxCode}
}

func (lox *Lox) Scan() error {
	lox.Scanner = scanner.NewScanner(lox.loxCode)
	lox.Scanner.ScanTokens()
	if lox.Scanner.HadErrors() {
		lox.hadErrors = true
	}
	return nil
}

func (lox *Lox) Parse() error {
	lox.Scan()
	lox.Parser = parser.NewParser(lox.Scanner.GetTokens())
	lox.Parser.Parse()
	if lox.Parser.HadErrors() {
		lox.hadErrors = true
	}
	lox.ASTPrinter = parser.NewASTPrinter()
	return nil
}

func (lox *Lox) Interpret() error {
	lox.Parse()
	lox.Interpreter = interpreter.NewInterpreter()
	lox.Interpreter.Interpret(lox.Parser.GetParsedExpression())
	if lox.Interpreter.HadErrors() {
		lox.hadErrors = true
	}
	return nil
}

func (lox *Lox) HadErrors() bool {
	return lox.hadErrors
}
