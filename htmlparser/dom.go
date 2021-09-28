package htmlparser

import (
	"fmt"
)

type Node interface {
	toString() string
	printNode(level int)
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
	return n.text
}

func (n *ElementNode) toString() string {
	return n.data.tagName
}

func (n *ElementNode) printNode(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("<%s>\n", n.toString())
	for _, c := range n.children {
		c.printNode(level + 1)
	}
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("</%s>\n", n.toString())
}

func (n *TextNode) printNode(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Println(n.toString())
}
