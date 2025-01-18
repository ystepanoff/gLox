package lox

import (
	"errors"

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
		return errors.New("scanner error")
	}
	return nil
}

func (lox *Lox) Parse() error {
	if err := lox.Scan(); err != nil {
		return err
	}
	lox.Parser = parser.NewParser(lox.Scanner.GetTokens())
	lox.Parser.Parse()
	if lox.Parser.HadErrors() {
		lox.hadErrors = true
		return errors.New("parser error")
	} else {
		lox.ASTPrinter = parser.NewASTPrinter()
	}
	return nil
}

func (lox *Lox) Interpret() error {
	if err := lox.Parse(); err != nil {
		return err
	}
	lox.Interpreter = interpreter.NewInterpreter()
	lox.Interpreter.Interpret(lox.Parser.GetParsedExpression())
	if lox.Interpreter.HadErrors() {
		lox.hadErrors = true
		return errors.New("runtime error")
	}
	return nil
}

func (lox *Lox) HadErrors() bool {
	return lox.hadErrors
}
