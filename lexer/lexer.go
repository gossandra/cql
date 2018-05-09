package lexer

import (
	"bytes"
	"strings"
	"unicode/utf8"

	"github.com/gossandra/cql/token"
)

type lexer struct {
	input []byte
	items chan Item

	start int
	pos   int
	width int
} // TODO: define

func (l *lexer) run() {
	for state := startState; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func (l *lexer) emit(t token.Token) {
	l.items <- Item{
		Token: t,
		Value: l.input[l.start:l.pos],
	}
}

// returns next rune from input
func (l *lexer) next() (r rune) {
	// TODO Check EOF
	r, l.width = utf8.DecodeRune(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) ignore() {
	l.start = l.pos
}
func (l *lexer) backup() {
	l.pos -= l.width
}
func (l *lexer) peek() (r rune) {
	r = l.next()
	l.backup()
	return r
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// Consume CQL keyword and emits Item. Case Insensitive
func (l *lexer) acceptKeyword(tok token.Token) bool {
	b := []byte(tok.String())
	if bytes.HasPrefix(l.input[l.pos:], b) ||
		bytes.HasPrefix(l.input[l.pos:], bytes.ToLower(b)) {
		l.pos += len(tok.String())
		l.emit(tok)
		return true
	}
	return false
}
