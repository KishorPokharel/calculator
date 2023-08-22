package parser

import (
	"log"
	"strconv"

	"github.com/KishorPokharel/calculator/ast"
	"github.com/KishorPokharel/calculator/lexer"
	"github.com/KishorPokharel/calculator/token"
)

type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

// nextToken consumes the token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// func (p *Parser) Parse() ast.Node {
// 	return ast.DivideNode{
// 		A: ast.AddNode{
// 			A: ast.NumberNode{Value: 5},
// 			B: ast.MultiplyNode{
// 				A: ast.NumberNode{Value: 2},
// 				B: ast.NumberNode{Value: 3},
// 			},
// 		},
// 		B: ast.NumberNode{Value: 2},
// 	}
// }

func (p *Parser) Parse() ast.Node {
	if p.curToken.Type != token.EOF {
		result := p.expr()
		return result
	}
	return nil
}

func (p *Parser) expr() ast.Node {
	result := p.term()
	for p.curToken.Type != token.EOF && (p.curToken.Type == token.PLUS || p.curToken.Type == token.SUBTRACT) {
		switch p.curToken.Type {
		case token.PLUS:
			p.nextToken()
			result = ast.AddNode{A: result, B: p.term()}
		case token.SUBTRACT:
			p.nextToken()
			result = ast.SubtractNode{A: result, B: p.term()}
		default:
			p.nextToken()
		}
	}
	return result
}

func (p *Parser) term() ast.Node {
	result := p.factor()
	for p.curToken.Type != token.EOF && (p.curToken.Type == token.MULTIPLY || p.curToken.Type == token.DIVIDE) {
		switch p.curToken.Type {
		case token.MULTIPLY:
			p.nextToken()
			result = ast.MultiplyNode{A: result, B: p.factor()}
		case token.DIVIDE:
			p.nextToken()
			result = ast.DivideNode{A: result, B: p.term()}
		default:
			p.nextToken()
		}
	}
	return result
}

func (p *Parser) factor() ast.Node {
	if p.curToken.Type == token.NUMBER {
		defer p.nextToken()
		f, err := strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			log.Fatal("could not parse float")
		}
		return ast.NumberNode{Value: f}
	}
	// ( E )
	if p.curToken.Type == token.LPAREN {
		p.nextToken()
		expr := p.expr()
		if p.curToken.Type == token.RPAREN {
			p.nextToken()
			return expr
		} else {
			log.Println("invalid expression, expected )")
			return nil
		}
	}
	return nil
}
