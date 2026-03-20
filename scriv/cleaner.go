package scriv

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/calmecac-dev/voluta/ast"
)

var (
	reScrivCharStyle = regexp.MustCompile(`<\$Scr_Cs::\d+>(.*?)<!\$Scr_Cs::\d+>`)
	reScrivParaStyle = regexp.MustCompile(`<[!]?\$Scr_Ps::\d+>`)
	reScrivHeading   = regexp.MustCompile(`<\$Scr_H::(\d+)>(.*?)<!\$Scr_H::\d+>`)
	reScrivLeftover  = regexp.MustCompile(`<[!]?\$Scr_[^>]+>`)
)

func cleanDocument(doc ast.Document) ast.Document {
	doc.Children = cleanNodes(doc.Children)
	return doc
}

func cleanNodes(nodes []ast.Node) []ast.Node {
	cleaned := make([]ast.Node, 0, len(nodes))
	for _, n := range nodes {
		n = cleanNode(n)
		cleaned = append(cleaned, n)
	}
	return cleaned
}

func cleanNode(n ast.Node) ast.Node {
	if n.Type == ast.NodeText {
		n.Value = cleanString(n.Value)
		return n
	}
	if n.Type == ast.NodeParagraph {
		if heading, ok := extractHeading(n); ok {
			return heading
		}
	}
	n.Children = cleanNodes(n.Children)
	return n
}

func cleanString(s string) string {
	s = reScrivCharStyle.ReplaceAllString(s, "$1")
	s = reScrivParaStyle.ReplaceAllString(s, "")
	s = reScrivLeftover.ReplaceAllString(s, "")
	return s
}

// extractHeading checks if a paragraph contains a Scrivener heading marker
// and returns a NodeHeading if found.
func extractHeading(n ast.Node) (ast.Node, bool) {
	// Collect full text of paragraph to check for heading markers
	text := collectText(n)
	match := reScrivHeading.FindStringSubmatch(text)
	if match == nil {
		return ast.Node{}, false
	}
	level := 1
	fmt.Sscanf(match[1], "%d", &level)
	content := match[2]
	// Clean remaining placeholders from content
	content = reScrivCharStyle.ReplaceAllString(content, "$1")
	content = reScrivLeftover.ReplaceAllString(content, "")
	return ast.Heading(level, ast.Text(content)), true
}

// collectText recursively collects all text values from a node tree.
func collectText(n ast.Node) string {
	if n.Type == ast.NodeText {
		return n.Value
	}
	var sb strings.Builder
	for _, child := range n.Children {
		sb.WriteString(collectText(child))
	}
	return sb.String()
}
