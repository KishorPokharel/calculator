package eval

import (
	"math"

	"github.com/KishorPokharel/calculator/ast"
)

var State = map[string]float64{}

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
	case ast.UnaryPlusNode:
		return Eval(v.A)
	case ast.AbsNode:
		res := Eval(v.A)
		if res < 0 {
			return -res
		}
		return res
	case ast.PowerNode:
		return math.Pow(Eval(v.A), Eval(v.B))
	case ast.FactorialNode:
		val := Eval(v.A)
		return factorial(int(val))
	case ast.AssignmentNode:
		result := Eval(v.A)
		State[v.ID] = result
		return result
	// case ast.IdentifierNode:
	// 	result := State[v.ID]
	// 	return result
	default:
		panic("Invalid ast node")
	}
}

func factorial(v int) float64 {
	res := 1
	for i := v; i > 0; i-- {
		res *= i
	}
	return float64(res)
}
