package htmlparser

import (
	"fmt"
)

type Node interface {
	toString() string
	PrintNode(level int)
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

func (n *ElementNode) PrintNode(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("<%s", n.toString())
	for k, v := range n.data.attributes {
		fmt.Printf(" %s=\"%s\"", k, v)
	}
	fmt.Println(">")
	for _, c := range n.children {
		c.PrintNode(level + 1)
	}
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("</%s>\n", n.toString())
}

func (n *TextNode) PrintNode(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Println(n.toString())
}
