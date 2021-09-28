package htmlparser

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Parsing struct and functions

// Parser
// Basic struct for parsing a string
type Parser struct {
	pos   int
	input string
}

// Return the character at pos
func (p *Parser) nextChar() rune {
	return rune(p.input[p.pos])
}

// Check if string (starting at pos) starts with s
func (p *Parser) startsWith(s string) bool {
	return strings.HasPrefix(p.input[p.pos:], s)
}

// Check if string has been fully consumed
func (p *Parser) eof() bool {
	return (p.pos >= len(p.input))
}

// Consume the character and advance pos
func (p *Parser) consumeChar() rune {
	c := rune(p.input[p.pos])
	p.pos++
	return c
}

// Error out if consumed char is not c
func (p *Parser) consumeCheck(c rune) {
	consumed := p.consumeChar()
	if c != consumed {
		fmt.Printf("Parser.consumeCheck(): Expected %c; got %c\n", c, consumed)
		os.Exit(1)
	}
}

// Consume all characters while test() returns true
// Return string containing all consumed characters
func (p *Parser) consumeWhile(test func(c rune) bool) string {
	b := new(bytes.Buffer)
	for !p.eof() && test(p.nextChar()) {
		b.WriteByte(byte(p.consumeChar()))
	}
	return b.String()
}

// Consume all whitespace characters from pos
func (p *Parser) consumeWhitespace() {
	p.consumeWhile(unicode.IsSpace)
}

// Consume all letters/digits from pos
// Return string representing a word
func (p *Parser) parseWord() string {
	return p.consumeWhile(func(c rune) bool {
		return unicode.IsLetter(c) || unicode.IsDigit(c)
	})
}

// Parse text up to the next '<'
// Return a new TextNode containing consumed text
func (p *Parser) parseText() *TextNode {
	t := new(TextNode)
	t.text = strings.TrimSpace(p.consumeWhile(func(c rune) bool { return c != '<' }))
	return t
}

// Parse opening/closing tags, attributes, and child nodes
// Return a new ElementNode containing all parsed data
func (p *Parser) parseElement() *ElementNode {
	p.consumeCheck('<')

	tagName := p.parseWord()
	attrs := p.parseAttrs()
	p.consumeCheck('>')

	children := p.parseNodes()

	p.consumeCheck('<')
	p.consumeCheck('/')
	closingTagName := p.parseWord()
	if closingTagName != tagName {
		fmt.Printf("Parser.parseElement(): Closing tag %s does not match opening tag %s\n", closingTagName, tagName)
		os.Exit(1)
	}
	p.consumeCheck('>')
	p.consumeWhitespace()

	e := new(ElementNode)
	e.data = ElementData{tagName, attrs}
	e.children = children

	return e
}

// Parse a single Text/Element Node
func (p *Parser) parseNode() Node {
	if p.nextChar() == '<' {
		return p.parseElement()
	}
	return p.parseText()
}

// Parse element attribute
// Return 2 strings: name, value
func (p *Parser) parseAttr() (string, string) {
	n := p.parseWord()
	c := p.consumeChar()
	if c != '=' {
		fmt.Printf("Parser.parseAttr(): Expected [=]; got [%c]\n", c)
		os.Exit(1)
	}

	v := p.parseAttrValue()

	return n, v
}

// Parse attribute value
// Consumes all characters from opening to closing quotes
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

// Parse attributes of element
// Returns a map where key = attr name, value = attr value
func (p *Parser) parseAttrs() map[string]string {
	r := make(map[string]string)

	p.consumeWhitespace()
	for p.nextChar() != '>' {
		p.consumeWhitespace()
		n, v := p.parseAttr()
		r[n] = v
	}
	return r
}

// Parse until eof and return slice of Nodes
func (p *Parser) parseNodes() []Node {
	r := make([]Node, 0)

	p.consumeWhitespace()
	for !p.eof() && !p.startsWith("</") {
		p.consumeWhitespace()
		r = append(r, p.parseNode())
	}

	return r
}
