package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Parser struct {
	pos   int
	input string
}

func (p *Parser) nextChar() rune {
	return rune(p.input[p.pos])
}

func (p *Parser) startsWith(s string) bool {
	return strings.HasPrefix(p.input[p.pos:], s)
}

func (p *Parser) eof() bool {
	return (p.pos >= len(p.input))
}

func (p *Parser) consumeChar() rune {
	c := rune(p.input[p.pos])
	p.pos++
	return c
}

func (p *Parser) consumeWhile(test func(c rune) bool) string {
	b := new(bytes.Buffer)
	for !p.eof() && test(p.nextChar()) {
		b.WriteByte(byte(p.consumeChar()))
	}
	return b.String()
}

func (p *Parser) consumeWhitespace() {
	p.consumeWhile(unicode.IsSpace)
}

func (p *Parser) parseTagName() string {
	return p.consumeWhile(func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsDigit(c)
	})
}

func (p *Parser) parseText() *TextNode {
	t := new(TextNode)
	t.text = p.consumeWhile(func(c rune) bool { return c != '<' })
	return t
}

func (p *Parser) parseElement() *ElementNode {
	c := p.consumeChar()
	if c != '<' {
		fmt.Printf("Parser.parseElement(): Expected [<]; got [%c]\n", c)
		os.Exit(1)
	}
	// Parse Attributes
	return nil
}

func (p *Parser) parseAttr() (string, string) {
	n := p.parseTagName()
	c := p.consumeChar()
	if c != '=' {
		fmt.Printf("Parser.parseAttr(): Expected [=]; got [%c]\n", c)
		os.Exit(1)
	}

	v := p.parseAttrValue()

	return n, v
}

func (p *Parser) parseAttrValue() string {
	q := p.consumeChar()
	if q != '"' && q != '\'' {
		fmt.Printf("Parser.parseAttrValue(): Expected opening quote; Got [%c]\n", q)
		os.Exit(1)
	}
	v := p.consumeWhile(func(c rune) bool { return c != q })
	q2 := p.consumeChar()
	if q != q2 {
		fmt.Printf("Parser.consumeWhile(): Expected closing quote [%c]; Got [%c]\n", q, q2)
		os.Exit(1)
	}

	return v
}

// TODO: Add parseAttrs
