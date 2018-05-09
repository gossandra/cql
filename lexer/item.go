package lexer

import (
	"fmt"

	"github.com/gossandra/cql/token"
)

type Item struct {
	Token token.Token
	Value []byte
}

func (i Item) String() string {
	switch i.Token {
	case token.EOF:
		return "EOF"
		// TODO
	}
	if len(i.Value) == 0 {
		return i.Token.String()
	}

	return fmt.Sprintf("%q", i.Value)
}
