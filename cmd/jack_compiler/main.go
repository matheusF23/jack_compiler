package main

import (
	"fmt"
	"os"

	"jack_compiler/scanner"
	"jack_compiler/token"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run ./cmd/compiler <arquivo.jack>")
		return
	}

	filename := os.Args[1]

	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	scan := scanner.NewScanner(input)

	fmt.Println("<tokens>")

	for tk := scan.NextToken(); tk.Type != token.EOF; tk = scan.NextToken() {
		fmt.Println(tk)
	}

	fmt.Println("</tokens>")
}
