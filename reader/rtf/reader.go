package rtf

import (
	"strings"

	"github.com/calmecac-dev/voluta/ast"
)

// Options configures the RTF reader behavior.
type Options struct {
	UnderlineAsHTML bool
	ImageHandler    func(data []byte, format string) (string, error)
}

// Read converts RTF data to an ast.Document using default options.
func Read(data []byte) (ast.Document, error) {
	return ReadWithOptions(data, Options{})
}

// ReadWithOptions converts RTF data to an ast.Document with custom options.
func ReadWithOptions(data []byte, opts Options) (ast.Document, error) {
	tokens, err := Tokenize(data)
	if err != nil {
		return ast.Document{}, err
	}
	c := newConverter(opts)
	c.processTokens(tokens)
	c.flush()
	c.doc = groupListItems(c.doc)
	return c.doc, nil
}

// --- Properties: formatting state for the current group ---

type properties struct {
	bold         bool
	italic       bool
	underline    bool
	strike       bool
	outlineLevel int // -1 = normal paragraph
	listID       int // 0 = not a list
	listLevel    int
}

func defaultProperties() properties {
	return properties{outlineLevel: -1}
}

// --- List stack ---

type listItem struct {
	listID    int
	listLevel int
	marker    string
	nodes     []ast.Node
}

// --- Converter ---

type converter struct {
	opts         Options
	propStack    []properties
	cur          properties
	paragraphBuf []ast.Node
	doc          ast.Document
	listStack    []listItem
	nextMarker   string
}

func newConverter(opts Options) *converter {
	return &converter{
		opts: opts,
		cur:  defaultProperties(),
	}
}

func (c *converter) pushGroup() {
	c.propStack = append(c.propStack, c.cur)
}

func (c *converter) popGroup() {
	if len(c.propStack) > 0 {
		c.cur = c.propStack[len(c.propStack)-1]
		c.propStack = c.propStack[:len(c.propStack)-1]
	}
}

// processTokens processes a flat list of tokens at the current group level.
func (c *converter) processTokens(tokens []Token) {
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		switch tok.Kind {
		case TokenGroup:
			c.processGroup(tok.Children)
		case TokenControlWord:
			c.handleControl(tok)
		case TokenControlSymbol:
			c.handleSymbol(tok.Value)
		case TokenText:
			c.handleText(tok.Value)
		}
	}
}

// processGroup handles a { ... } group, identifying special groups by
// their first token — exactly like Pandoc does.
func (c *converter) processGroup(children []Token) {
	if len(children) == 0 {
		return
	}

	first := children[0]

	// Extension group {\* ...} — identify by content
	if first.Kind == TokenControlSymbol && first.Value == "*" {
		// ignored extension group — check for pntext inside
		return
	}

	// {\pntext ...} — list marker group
	if first.Kind == TokenControlWord && first.Value == "pntext" {
		c.nextMarker = extractText(children[1:])
		return
	}

	// Ignored metadata/formatting groups
	if first.Kind == TokenControlWord {
		switch first.Value {
		case "fonttbl", "colortbl", "stylesheet", "info",
			"listtable", "listoverridetable",
			"pict", "object", "fldinst", "fldrslt",
			"pgdsc", "wgrffmtfilter", "themedata",
			"colorschememapping", "datastore", "latentstyles",
			"pntxta", "pntxtb", "xmlnstbl", "filetbl",
			"expandedcolortbl", "revtbl", "header", "footer",
			"headerl", "headerr", "footerr", "footerl":
			return
		}
	}

	// Normal group — inherit current properties, process children, restore
	c.pushGroup()
	c.processTokens(children)
	c.popGroup()
}

func (c *converter) handleControl(tok Token) {
	switch tok.Value {
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
	case "par":
		c.flushParagraph()
	case "pard":
		c.flushParagraph()
		c.cur.outlineLevel = -1
		c.cur.listID = 0
		c.cur.listLevel = 0
	case "line":
		c.paragraphBuf = append(c.paragraphBuf, ast.Node{Type: ast.NodeLineBreak})
	case "u":
		if tok.HasParam {
			r := rune(tok.Param)
			if r < 0 {
				r += 65536
			}
			c.handleText(string(r))
		}
	case "ls":
		if tok.HasParam {
			c.cur.listID = tok.Param
		}
	case "ilvl":
		if tok.HasParam {
			c.cur.listLevel = tok.Param
		}
	case "plain":
		ol := c.cur.outlineLevel
		lid := c.cur.listID
		ll := c.cur.listLevel
		c.cur = defaultProperties()
		c.cur.outlineLevel = ol
		c.cur.listID = lid
		c.cur.listLevel = ll
	case "outlinelevel":
		if tok.HasParam {
			c.cur.outlineLevel = tok.Param
		}
	case "listtext":
		// old-style list marker — handled as group in processGroup
	}
}

func (c *converter) handleSymbol(sym string) {
	switch sym {
	case "-":
		c.handleText("\u00ad") // soft hyphen
	case "~":
		c.handleText("\u00a0") // non-breaking space
	case "_":
		c.handleText("\u2011") // non-breaking hyphen
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
	var node ast.Node
	if c.cur.outlineLevel >= 0 {
		node = ast.Heading(c.cur.outlineLevel+1, c.paragraphBuf...)
	} else if c.cur.listID > 0 {
		marker := c.nextMarker
		if marker == "" {
			marker = "-"
		}
		node = ast.Node{
			Type:     ast.NodeListItem,
			Value:    marker,
			Level:    c.cur.listLevel,
			Children: c.paragraphBuf,
		}
		c.nextMarker = ""
	} else {
		node = ast.Paragraph(c.paragraphBuf...)
	}
	c.doc.Children = append(c.doc.Children, node)
	c.paragraphBuf = nil
}

func (c *converter) flush() {
	c.flushParagraph()
}

// extractText recursively collects visible text from a token list,
// used to extract list markers from {\pntext} groups.
func extractText(tokens []Token) string {
	var sb strings.Builder
	for _, tok := range tokens {
		switch tok.Kind {
		case TokenText:
			t := strings.TrimSpace(tok.Value)
			if t != "" && t != "\t" {
				sb.WriteString(t)
			}
		case TokenGroup:
			sb.WriteString(extractText(tok.Children))
		}
	}
	return sb.String()
}

func paramIsOn(tok Token) bool {
	if !tok.HasParam {
		return true
	}
	return tok.Param != 0
}

func groupListItems(doc ast.Document) ast.Document {
	doc.Children = groupNodes(doc.Children)
	return doc
}

func groupNodes(nodes []ast.Node) []ast.Node {
	var result []ast.Node
	i := 0
	for i < len(nodes) {
		if nodes[i].Type == ast.NodeListItem {
			list := ast.Node{Type: ast.NodeList}
			for i < len(nodes) && nodes[i].Type == ast.NodeListItem {
				list.Children = append(list.Children, nodes[i])
				i++
			}
			result = append(result, list)
		} else {
			result = append(result, nodes[i])
			i++
		}
	}
	return result
}
