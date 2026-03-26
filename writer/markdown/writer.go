// Package markdown implements a writer that converts an ast.Document to Markdown.
package markdown

import (
	"strings"

	"github.com/calmecac-dev/cuicatl/ast"
)

// Write converts an ast.Document to Markdown text.
func Write(doc ast.Document) (string, error) {
	var sb strings.Builder
	for i, node := range doc.Children {
		if i > 0 {
			sb.WriteString("\n\n")
		}
		sb.WriteString(writeNode(node))
	}
	return strings.TrimSpace(sb.String()), nil
}

func writeNode(n ast.Node) string {
	switch n.Type {
	case ast.NodeParagraph:
		return writeChildren(n)
	case ast.NodeHeading:
		prefix := strings.Repeat("#", clamp(n.Level, 1, 6))
		return prefix + " " + writeChildren(n)
	case ast.NodeText:
		return n.Value
	case ast.NodeBold:
		return "**" + writeChildren(n) + "**"
	case ast.NodeItalic:
		return "*" + writeChildren(n) + "*"
	case ast.NodeStrikethrough:
		return "~~" + writeChildren(n) + "~~"
	case ast.NodeUnderline:
		return "<u>" + writeChildren(n) + "</u>"
	case ast.NodeLineBreak:
		return "  \n"
	case ast.NodeHorizontalRule:
		return "---"
	case ast.NodeImage:
		src := n.Attrs["src"]
		alt := n.Attrs["alt"]
		return "![" + alt + "](" + src + ")"
	case ast.NodeBlockQuote:
		lines := strings.Split(writeChildren(n), "\n")
		for i, l := range lines {
			lines[i] = "> " + l
		}
		return strings.Join(lines, "\n")
	case ast.NodeList:
		return writeList(n)
	case ast.NodeListItem:
		indent := strings.Repeat("  ", n.Level)
		marker := n.Value
		if marker == "" {
			marker = "-"
		}
		return indent + marker + " " + writeChildren(n)
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

func writeList(n ast.Node) string {
	var lines []string
	for _, item := range n.Children {
		if item.Type != ast.NodeListItem {
			continue
		}
		lines = append(lines, writeNode(item))
	}
	return strings.Join(lines, "\n")
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
