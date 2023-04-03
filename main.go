package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KishorPokharel/calculator/eval"
	"github.com/KishorPokharel/calculator/lexer"
	"github.com/KishorPokharel/calculator/parser"
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
		tree := p.Parse()
		fmt.Printf("%f\n", eval.Eval(tree))
		// fmt.Printf("%s\n", tree)
	}
}
