// Package rtf implements a reader that converts RTF documents to Voluta's AST.
package rtf

import (
	"fmt"
	"strconv"
)

type tokenType int

const (
	tokenGroupOpen  tokenType = iota // {
	tokenGroupClose                  // }
	tokenControl                     // \keyword or \keywordN
	tokenText                        // literal text
)

type token struct {
	kind     tokenType
	value    string
	param    int
	hasParam bool
}

type lexer struct {
	src []byte
	pos int
}

func newLexer(data []byte) *lexer {
	return &lexer{src: data, pos: 0}
}

func (l *lexer) tokenize() ([]token, error) {
	var tokens []token
	for l.pos < len(l.src) {
		ch := l.src[l.pos]
		switch {
		case ch == '{':
			tokens = append(tokens, token{kind: tokenGroupOpen})
			l.pos++
		case ch == '}':
			tokens = append(tokens, token{kind: tokenGroupClose})
			l.pos++
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

func (l *lexer) readControl() (token, error) {
	l.pos++ // consume '\'

	if l.pos >= len(l.src) {
		return token{}, fmt.Errorf("rtf: incomplete control word at end of file")
	}

	ch := l.src[l.pos]

	// Hex escape: \'XX
	if ch == '\'' {
		l.pos++
		if l.pos+2 > len(l.src) {
			return token{}, fmt.Errorf("rtf: incomplete hex sequence")
		}
		hex := string(l.src[l.pos : l.pos+2])
		l.pos += 2
		n, err := strconv.ParseInt(hex, 16, 32)
		if err != nil {
			return token{}, fmt.Errorf("rtf: invalid hex %q", hex)
		}
		return token{kind: tokenText, value: string(rune(n))}, nil
	}

	// Single character control symbol
	if !isLetter(ch) {
		l.pos++
		return token{kind: tokenControl, value: string(ch)}, nil
	}

	// Control word: \keyword or \keywordN
	start := l.pos
	for l.pos < len(l.src) && isLetter(l.src[l.pos]) {
		l.pos++
	}
	keyword := string(l.src[start:l.pos])

	// Optional numeric parameter
	negative := false
	if l.pos < len(l.src) && l.src[l.pos] == '-' {
		negative = true
		l.pos++
	}
	paramStart := l.pos
	for l.pos < len(l.src) && isDigit(l.src[l.pos]) {
		l.pos++
	}

	tok := token{kind: tokenControl, value: keyword}
	if l.pos > paramStart {
		n, _ := strconv.Atoi(string(l.src[paramStart:l.pos]))
		if negative {
			n = -n
		}
		tok.param = n
		tok.hasParam = true
	}

	// Consume delimiter space
	if l.pos < len(l.src) && l.src[l.pos] == ' ' {
		l.pos++
	}

	return tok, nil
}

func (l *lexer) readText() token {
	start := l.pos
	for l.pos < len(l.src) {
		ch := l.src[l.pos]
		if ch == '{' || ch == '}' || ch == '\\' || ch == '\r' || ch == '\n' {
			break
		}
		l.pos++
	}
	return token{kind: tokenText, value: string(l.src[start:l.pos])}
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
