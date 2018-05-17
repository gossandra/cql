/*
Native types docs: http://cassandra.apache.org/doc/latest/cql/definitions.html#constants
*/
package lexer

import (
	"bytes"
	"errors"
	"log"

	"github.com/gossandra/cql/token"
)

var InvalidUUID = errors.New("Invalid UUID")

type parsingState int8

const (
	cuuid parsingState = iota
	cinteger
	cfloat
)

func lexTerm(l *lexer) error {

	log.Print(string(l.input[l.pos:]))
	if err := lexConstant(l); err == nil {
		return nil
	}
	l.reset()

	if err := lexLiteral(l); err == nil {
		return nil
	}
	l.reset()
	log.Print("lex ERROR")
	return errors.New("not a term")

	// lexIdentifier + ()
}

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
	case isHex(r) || r == '-': // UUID or Number
		return lexNumber(l)
	}
	return errors.New("not a constant")
}

const (
	nint uint8 = 1 << iota
	ndec
	nexpint
	nexpdec
)

func lexNumber(l *lexer) error {
	var (
		t          = token.INT
		r          rune
		dashPrefix bool  = l.peek() == '-'
		state      uint8 = nint
	)
	if dashPrefix {
		l.next()
	}

	digits := "0123456789"

	for {
		l.acceptRun(digits)
		r = l.next()
		switch {
		case isNum(r):
			continue
		case r == RuneEOF:
			return ErrorEOF
		case r == '.':
			t = token.FLOAT
			if !(state == nint || state == nexpint) {
				return errors.New("Invalid number")
			}
			l.acceptRun(digits)
			state = state << 1
			continue
		case r == 'e' || r == 'E':
			if state > ndec {
				return errors.New("invalid exponential notation")
			}
			t = token.FLOAT
			l.accept("+-")
			state = nexpint

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
			if r == RuneEOF {
				return ErrorEOF
			}
		}
	}

	for r = l.next(); ; r = l.next() {
		if r == '\'' && l.peek() != '\'' {
			l.emit(token.STRING)
			return nil
		}
		if r == RuneEOF {
			return ErrorEOF
		}
	}
	return errors.New("not a string")
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
		if r == RuneEOF {
			return ErrorEOF
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
