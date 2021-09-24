package main

type Node interface {
}

type TextNode struct {
	children []Node
	text     string
}

type ElementNode struct {
	children []Node
	data     ElementData
}

type AttrMap map[string]string

type ElementData struct {
	tagName    string
	attributes AttrMap
}
