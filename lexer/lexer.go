package lexer

import (
	"bytes"
	"fmt"
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
	if l.start == l.pos {
		return
	}
	l.items <- Item{
		Token: t,
		Value: l.input[l.start:l.pos],
	}
	l.ignore()
}

// returns next rune from input
const RuneEOF rune = -1

func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		return RuneEOF
	}
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

// Consume CQL token and emits Itema.
func (l *lexer) acceptToken(tok token.Token) bool {
	b := []byte(tok.String())
	if bytes.HasPrefix(l.input[l.pos:], b) {
		l.pos += len(tok.String())
		l.emit(tok)
		return true
	}
	return false
}

// Skips the whitespaces.
// Do not use in comments context
func (l *lexer) skip() {
	l.acceptRun(" \n\t")
	l.ignore()
}

// error returns an error token and terminates the scan
// by passing back a nil pointer that will be the next
// state, terminating l.run.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- Item{
		token.ILLEGAL,
		[]byte(fmt.Sprintf(format, args...)),
	}
	return nil
}
