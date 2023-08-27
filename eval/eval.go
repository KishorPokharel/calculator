package eval

import (
	"github.com/KishorPokharel/calculator/ast"
)

func Eval(tree ast.Node) float64 {
	switch v := tree.(type) {
	case ast.NumberNode:
		return v.Value
	case ast.AddNode:
		return Eval(v.A) + Eval(v.B)
	case ast.SubtractNode:
		return Eval(v.A) - Eval(v.B)
	case ast.MultiplyNode:
		return Eval(v.A) * Eval(v.B)
	case ast.DivideNode:
		return Eval(v.A) / Eval(v.B)
	case ast.NegationNode:
		return -Eval(v.A)
	default:
		panic("Invalid ast node")
	}
}
