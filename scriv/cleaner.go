package scriv

import (
	"regexp"

	"github.com/calmecac-dev/voluta/ast"
)

var (
	reScrivCharStyle = regexp.MustCompile(`<\$Scr_Cs::\d+>(.*?)<!\$Scr_Cs::\d+>`)
	reScrivParaStyle = regexp.MustCompile(`<[!]?\$Scr_Ps::\d+>`)
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
	n.Children = cleanNodes(n.Children)
	return n
}

func cleanString(s string) string {
	s = reScrivCharStyle.ReplaceAllString(s, "$1")
	s = reScrivParaStyle.ReplaceAllString(s, "")
	s = reScrivLeftover.ReplaceAllString(s, "")
	return s
}
