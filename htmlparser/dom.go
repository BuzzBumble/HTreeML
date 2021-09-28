package htmlparser

type Node interface {
	toString() string
}

type TextNode struct {
	children []Node
	text     string
}

type ElementNode struct {
	children []Node
	data     ElementData
}

func (n *TextNode) toString() string {
	return "TextNode"
}

func (n *ElementNode) toString() string {
	return "ElementNode"
}

type AttrMap map[string]string

type ElementData struct {
	tagName    string
	attributes AttrMap
}
