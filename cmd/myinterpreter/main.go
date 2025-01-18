package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/pkg/lox"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file '%s': %v\n", filename, err)
		os.Exit(1)
	}
	if len(fileContents) > 0 {
		lox := lox.NewLox(string(fileContents))
		switch command {
		case "tokenize":
			lox.Scan()
			for _, token := range lox.Scanner.GetTokens() {
				fmt.Println(&token)
			}
			if lox.HadErrors() {
				os.Exit(65)
			}
		case "parse":
			lox.Parse()
			if lox.HadErrors() {
				os.Exit(65)
			}
			fmt.Println(lox.ASTPrinter.Print(lox.Parser.GetParsedExpression()))
		case "evaluate":
			lox.Interpret()
		}
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
