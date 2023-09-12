package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/KishorPokharel/calculator/eval"
	"github.com/KishorPokharel/calculator/lexer"
	"github.com/KishorPokharel/calculator/parser"

	"github.com/chzyer/readline"
)

const PROMPT = ">> "

func main() {
	rl, err := readline.New(PROMPT)
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
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

// func printTokens(l *lexer.Lexer) {
// 	for {
// 		tok := l.NextToken()
// 		if tok.Type == token.EOF {
// 			break
// 		}
// 		fmt.Println(tok)
// 	}
// }
