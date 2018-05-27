package lexer

import (
	"errors"

	"github.com/gossandra/cql/token"
)

var NotAFunctionError = errors.New("Identifier is not a function call")

func lexFunctionParams(l *lexer) error {
	for {
		l.skip()
		switch l.peek() {
		case ')':
			l.acceptToken(token.RPAREN)
			return nil
		case ',':
			l.acceptToken(token.COMMA)
		default:
			if err := lexTerm(l); err != nil {
				return err
			}
		}

	}
}

func lexFunction(l *lexer) error {
	var (
		err error
	)
	err = lexIdentifier(l)
	if err != nil {
		return err
	}
	if !l.acceptToken(token.LPAREN) {
		return NotAFunctionError
	}
	return lexFunctionParams(l)

}
