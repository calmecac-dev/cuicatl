package rtf

import (
	"strings"

	"github.com/calmecac-dev/voluta/ast"
)

// Options configures the RTF reader behavior.
type Options struct {
	// UnderlineAsHTML wraps underlined text in <u> instead of ignoring it.
	UnderlineAsHTML bool
	// ImageHandler is called with the bytes of each image found.
	// It should return the path or URL to use in the ast.NodeImage.
	// If nil, images are skipped.
	ImageHandler func(data []byte, format string) (string, error)
}

// Read converts RTF data to an ast.Document using default options.
func Read(data []byte) (ast.Document, error) {
	return ReadWithOptions(data, Options{})
}

// ReadWithOptions converts RTF data to an ast.Document with custom options.
func ReadWithOptions(data []byte, opts Options) (ast.Document, error) {
	l := newLexer(data)
	tokens, err := l.tokenize()
	if err != nil {
		return ast.Document{}, err
	}
	c := &converter{tokens: tokens, opts: opts}
	return c.convert()
}

// --- converter: tokens → ast.Document ---

type converterState struct {
	bold      bool
	italic    bool
	underline bool
	strike    bool
	ignore    bool
}

type converter struct {
	tokens       []token
	opts         Options
	stateStack   []converterState
	cur          converterState
	paragraphBuf []ast.Node
	doc          ast.Document
}

func (c *converter) convert() (ast.Document, error) {
	for _, tok := range c.tokens {
		switch tok.kind {
		case tokenGroupOpen:
			c.stateStack = append(c.stateStack, c.cur)
		case tokenGroupClose:
			if len(c.stateStack) > 0 {
				c.cur = c.stateStack[len(c.stateStack)-1]
				c.stateStack = c.stateStack[:len(c.stateStack)-1]
			}
		case tokenControl:
			if c.cur.ignore {
				continue
			}
			c.handleControl(tok)
		case tokenText:
			if c.cur.ignore {
				continue
			}
			c.handleText(tok.value)
		}
	}
	c.flushParagraph()
	return c.doc, nil
}

func (c *converter) handleControl(tok token) {
	switch tok.value {
	case "*":
		c.cur.ignore = true
	case "b":
		c.cur.bold = paramIsOn(tok)
	case "i":
		c.cur.italic = paramIsOn(tok)
	case "ul":
		c.cur.underline = true
	case "ulnone":
		c.cur.underline = false
	case "strike", "striked":
		c.cur.strike = paramIsOn(tok)
	case "par", "pard":
		c.flushParagraph()
	case "line":
		c.paragraphBuf = append(c.paragraphBuf, ast.Node{Type: ast.NodeLineBreak})
	case "u":
		if tok.hasParam {
			r := rune(tok.param)
			if r < 0 {
				r += 65536
			}
			c.handleText(string(r))
		}
	case "fonttbl", "colortbl", "stylesheet", "info",
		"pict", "object", "fldinst", "fldrslt":
		c.cur.ignore = true
	}
}

func (c *converter) handleText(value string) {
	if strings.TrimSpace(value) == "" {
		return
	}
	node := ast.Text(value)
	if c.cur.strike {
		node = ast.Node{Type: ast.NodeStrikethrough, Children: []ast.Node{node}}
	}
	if c.cur.underline {
		node = ast.Node{Type: ast.NodeUnderline, Children: []ast.Node{node}}
	}
	if c.cur.italic {
		node = ast.Italic(node)
	}
	if c.cur.bold {
		node = ast.Bold(node)
	}
	c.paragraphBuf = append(c.paragraphBuf, node)
}

func (c *converter) flushParagraph() {
	if len(c.paragraphBuf) == 0 {
		return
	}
	c.doc.Children = append(c.doc.Children, ast.Paragraph(c.paragraphBuf...))
	c.paragraphBuf = nil
}

func paramIsOn(tok token) bool {
	if !tok.hasParam {
		return true
	}
	return tok.param != 0
}
