package lexer

import "github.com/gossandra/cql/token"

func lexLiteral(l *lexer) error {
	switch l.peek() {
	case '{':
		// TODO try constant -> set/map
		// Try identifier -> UDT
	case '[':
		return lexListLiteral(l)
	case '(':
		return lexTupleLiteral(l)
	}
	return nil
}

func lexMapLiteral(l *lexer) error {
	return nil
}

func lexSetLiteral(l *lexer) error {
	return nil
}

func lexListLiteral(l *lexer) error {
	return nil
}

func lexTupleLiteral(l *lexer) error {
	var (
		err error
		r   rune
	)
	l.acceptToken(token.LPAREN)
	for {
		r = l.next()
		switch {
		case l.acceptToken(token.RPAREN):
			return nil
		case isWhiteSpace(r):
			l.skip()
		case r == ',':
			l.emit(token.COMMA)

		default:
			l.backup()
			if err = lexTerm(l); err != nil {
				return err
			}
		}
	}
}

func lexUDTLiteral(l *lexer) error {
	return nil
}
