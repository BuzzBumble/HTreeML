package main

import (
	"bytes"
	"strings"
)

type Parser struct {
	pos   int
	input string
}

func (p *Parser) nextChar() byte {
	return p.input[p.pos]
}

func (p *Parser) startsWith(s string) bool {
	return strings.HasPrefix(p.input[p.pos:], s)
}

func (p *Parser) eof() bool {
	return (p.pos >= len(p.input))
}

func (p *Parser) consumeChar() byte {
	c := p.input[p.pos]
	p.pos++
	return c
}

func (p *Parser) consumeWhile(test func(c byte) bool) string {
	b := new(bytes.Buffer)
	for !p.eof() && test(p.nextChar()) {
		b.WriteByte(p.consumeChar())
	}
	return b.String()
}
