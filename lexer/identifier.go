package lexer

import (
	"fmt"

	"github.com/gossandra/cql/token"
)

func lexIdentifier(l *lexer) error {
	var r rune
	if r = l.peek(); !isAlpha(r) {
		return lexQuotedIdentifier(l)
	}
	l.next()
	for {
		r = l.peek()
		if !isAlphaNum(r) && r != '_' {
			l.emit(token.IDENT)
			return nil
		}
		l.next()
	}
}

// Helper
func lexQuotedIdentifier(l *lexer) error {
	var r rune
	if r = l.next(); r != '"' { // FirstQuote
		return fmt.Errorf("Syntax error at: %s", l.start)
	}

	for {
		r = l.next()
		if r == '"' {
			if l.peek() == '"' { // DoubleQuoted
				l.next()
				continue
			}
			l.emit(token.QIDENT)
			return nil
		}

	}

}

func lexTableNameInside(next stateFn) stateFn {
	return func(l *lexer) stateFn {
		if err := lexIdentifier(l); err != nil {
			return l.errorf("invalid table name: %v", err)
		}

		if l.peek() == '.' {
			l.next()
			l.emit(token.PERIOD)
			if err := lexIdentifier(l); err != nil {
				return l.errorf("invalid table name: %v", err)

			}

		}
		return next
	}
}

// TODO: Check to EOF/END
