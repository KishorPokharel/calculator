package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/KishorPokharel/calculator/eval"
	"github.com/KishorPokharel/calculator/lexer"
	"github.com/KishorPokharel/calculator/parser"
	"github.com/KishorPokharel/calculator/token"
)

const PROMPT = ">> "

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		tree, err := p.Parse()
		if err != nil {
			switch {
			case errors.Is(parser.ErrNoTokens, err):
				continue
			default:
				fmt.Printf("ERROR: %v\n", err)
				continue
			}
		}
		fmt.Printf("%f\n", eval.Eval(tree))
	}
}

func printTokens(l *lexer.Lexer) {
	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}
		fmt.Println(tok)
	}
}
