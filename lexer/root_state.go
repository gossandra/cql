package lexer

import (
	"github.com/gossandra/cql/token"
)

func rootState(l *lexer) stateFn {
	l.skip()

	switch {
	case l.acceptIToken(token.K_BATCH):
		return nil
	// DDL
	case l.acceptIToken(token.K_ALTER):
		return nil
	case l.acceptIToken(token.K_CREATE):
		return createState
	case l.acceptIToken(token.K_DROP):
		return nil
	case l.acceptIToken(token.K_GRANT):
		return nil
	case l.acceptIToken(token.K_REVOKE):
		return nil

	// DML
	case l.acceptIToken(token.K_SELECT):
		return nil
	case l.acceptIToken(token.K_INSERT):
		return nil
	case l.acceptIToken(token.K_UPDATE):
		return nil
	case l.acceptIToken(token.K_DELETE):
		return nil

	default:
		l.items <- Item{Token: token.ILLEGAL, Value: []byte("Unexpected: error ")}
		return nil
	}
}

func createState(l *lexer) stateFn {
	l.skip()
	switch {
	case l.acceptIToken(token.K_AGGREGATE):
		return nil
	// DDL
	case l.acceptIToken(token.K_INDEX):
		return nil
	//case l.acceptKeyword(token.K_FUNCTION):
	//return createState
	case l.acceptIToken(token.K_KEYSPACE):
		return nil
	case l.acceptIToken(token.K_TABLE):
		return createTableState
	case l.acceptIToken(token.K_REVOKE):
		return nil

		//l.items <- Item{Token: token.K_TABLE}

	}
	return nil
}
