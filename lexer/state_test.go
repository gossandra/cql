package lexer

import (
	"log"
	"testing"
)

var q1 = `SELECT name, occupation FROM users WHERE userid IN (199, 200, 207);
SELECT JSON name, occupation FROM users WHERE userid = 199;
SELECT name AS user_name, occupation AS user_occupation FROM users;

SELECT time, value
FROM events
WHERE event_type = 'myEvent'
  AND time > '2011-02-03'
  AND time <= '2012-01-01'

SELECT COUNT (*) AS user_count FROM users;`

var q0 = `CREATE TABLE IF NOT EXISTS keyspace1.test_table1 (
		col1 Text,
		col2 Int static,
		col3 UUID,
		PRIMARY KEY (col1, col2)
	) WITH CLUSTERING ORDER BY ololo DESC;`

func TestLexer(t *testing.T) {
	var queries = []string{
		q0,
		q1,
	}

	for _, q := range queries {
		_, items := Lex([]byte(q))

		for item := range items {
			t.Log(item)
		}
		t.Log("\n\n====================")
	}
}

func TestLexConstant(t *testing.T) {

	tt := []struct {
		name   string
		input  string
		hasErr bool
	}{
		{"UUID", "abcdefab-abcd-abcd-abcd-0123456789012,", false},
		{"UUID incorrect", "abcdefab0abcd-abcd-abcd-0123456789012,", true},
		{"UUID incorrect", "-bcdefab-abcd-abcd-abcd-0123456789012,", true},
		{"Float", "-1795.65734E+17.472 ", false},
		{"Integer", "1488}", false},
		{"String", "'somestring '' with doeblequoted escape'", false},
		{"String dollar quoted", "$$ somestring '\"!@#$%^&*() $$, ", false},
	}

	for _, r := range tt {
		t.Run(r.name, func(t *testing.T) {
			l := New([]byte(r.input))
			// Log items chan
			go logItems(l.items)

			err := lexConstant(l)
			if (err != nil) != r.hasErr {
				t.Log(err, l.pos)
				t.Fail()
			}
		})

	}
}

func logItems(items <-chan Item) {
	for i := range items {
		log.Print(i)
	}
}
