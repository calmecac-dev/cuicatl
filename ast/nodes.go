// Package ast defines the internal document representation used by Voluta
// as an intermediate format between readers and writers.
package ast

// NodeType identifies the type of each node in the document tree.
type NodeType int

const (
	NodeParagraph NodeType = iota
	NodeHeading
	NodeText
	NodeBold
	NodeItalic
	NodeUnderline
	NodeStrikethrough
	NodeLineBreak
	NodeHorizontalRule
	NodeBlockQuote
	NodeImage
	NodeList
	NodeListItem
	NodeTable
	NodeTableRow
	NodeTableCell
)

// Meta holds document-level metadata.
type Meta struct {
	Title    string
	Author   string
	Language string
	Extra    map[string]string
}

// Document is the root of the document tree.
type Document struct {
	Meta     Meta
	Children []Node
}

// Node represents a single element in the document tree.
type Node struct {
	Type     NodeType
	Value    string            // leaf text content
	Level    int               // heading level (1-6), list nesting depth
	Attrs    map[string]string // src, alt, href, ordered, etc.
	Children []Node
}

// Helpers for building common nodes

func Text(value string) Node {
	return Node{Type: NodeText, Value: value}
}

func Paragraph(children ...Node) Node {
	return Node{Type: NodeParagraph, Children: children}
}

func Heading(level int, children ...Node) Node {
	return Node{Type: NodeHeading, Level: level, Children: children}
}

func Bold(children ...Node) Node {
	return Node{Type: NodeBold, Children: children}
}

func Italic(children ...Node) Node {
	return Node{Type: NodeItalic, Children: children}
}

func Image(src, alt string) Node {
	return Node{
		Type:  NodeImage,
		Attrs: map[string]string{"src": src, "alt": alt},
	}
}
