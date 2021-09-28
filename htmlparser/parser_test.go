package htmlparser

import (
	"os"
	"testing"
)

var testFile = "test.html"
var p = new(Parser)
var data, err = os.ReadFile(testFile)
var htmlstring = string(data)

func reset(p *Parser) {
	p.pos = 0
	p.input = htmlstring
}

func TestParseElement(t *testing.T) {
	reset(p)

	e := p.parseElement()

	if e.data.tagName != "html" {
		t.Errorf("Expected tagName = html; Got %s", e.data.tagName)
	}
	if len(e.children) != 2 {
		t.Errorf("Expected 2 children; Got %d", len(e.children))
	}

	e.printNode(0)
}
