package rtf

import (
	"fmt"
	"strconv"
)

// Token represents a single element in the RTF token tree.
type Token struct {
	Kind     TokenKind
	Value    string // control word name or text content
	Param    int
	HasParam bool
	Children []Token // only for TokenGroup
}

type TokenKind int

const (
	TokenControlWord   TokenKind = iota // \keyword or \keywordN
	TokenControlSymbol                  // single non-alpha char after \
	TokenText                           // literal text
	TokenGroup                          // { ... } — contains Children
)

// Tokenize parses RTF bytes into a token tree.
// The top level is a slice of tokens (usually one root group).
func Tokenize(data []byte) ([]Token, error) {
	l := &lexer{src: data, pos: 0}
	return l.tokenizeGroup()
}

type lexer struct {
	src []byte
	pos int
}

// tokenizeGroup reads tokens until EOF or closing '}'.
func (l *lexer) tokenizeGroup() ([]Token, error) {
	var tokens []Token
	for l.pos < len(l.src) {
		ch := l.src[l.pos]
		switch {
		case ch == '{':
			l.pos++
			children, err := l.tokenizeGroup()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, Token{Kind: TokenGroup, Children: children})
		case ch == '}':
			l.pos++ // consume closing brace
			return tokens, nil
		case ch == '\\':
			tok, err := l.readControl()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, tok)
		case ch == '\r' || ch == '\n':
			l.pos++
		default:
			tokens = append(tokens, l.readText())
		}
	}
	return tokens, nil
}

func (l *lexer) readControl() (Token, error) {
	l.pos++ // consume '\'

	if l.pos >= len(l.src) {
		return Token{}, fmt.Errorf("rtf: incomplete control word at end of file")
	}

	ch := l.src[l.pos]

	// Hex escape: \'XX
	if ch == '\'' {
		l.pos++
		if l.pos+2 > len(l.src) {
			return Token{}, fmt.Errorf("rtf: incomplete hex escape")
		}
		hex := string(l.src[l.pos : l.pos+2])
		l.pos += 2
		n, err := strconv.ParseInt(hex, 16, 32)
		if err != nil {
			return Token{}, fmt.Errorf("rtf: invalid hex %q", hex)
		}
		return Token{Kind: TokenText, Value: string(rune(n))}, nil
	}

	// Single character control symbol (non-alpha)
	if !isLetter(ch) {
		l.pos++
		return Token{Kind: TokenControlSymbol, Value: string(ch)}, nil
	}

	// Control word: \keyword or \keywordN
	start := l.pos
	for l.pos < len(l.src) && isLetter(l.src[l.pos]) {
		l.pos++
	}
	name := string(l.src[start:l.pos])

	// Optional numeric parameter (may be negative)
	negative := false
	if l.pos < len(l.src) && l.src[l.pos] == '-' {
		negative = true
		l.pos++
	}
	paramStart := l.pos
	for l.pos < len(l.src) && isDigit(l.src[l.pos]) {
		l.pos++
	}

	tok := Token{Kind: TokenControlWord, Value: name}
	if l.pos > paramStart {
		n, _ := strconv.Atoi(string(l.src[paramStart:l.pos]))
		if negative {
			n = -n
		}
		tok.Param = n
		tok.HasParam = true
	}

	// Consume optional space delimiter
	if l.pos < len(l.src) && l.src[l.pos] == ' ' {
		l.pos++
	}

	return tok, nil
}

func (l *lexer) readText() Token {
	start := l.pos
	for l.pos < len(l.src) {
		ch := l.src[l.pos]
		if ch == '{' || ch == '}' || ch == '\\' || ch == '\r' || ch == '\n' {
			break
		}
		l.pos++
	}
	return Token{Kind: TokenText, Value: string(l.src[start:l.pos])}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
