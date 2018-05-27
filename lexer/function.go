package lexer

import "github.com/gossandra/cql/token"

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
	l.acceptToken(token.LPAREN)
	return lexFunctionParams(l)

}
