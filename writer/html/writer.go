// Package html implements a writer that converts an ast.Document to HTML.
package html

import (
	"fmt"
	"strings"

	"github.com/calmecac-dev/cuicatl/ast"
)

// Write converts an ast.Document to HTML.
func Write(doc ast.Document) (string, error) {
	var sb strings.Builder
	for _, node := range doc.Children {
		sb.WriteString(writeNode(node))
		sb.WriteString("\n")
	}
	return sb.String(), nil
}

func writeNode(n ast.Node) string {
	switch n.Type {
	case ast.NodeParagraph:
		return "<p>" + writeChildren(n) + "</p>"
	case ast.NodeHeading:
		l := clamp(n.Level, 1, 6)
		return fmt.Sprintf("<h%d>%s</h%d>", l, writeChildren(n), l)
	case ast.NodeText:
		return n.Value
	case ast.NodeBold:
		return "<strong>" + writeChildren(n) + "</strong>"
	case ast.NodeItalic:
		return "<em>" + writeChildren(n) + "</em>"
	case ast.NodeStrikethrough:
		return "<s>" + writeChildren(n) + "</s>"
	case ast.NodeUnderline:
		return "<u>" + writeChildren(n) + "</u>"
	case ast.NodeLineBreak:
		return "<br>"
	case ast.NodeHorizontalRule:
		return "<hr>"
	case ast.NodeImage:
		return fmt.Sprintf(`<img src="%s" alt="%s">`, n.Attrs["src"], n.Attrs["alt"])
	case ast.NodeBlockQuote:
		return "<blockquote>" + writeChildren(n) + "</blockquote>"
	case ast.NodeList:
		tag := "ul"
		if n.Attrs["ordered"] == "true" {
			tag = "ol"
		}
		return "<" + tag + ">" + writeChildren(n) + "</" + tag + ">"
	case ast.NodeListItem:
		return "<li>" + writeChildren(n) + "</li>"
	default:
		return writeChildren(n)
	}
}

func writeChildren(n ast.Node) string {
	var sb strings.Builder
	for _, child := range n.Children {
		sb.WriteString(writeNode(child))
	}
	return sb.String()
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
