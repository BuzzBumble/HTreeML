package main

type Node interface {
}

type TextNode struct {
	children []Node
}

type ElementNode struct {
	children []Node
}

type AttrMap map[string]string

type ElementData struct {
	tagName    string
	attributes AttrMap
}
