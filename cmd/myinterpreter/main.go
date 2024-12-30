package main

import (
	"fmt"
	"github.com/codecrafters-io/interpreter-starter-go/pkg/scanner"
	"os"
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
				fmt.Println(token)
			}
			if scanner_.HadErrors() {
				os.Exit(65)
			}
		case "parse":
			fmt.Println("Not implemented")
		}
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
