package htmlparser

import (
	"os"
	"testing"
)

var testFile = "test.html"
var p = new(Parser)
var data, err = os.ReadFile(testFile)
var s = string(data)

func reset(p *Parser) {
	p.pos = 0
	p.input = s
}

func TestNextChar(t *testing.T) {
	reset(p)
	c := p.nextChar()
	c2 := rune(s[0])
	if c != c2 {
		t.Errorf("Got %c; want %c", c, c2)
	}
}
