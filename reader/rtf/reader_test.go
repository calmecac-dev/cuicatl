package rtf_test

import (
	"testing"

	"github.com/calmecac-dev/cuicatl/ast"
	"github.com/calmecac-dev/cuicatl/reader/rtf"
)

func TestReadSimpleParagraph(t *testing.T) {
	input := `{\rtf1\ansi Hola mundo\par}`
	doc, err := rtf.Read([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(doc.Children) == 0 {
		t.Fatal("expected at least one paragraph")
	}
	if doc.Children[0].Type != ast.NodeParagraph {
		t.Errorf("expected NodeParagraph, got %v", doc.Children[0].Type)
	}
}

func TestReadBold(t *testing.T) {
	input := `{\rtf1 \b bold text\b0\par}`
	doc, err := rtf.Read([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(doc.Children) == 0 {
		t.Fatal("expected at least one paragraph")
	}
	p := doc.Children[0]
	hasBold := false
	for _, child := range p.Children {
		if child.Type == ast.NodeBold {
			hasBold = true
		}
	}
	if !hasBold {
		t.Error("expected at least one Bold node")
	}
}

func TestReadItalic(t *testing.T) {
	input := `{\rtf1 \i italic text\i0\par}`
	doc, err := rtf.Read([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	p := doc.Children[0]
	hasItalic := false
	for _, child := range p.Children {
		if child.Type == ast.NodeItalic {
			hasItalic = true
		}
	}
	if !hasItalic {
		t.Error("expected at least one Italic node")
	}
}

func TestReadHexEscape(t *testing.T) {
	// \'e9 = é in latin-1
	input := `{\rtf1 caf\'e9\par}`
	doc, err := rtf.Read([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(doc.Children) == 0 {
		t.Fatal("expected at least one paragraph")
	}
}

func TestReadMultipleParagraphs(t *testing.T) {
	input := `{\rtf1 First\par Second\par}`
	doc, err := rtf.Read([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(doc.Children) != 2 {
		t.Errorf("expected 2 paragraphs, got %d", len(doc.Children))
	}
}

func TestIgnoreExtensionGroups(t *testing.T) {
	input := `{\rtf1 {\*\ignoreme internal data} visible text\par}`
	doc, err := rtf.Read([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(doc.Children) == 0 {
		t.Fatal("expected at least one paragraph")
	}
}
