package compiler

import (
	"fmt"
	"github.com/rhysd/gocaml/ast"
	"github.com/rhysd/gocaml/lexer"
	"github.com/rhysd/gocaml/parser"
	"github.com/rhysd/gocaml/token"
	"os"
)

type Compiler struct {
	// Compiler options (e.g. optimization level) go here.
}

func (c *Compiler) Compile(source *token.Source) {
	// TODO
}

func (c *Compiler) Lex(src *token.Source) chan token.Token {
	ch := make(chan token.Token)
	l := lexer.NewLexer(src, ch)
	go l.Lex()
	return ch
}

func (c *Compiler) PrintTokens(src *token.Source) {
	tokens := c.Lex(src)
	for {
		select {
		case t := <-tokens:
			fmt.Println(t.String())
			switch t.Kind {
			case token.EOF, token.ILLEGAL:
				return
			}
		}
	}
}

func (c *Compiler) Parse(src *token.Source) (*ast.AST, error) {
	tokens := c.Lex(src)
	root, err := parser.Parse(tokens)

	if err != nil {
		return nil, err
	}

	ast := &ast.AST{
		File: src,
		Root: root,
	}

	return ast, nil
}

func (c *Compiler) PrintAST(src *token.Source) {
	a, err := c.Parse(src)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	ast.Print(a)
}
