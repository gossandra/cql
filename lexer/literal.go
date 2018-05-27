package lexer

import (
	"errors"
	"fmt"
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
	var err error
	l.acceptToken(token.LBRACE)
	if l.acceptToken(token.RBRACE) {
		return nil // empty collection
	}
	err = lexTerm(l)

	if err == nil {
		return lexSetMapLiteral(l)
	}
	if err == NotAFunctionError {
		return lexUDTLiteral(l)
	}

	return errors.New("literal syntax error")
}

func lexUDTLiteral(l *lexer) error {
	log.Print("UDT LITERAL")
	//l.reset()
	// starts with colon, because {FirstField is lexed by `lexSetMapUDT`
	for {
		l.skip()
		if l.peek() == RuneEOF {
			return ErrorEOF
		}

		// colon:
		l.skip()
		if !l.acceptToken(token.COLON) {
			return fmt.Errorf("unexpected %s, expecting COLON", string(l.peek()))
		}

		// value:
		l.skip()
		if err := lexTerm(l); err != nil {
			return err
		}

		// comma:
		l.skip()
		if l.acceptToken(token.RBRACE) {
			return nil
		}
		if !l.acceptToken(token.COMMA) {
			return fmt.Errorf("unexpected %s, expecting COMMA", string(l.peek()))
		}
		// key:
		l.skip()
		if err := lexIdentifier(l); err != nil {
			return err
		}

	}
}

func lexSetMapLiteral(l *lexer) error {
	var (
		isMap bool = false
	)

	l.skip()
	if l.acceptToken(token.COLON) {
		isMap = true
	}

	for {
		l.skip()
		// MAP part
		if isMap {

			// value:
			l.skip()
			if err := lexTerm(l); err != nil {
				return err
			}
		}

		// comma:
		l.skip()
		if l.acceptToken(token.RBRACE) {
			return nil
		}
		if !l.acceptToken(token.COMMA) {
			return fmt.Errorf("unexpected %s, expecting COMMA", string(l.peek()))
		}

		// key
		if err := lexTerm(l); err != nil {
			return err
		}

		if isMap {
			// colon:
			l.skip()
			if !l.acceptToken(token.COLON) {
				return fmt.Errorf("unexpected %s, expecting COLON", string(l.peek()))
			}
		}

	}
}
