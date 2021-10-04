package htmlparser

import (
	"os"
	"testing"
)

var testFile = "test.html"
var p = new(HTMLParser)
var data, err = os.ReadFile(testFile)
var htmlstring = string(data)

func reset(p *HTMLParser) {
	p.pos = 0
	p.input = htmlstring
}

func TestParseElement(t *testing.T) {
	reset(p)

	n := Parse(htmlstring)
	n.PrintNode(0)
}
