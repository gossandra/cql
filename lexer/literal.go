package lexer

import (
	"errors"
	"log"

	"github.com/gossandra/cql/token"
)

func lexLiteral(l *lexer) error {
	switch l.peek() {
	case '{':
		return lexSetMapUDT(l)
	case '[', '(':
		return lexListTupleLiteral(l)
	}
	return errors.New("Not a literal")
}

func lexListTupleLiteral(l *lexer) error {
	var (
		err  error
		r    rune
		last token.Token = token.RBRACK
	)
	if l.acceptToken(token.LPAREN) {
		last = token.RPAREN
	}
	for {
		r = l.next()
		switch {
		case l.acceptToken(last):
			return nil
		case isWhiteSpace(r):
			l.skip()
		case r == ',':
			l.emit(token.COMMA)
		case r == RuneEOF:
			return ErrorEOF

		default:
			l.backup()
			if err = lexTerm(l); err != nil {
				return err
			}
		}
	}
}

func lexSetMapUDT(l *lexer) error {
	l.acceptToken(token.LBRACE)
	if err := lexConstant(l); err == nil {
		return lexSetMapLiteral(l)
	}
	if err := lexIdentifier(l); err == nil {
		return lexUDTLiteral(l)
	}

	return errors.New("literal syntax error")
}

func lexUDTLiteral(l *lexer) error {
	afterColon := false
	for {
		switch l.peek() {
		case '}':
			l.acceptToken(token.RBRACE)
			return nil
		case ',':
			l.acceptToken(token.COMMA)
			afterColon = false
		case ':':
			l.acceptToken(token.COLON)
			afterColon = true
		case RuneEOF:
			return ErrorEOF
		default:
		}
		if afterColon {
			if err := lexTerm(l); err != nil {
				return err
			}
		} else {
			if err := lexIdentifier(l); err != nil {
				return err
			}

		}
	}
	return nil
}

const (
	lkey int8 = 1 << iota
	lcolon
	lvalue
	lcomma
)

func lexSetMapLiteral(l *lexer) error {
	var (
		isMap bool = false
		state int8 = lkey
	)
	for {
		switch l.peek() {
		case ' ', '\n', '\t':
			l.skip()
		case '}':
			l.acceptToken(token.RBRACE)
			return nil
		case ',':
			if state != lcomma {
				return errors.New("unexpected comma")
			}
			state = lkey
			l.acceptToken(token.COMMA)
			continue
		case ':':
			isMap = true
			state = lvalue
			l.acceptToken(token.COLON)
			continue
		case RuneEOF:
			return ErrorEOF
		default:
		}

		log.Print(string(l.input[l.pos:]))
		if err := lexTerm(l); err != nil {
			return err
		}

		if state == lvalue || !isMap {
			state = lcomma
		} else {
			state = lcolon
		}
	}
}
