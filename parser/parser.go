package parser

import (
	"errors"
	"fmt"
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

var ErrNoTokens = errors.New("no tokens")

func (p *Parser) Parse() (ast.Node, error) {
	if p.curToken.Type == token.EOF {
		return nil, ErrNoTokens
	}
	result, err := p.expr()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *Parser) expr() (ast.Node, error) {
	result, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.curToken.Type != token.EOF && (p.curToken.Type == token.PLUS || p.curToken.Type == token.SUBTRACT) {
		switch p.curToken.Type {
		case token.PLUS:
			p.nextToken()
			B, err := p.term()
			if err != nil {
				return nil, err
			}
			result = ast.AddNode{A: result, B: B}
		case token.SUBTRACT:
			p.nextToken()
			B, err := p.term()
			if err != nil {
				return nil, err
			}
			result = ast.SubtractNode{A: result, B: B}
		default:
			p.nextToken()
		}
	}
	return result, nil
}

func (p *Parser) term() (ast.Node, error) {
	result, err := p.factor()
	if err != nil {
		return nil, err
	}
	for p.curToken.Type != token.EOF && (p.curToken.Type == token.MULTIPLY || p.curToken.Type == token.DIVIDE) {
		switch p.curToken.Type {
		case token.MULTIPLY:
			p.nextToken()
			B, err := p.factor()
			if err != nil {
				return nil, err
			}
			result = ast.MultiplyNode{A: result, B: B}
		case token.DIVIDE:
			p.nextToken()
			B, err := p.term()
			if err != nil {
				return nil, err
			}
			result = ast.DivideNode{A: result, B: B}
		default:
			p.nextToken()
		}
	}
	return result, nil
}

func (p *Parser) factor() (ast.Node, error) {
	if p.curToken.Type == token.NUMBER {
		defer p.nextToken()
		f, err := strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			log.Fatal("could not parse float")
		}
		return ast.NumberNode{Value: f}, nil
	}
	// ( E )
	if p.curToken.Type == token.LPAREN {
		p.nextToken()
		expr, err := p.expr()
		if err != nil {
			return nil, err
		}
		if p.curToken.Type == token.RPAREN {
			p.nextToken()
			return expr, nil
		} else {
			return nil, fmt.Errorf("invalid expression, expected )")
		}
	}
	return nil, fmt.Errorf("illegal token \"%s\"; expected a NUMBER or \"(\" token", p.curToken.Literal)
}
