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
		{input: "3", output: 3.0, parseError: nil},
		{input: "-3", output: -3.0, parseError: nil},
		{input: "+3", output: 3.0, parseError: nil},
		{input: "+3-2", output: 1.0, parseError: nil},
		{input: "+4-9", output: -5.0, parseError: nil},
		{input: "-4-9", output: -13.0, parseError: nil},
		{input: "+(4-9)", output: -5.0, parseError: nil},
		{input: "-(4-9)", output: 5.0, parseError: nil},

		{input: "3 + 4", output: 7.0, parseError: nil},
		{input: "(3 + 4)", output: 7.0, parseError: nil},
		{input: "(3 + 4) * 2", output: 14.0, parseError: nil},
		{input: "(3 + 4) - (2)", output: 5.0, parseError: nil},
		{input: "3 + 4 * 5 / 2", output: 13.0, parseError: nil},
		{input: "(3 + 4 * 5 / 2)", output: 13.0, parseError: nil},

		{input: "(3 + 4", parseError: parser.ErrSyntax},
		{input: "(", parseError: parser.ErrSyntax},

		// vars
		{input: "a=3", output: 3.0, parseError: nil},
		{input: "a", output: 3.0, parseError: nil},
		{input: "a+6", output: 9.0, parseError: nil},
		{input: "b=(45-3)", output: 42.0, parseError: nil},
		{input: "b=(a+b)", output: 45.0, parseError: nil},
		{input: "b=(b/5)", output: 9.0, parseError: nil},

		{input: "", parseError: parser.ErrNoTokens},
		{input: "c", parseError: parser.ErrUndeclaredVariable},

		// abs
		{input: "|45|", output: 45.0, parseError: nil},
		{input: "|-45|", output: 45.0, parseError: nil},
		{input: "|-((3 + 4) - (2))|", output: 5.0, parseError: nil},
		{input: "|-5/5|", output: 1.0, parseError: nil},
		{input: "|55-65/5|", output: 42.0, parseError: nil},
		{input: "|(55-65)/5|", output: 2.0, parseError: nil},

		{input: "c=-3", output: -3.0, parseError: nil},
		{input: "|c|", output: 3.0, parseError: nil},
		{input: "|c-5|", output: 8.0, parseError: nil},

		{input: "|d|", parseError: parser.ErrUndeclaredVariable},
		{input: "|c", parseError: parser.ErrSyntax},
		{input: "|-((3 + 4) - (2))", parseError: parser.ErrSyntax},
		{input: "|-45", parseError: parser.ErrSyntax},

		// power
		{input: "2^2", output: 4.0, parseError: nil},
		{input: "2^(2+1)", output: 8.0, parseError: nil},
		{input: "2^(2+1)/2", output: 4.0, parseError: nil},
		{input: "2^(3)+2", output: 10.0, parseError: nil},
		{input: "(2^(3)+2)", output: 10.0, parseError: nil},
		{input: "2^3 + 1", output: 9.0, parseError: nil},
		{input: "(2^3 + 1)/3", output: 3.0, parseError: nil},
		{input: "(2^3 + 2^3)", output: 16.0, parseError: nil},
		{input: "(2^3 + 2^3) / 2", output: 8.0, parseError: nil},
		{input: "(2^3 + 2^3) / 2 + 4 - 2", output: 10.0, parseError: nil},

		{input: "-2^2", output: 4.0, parseError: nil},
		{input: "-2^3", output: -8.0, parseError: nil},
		{input: "(-2^3 + -2^3) / 2 + 4 - 2", output: -6.0, parseError: nil},

		{input: "2^3^4", output: 2417851639229258349412352.0, parseError: nil},
		{input: "((2+3)^4)/5", output: 125.0, parseError: nil},

		{input: "((2+3^4)/5", parseError: parser.ErrSyntax},

		// factorial
		{input: "2!", parseError: parser.ErrSyntax},
		{input: "(2)!", parseError: parser.ErrSyntax},
		{input: "|(2)!|", parseError: parser.ErrSyntax},
		{input: "((2+3^4)/5!", parseError: parser.ErrSyntax},

		{input: "|2|!", output: 2, parseError: nil},
		{input: "|2+1|!", output: 6, parseError: nil},
		{input: "|2+2|!", output: 24, parseError: nil},
		{input: "|2+2*3|!", output: 40320, parseError: nil},
		{input: "|2+2*3/2|!", output: 120, parseError: nil},
		{input: "|2^4|!", output: 20_922_789_888_000, parseError: nil},
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
