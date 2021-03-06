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

func TestLexTerm(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		hasErr bool
	}{
		{"BLOB", "0xabcdefababcdabcdabcd0123456789012", false},
		{"BLOB", "0xABCDEFABABCDABCDABCD0123456789012", false},
		{"BLOB", "0XABCDEFABABCDABCDABCD0123456789012", false},

		{"UUID", "abcdefab-abcd-abcd-abcd-0123456789012,", false},
		{"UUID incorrect", "abcdefab0abcd-abcd-abcd-0123456789012,", true},
		{"UUID incorrect", "-bcdefab-abcd-abcd-abcd-0123456789012,", true},

		{"Float", "-1795.65734E+17.472 ", false},
		{"Float", "-1795..65734E+17.472 ", true},
		{"Float", "-1795.65734EE+17.472 ", true},
		{"Float", "+17.65", true},
		{"Integer", "1488}", false},
		{"Integer", "77", false},

		{"String", "'somestring '' with doeblequoted escape'", false},
		{"String dollar quoted", "$$ somestring '\"!@#$%^&*() $$, ", false},

		{"Map<string, string>", "{'key1': 'value1', 'key2': 'value2'}", false},
		{"Map<string, string>", "{'key1': ,'value1', 'key2': 'value2'}", true},
		{"Map<int, float>", "{42: 36.6, 33: 777.734, -6: -99.3e+55.3}", false},
		{"Map<int, float>", "{43: 36..6, 33: 777.734, -6: -99.3e+55.3}", true},
		{"Map with bind args", "{'key1': ?, 'key2': ?}", false},
		{"Map with bind named args", "{'key1': :ke1_val, 'key2': :key2_val}", false},
		{"Unclosed inner map", "{44: {36.6, 33: 777.734, -6: -99.3e+55.3}", true},
		{"Unclosed inner map with syntax error", "{44: {36..6, 33: 777.734, -6: -99.3e+55.3}", true},
		{"Arithmetic add", "16.24 + 74", false},

		{"Function", "somefunc()", false},
		{"Function with args", "somefunc(false, 17)", false},
		{"Function with invalid args", "somefunc(}, 17)", true},
		{"Function with math args", "somefunc(17 + 44, 17.37)", false},
		{"Function with quoted identifier", "\"somefunc\"(17 + 44, 17.37)", false},
		{"Function with quoted identifier escape", `"some""func"(17 + 44, 17.37)`, false},
		{"Function with bind args", "somefunc(?, ?, ?)", false},
		{"Function with named bind args", "somefunc(:arg1, :arg2, true)", false},

		{"UDT", `{Data: 0xAFFCAC6C, Name: 'stringvalue'}`, false},
		{"UDT", `{Data: 0xAFFCAC6C, Age: 21}`, false},
		{"UDT", `{Data: 0xAFFCAC6C, Name: 'stringvalue'} `, false},
	}

	for _, r := range tt {
		t.Run(r.name, func(t *testing.T) {
			l := New([]byte(r.input))
			// Log items chan
			go logItems(l.items)

			err := lexTerm(l)
			if (err != nil) != r.hasErr {
				t.Log(err, l.pos)
				t.Fail()
			}
		})

	}
}

func logItems(items <-chan Item) {
	for i := range items {
		log.Print("ITEM: ", i)
	}
}
