package parser

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/KishorPokharel/calculator/ast"
	"github.com/KishorPokharel/calculator/eval"
	"github.com/KishorPokharel/calculator/lexer"
	"github.com/KishorPokharel/calculator/token"
)

var (
	ErrNoTokens           = errors.New("no tokens")
	ErrUndeclaredVariable = errors.New("undeclared variable")
	ErrSyntax             = errors.New("syntax error")
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

func (p *Parser) Parse() (ast.Node, error) {
	if p.curToken.Type == token.EOF {
		return nil, ErrNoTokens
	}
	// assignment
	if p.curToken.Type == token.IDENTIFIER && p.peekToken.Type == token.ASSIGN {
		id := p.curToken.Literal
		p.nextToken() // =
		p.nextToken()
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		if p.curToken.Type != token.EOF {
			return nil, fmt.Errorf("%w: expected assignment statement to end, found %s", ErrSyntax, p.curToken.Literal)
		}
		return ast.AssignmentNode{ID: id, A: expr}, nil
	}

	// expression
	result, err := p.expression()
	if err != nil {
		return nil, err
	}
	if p.curToken.Type != token.EOF {
		return nil, fmt.Errorf("%w: expected expression to end, found %s", ErrSyntax, p.curToken.Literal)
	}
	return result, nil
}

func (p *Parser) expression() (ast.Node, error) {
	if p.curToken.Type == token.EOF {
		return nil, fmt.Errorf("%w", ErrSyntax)
	}
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
		}
	}
	return result, nil
}

func (p *Parser) term() (ast.Node, error) {
	result, err := p.primary()
	if err != nil {
		return nil, err
	}
	for p.curToken.Type != token.EOF && (p.curToken.Type == token.MULTIPLY || p.curToken.Type == token.DIVIDE) {
		switch p.curToken.Type {
		case token.MULTIPLY:
			p.nextToken()
			B, err := p.primary()
			if err != nil {
				return nil, err
			}
			result = ast.MultiplyNode{A: result, B: B}
		case token.DIVIDE:
			p.nextToken()
			B, err := p.primary()
			if err != nil {
				return nil, err
			}
			result = ast.DivideNode{A: result, B: B}
		}
	}
	return result, nil
}

func (p *Parser) primary() (ast.Node, error) {
	result, err := p.factor()
	if err != nil {
		return nil, err
	}
	if p.curToken.Type != token.EOF && p.curToken.Type == token.POWER {
		p.nextToken()
		B, err := p.primary() // right associative
		if err != nil {
			return nil, err
		}
		result = ast.PowerNode{A: result, B: B}
	}
	return result, nil
}

func (p *Parser) factor() (ast.Node, error) {
	if p.curToken.Type == token.NUMBER {
		f, err := strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse float")
		}
		p.nextToken()
		return ast.NumberNode{Value: f}, nil
	}

	// ID
	if p.curToken.Type == token.IDENTIFIER {
		id := p.curToken.Literal
		val, ok := eval.State[id]
		if !ok {
			return nil, fmt.Errorf("%w \"%s\"", ErrUndeclaredVariable, p.curToken.Literal)
		}
		p.nextToken()
		return ast.NumberNode{Value: val}, nil
	}

	// "-" Factor
	if p.curToken.Type == token.SUBTRACT {
		p.nextToken()
		res, err := p.factor()
		if err != nil {
			return nil, err
		}
		return ast.NegationNode{A: res}, nil
	}

	// "+" Factor
	if p.curToken.Type == token.PLUS {
		p.nextToken()
		res, err := p.factor()
		if err != nil {
			return nil, err
		}
		return ast.UnaryPlusNode{A: res}, nil
	}

	// "|" E "|"
	if p.curToken.Type == token.BAR {
		p.nextToken()
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		if p.curToken.Type == token.BAR {
			p.nextToken()
			return ast.AbsNode{A: expr}, nil
		} else {
			return nil, fmt.Errorf("%w: invalid expression, expected a closing |", ErrSyntax)
		}
	}

	// "(" E ")"
	if p.curToken.Type == token.LPAREN {
		p.nextToken()
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		if p.curToken.Type == token.RPAREN {
			p.nextToken()
			return expr, nil
		} else {
			return nil, fmt.Errorf("%w: invalid expression, expected a closing \")\"", ErrSyntax)
		}
	}

	return nil, fmt.Errorf("%w: illegal token \"%s\"", ErrSyntax, p.curToken.Literal)
}
