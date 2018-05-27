package lexer

import (
	"errors"

	"github.com/gossandra/cql/token"
)

var NotAMarkerError = errors.New("Not a marker")

func lexBindMarker(l *lexer) error {
	switch l.next() {
	case ':': // named marker
		l.emit(token.NAMED_MARKER)
		return lexIdentifier(l)
	case '?':
		l.emit(token.MARKER)
		return nil
	default:
		l.backup()
		return NotAMarkerError
	}
}
