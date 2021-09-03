package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KishorPokharel/calculator/lexer"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}