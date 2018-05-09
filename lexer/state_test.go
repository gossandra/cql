package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	_, items := lex([]byte(`CREATE TABLE IF NOT EXISTS keyspace1.test_table1 (
		col1 Text,
		col2 Int,
		col3 UUID,
		PRIMARY KEY (col1, col2)
	);`))

	for item := range items {
		t.Log(item)
	}
}
