package lexer

import (
	"fmt"

	"github.com/gossandra/cql/token"
	"github.com/pkg/errors"
)

// Lexing CREATE TABLE[WE ARE HERE] ... statement
func createTableState(l *lexer) stateFn {
	l.skip()
	l.acceptIToken(token.IfNotExists) // ... IF NOT EXISTS ...
	l.skip()
	return lexTableNameInside(createTableNameState)
}

func createTableNameState(l *lexer) stateFn {
	l.skip()
	if l.acceptToken(token.LPAREN) {
		return createTableDefinitionState
	}
	return l.errorf("unexpected: %s, expect '(' in %v", l.peek(), l.start)

}
func createTableDefinitionState(l *lexer) stateFn {
	var r rune
	for {
		r = l.peek()
		switch {
		case r == ')':
			l.next()
			return tableOptionsState // TODO: TableOptionsState
		case isWhiteSpace(r):
			l.next()
			l.ignore()
			continue
		case r == ',':
			l.next()
			l.emit(token.COMMA)
		case (r == 'P' || r == 'p') && l.acceptToken(token.PrimaryKey):
			if err := lexPKDefinition(l); err != nil { // PrimaryKeyDefState
				return l.errorf("%s", err)
			}
		default:
			if err := lexColumnDescription(l); err != nil { // PrimaryKeyDefState
				return l.errorf("%s", err)
			}
		}

	}
	return nil
}

func tableOptionsState(l *lexer) stateFn {
	l.skip()
	if !l.acceptToken(token.K_WITH) {
		l.skip()
		l.acceptToken(token.SEMICOLON)
		return nil // NO TABLE OPTIONS
	}

	var r rune
	for {
		r = l.peek()
		switch {
		case isWhiteSpace(r):
			l.next()
			l.ignore()
			//continue
		case r == ',':
			l.next()
			l.emit(token.COMMA)
		case r == ';' || r == RuneEOF:
			return nil // END
		case l.acceptToken(token.CompactStorage):
			return nil // TODO
		case l.acceptToken(token.ClusteringOrderBy):
			return nil
			// parse order by
		default:
			return nil
			// parse option
		}
	}
}

func lexColumnDescription(l *lexer) error {
	for {
		if err := lexIdentifier(l); err != nil {
			return errors.Wrap(err, "invalid column name")
		}
		// col type
		l.skip()
		if err := lexIdentifier(l); err != nil {
			return errors.Wrap(err, "invalid column type")
		}

		// Do not depend on order. PK STATIC, STATIC PK ...
		l.skip()
		l.acceptIToken(token.K_STATIC)
		l.skip()
		l.acceptIToken(token.PrimaryKey)
		l.skip()
		l.acceptIToken(token.K_STATIC)
		return nil // TODO: Error
	}
}

// Lex PK definition AFTER "PRIMARY KEY" KEYWORD.
func lexPKDefinition(l *lexer) error {
	var openParen int
	l.skip()
	if !l.acceptToken(token.LPAREN) {
		return fmt.Errorf("invalid PRIMARY KEY definition at")
	}
	openParen++
	for {
		l.skip()
		switch {
		case openParen == 0:
			return nil
		case l.acceptToken(token.LPAREN):
			openParen++
		case l.acceptToken(token.RPAREN):
			openParen--
		case l.acceptToken(token.COMMA):
		default:
			lexIdentifier(l)
		}
	}
	return nil
}
