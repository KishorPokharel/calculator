package ast

import "fmt"

type Node interface {
	String() string
}

type NumberNode struct {
	Value float64
}

func (n NumberNode) String() string {
	return fmt.Sprintf("%f", n.Value)
}

type AddNode struct {
	A Node
	B Node
}

func (n AddNode) String() string {
	return fmt.Sprintf("(%s + %s)", n.A, n.B)
}

type SubtractNode struct {
	A Node
	B Node
}

func (n SubtractNode) String() string {
	return fmt.Sprintf("(%s - %s)", n.A, n.B)
}

type MultiplyNode struct {
	A Node
	B Node
}

func (n MultiplyNode) String() string {
	return fmt.Sprintf("(%s * %s)", n.A, n.B)
}

type DivideNode struct {
	A Node
	B Node
}

func (n DivideNode) String() string {
	return fmt.Sprintf("(%s / %s)", n.A, n.B)
}

type NegationNode struct {
	A Node
}

func (n NegationNode) String() string {
	return fmt.Sprintf("( - %s)", n.A)
}

type AbsNode struct {
	A Node
}

func (n AbsNode) String() string {
	return fmt.Sprintf("(|%s|)", n.A)
}

type PowerNode struct {
	A Node
	B Node
}

func (n PowerNode) String() string {
	return fmt.Sprintf("(%s ^ %s)", n.A, n.B)
}

type AssignmentNode struct {
	ID string
	A  Node
}

func (n AssignmentNode) String() string {
	return fmt.Sprintf("(%s = %s)", n.ID, n.A)
}

// type IdentifierNode struct {
// 	ID string
// }

// func (n IdentifierNode) String() string {
// 	return fmt.Sprintf("(%s)", n.ID)
// }
