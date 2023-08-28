package main

import (
	"errors"
	"testing"

	"github.com/KishorPokharel/calculator/eval"
	"github.com/KishorPokharel/calculator/lexer"
	"github.com/KishorPokharel/calculator/parser"
)

func TestMain(t *testing.T) {
	tests := []struct {
		input      string
		output     float64
		parseError error
	}{
		{input: "3 + 4", output: 7.0, parseError: nil},
		{input: "(3 + 4)", output: 7.0, parseError: nil},
		{input: "(3 + 4) * 2", output: 14.0, parseError: nil},
		{input: "(3 + 4) - (2)", output: 5.0, parseError: nil},
		{input: "3 + 4 * 5 / 2", output: 13.0, parseError: nil},
		{input: "(3 + 4 * 5 / 2)", output: 13.0, parseError: nil},
		{input: "a=3", output: 3.0, parseError: nil},
		{input: "a", output: 3.0, parseError: nil},
		{input: "a+6", output: 9.0, parseError: nil},
		{input: "b=(45-3)", output: 42.0, parseError: nil},
		{input: "b=(a+b)", output: 45.0, parseError: nil},
		{input: "b=(b/5)", output: 9.0, parseError: nil},
		{input: "", parseError: parser.ErrNoTokens},
		{input: "c", parseError: parser.ErrUndeclaredVariable},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := parser.New(l)
		tree, err := p.Parse()
		testname := test.input
		t.Run(testname, func(t *testing.T) {
			if !errors.Is(err, test.parseError) {
				t.Fatalf("expected parseErr to be %v, got error \"%v\"", test.parseError, err)
			}
			if err == nil {
				out := eval.Eval(tree)
				if out != test.output {
					t.Fatalf("expected %s=%f , got %f ", test.input, test.output, out)
				}
			}
		})
	}
}
