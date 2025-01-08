package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/pkg/parser"
	"github.com/codecrafters-io/interpreter-starter-go/pkg/scanner"
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
		switch command {
		case "tokenize":
			scanner_ := scanner.NewScanner(string(fileContents))
			scanner_.ScanTokens()
			for _, token := range scanner_.GetTokens() {
				fmt.Println(&token)
			}
			if scanner_.HadErrors() {
				os.Exit(65)
			}
		case "parse":
			scanner_ := scanner.NewScanner(string(fileContents))
			scanner_.ScanTokens()
			if scanner_.HadErrors() {
				return
			}
			parser_ := parser.NewParser(scanner_.GetTokens())
			expr := parser_.Parse()
			if parser_.HadErrors() {
				return
			}
			printer := parser.NewASTPrinter()
			fmt.Println(printer.Print(expr))
		}
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
