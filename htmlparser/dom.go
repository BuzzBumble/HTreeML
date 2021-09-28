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

type ElementData struct {
	tagName    string
	attributes AttrMap
}

type AttrMap map[string]string

func (n *TextNode) toString() string {
	return "TextNode"
}

func (n *ElementNode) toString() string {
	return "ElementNode"
}
