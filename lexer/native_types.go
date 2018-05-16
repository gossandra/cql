/*
Native types docs: http://cassandra.apache.org/doc/latest/cql/definitions.html#constants
*/
package lexer

import (
	"bytes"
	"errors"

	"github.com/gossandra/cql/token"
)

var InvalidUUID = errors.New("Invalid UUID")

type parsingState int8

const (
	cuuid parsingState = iota
	cinteger
	cfloat
)

var lexTerm = lexConstant

func lexConstant(l *lexer) error {
	var r rune
	l.ignore()

	if l.acceptIToken(token.K_NAN) ||
		l.acceptIToken(token.K_INFINITY) ||
		l.acceptIToken(token.K_TRUE) ||
		l.acceptIToken(token.K_FALSE) ||
		l.acceptIToken(token.K_NULL) {
		return nil
	}

	// check BLOB
	if bytes.HasPrefix(l.input[l.pos:], []byte("0x")) ||
		bytes.HasPrefix(l.input[l.pos:], []byte("0X")) {
		return lexBLOB(l)
	}

	r = l.peek()
	switch {
	case r == '\'' || r == '$': // TODO: Double $$
		return lexString(l)
	case isAlphaNum(r) || r == '-':
		return lexNumber(l)
	}
	return nil
}

func lexNumber(l *lexer) error {
	var (
		t          = token.INT
		r          rune
		dashPrefix bool = l.peek() == '-'
	)
	if dashPrefix {
		l.next()
	}

	for {
		r = l.next()
		switch {
		case isNum(r):
			continue
		case r == '.':
			t = token.FLOAT
		case r == 'e' || r == 'E':
			t = token.FLOAT
			l.accept("+-")

		case (isHex(r) || r == '-'):
			if dashPrefix {
				return errors.New("invalid number syntax")
			}
			return lexUUID(l)

		default:
			l.backup()
			l.emit(t)
			return nil
		}
	}
	return errors.New("unexpected error")
}

// lexString lexes SingleQuoted or DollarQuoted strings
func lexString(l *lexer) error {
	r := l.next() // get and accept first symbol
	if r == '$' && l.next() == '$' {
		for r = l.next(); ; r = l.next() {
			if r == '$' && l.next() == '$' {
				l.emit(token.DSTRING)
				return nil
			}
		}
	}

	for r = l.next(); ; r = l.next() {
		if r == '\'' && l.peek() != '\'' {
			l.emit(token.STRING)
			return nil
		}
	}
	return nil
}

func lexUUID(l *lexer) error {
	var (
		r   rune
		pos int = l.relative()
	)
	for ; pos < 36; pos++ {
		r = l.next()
		if pos == 8 || pos == 13 || pos == 18 || pos == 23 {
			if r != '-' {
				return InvalidUUID
			}
			continue
		}
		if !isHex(r) {
			return InvalidUUID
		}
	}
	l.emit(token.UUID)
	return nil
}

func lexBLOB(l *lexer) error {
	l.next()                              // 0
	l.next()                              // X
	l.acceptRun("0123456789abcdefABCDEF") // HEX
	l.emit(token.BLOB)
	return nil
}
