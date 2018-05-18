package lexer

import (
	"errors"

	"github.com/gossandra/cql/token"
)

func lexArithmetic(l *lexer) error {
	l.skip()

	switch l.peek() {
	case RuneEOF:
		return ErrorEOF
	case '+':
		l.acceptToken(token.ADD)
	case '-':
		l.acceptToken(token.SUB)
	case '*':
		l.acceptToken(token.MUL)
	case '/':
		l.acceptToken(token.QUO)
	case '%':
		l.acceptToken(token.REM)
	default:
		return errors.New("Not an Arithmetic operation")
	}
	return nil
}
