/*
List of possible CQL3 syntax tokens

Source: Golang parser


TODO: Check literal part and operators

*/

package token

// Token is the set of lexical tokens of the CQL query language.
type Token int

func (t Token) String() string {
	return tokens[int(t)]
	//if str, ok := tokens[t]; ok {
	//return str
	//}
	//return "Token: " + strconv.Itoa(t)

}

// The list of tokens and keywords
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT   // fieldname
	QIDENT  // "fieldName"
	INT     // 12345
	FLOAT   // 123.45
	IMAG    // 123.45i
	STRING  // 'abc'
	DSTRING // $$abc$$
	BLOB    // 0x
	UUID    // 8-4-4-4-12
	literal_end

	operator_beg
	// Operators and delimiters
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=

	LAND  // &&
	LOR   // ||
	ARROW // <-
	INC   // ++
	DEC   // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	DEFINE   // :=
	ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	BLOBPrefix // 0x
	operator_end

	keyword_beg
	// CQL Keywords

	// --> COMBINED
	IfNotExists
	PrimaryKey
	CompactStorage
	ClusteringOrderBy

	// Not a CQL
	K_NULL
	K_TRUE
	K_FALSE
	//
	// --> ALL
	K_ADD
	K_AGGREGATE
	K_ALL
	K_ALLOW
	K_ALTER
	K_AND
	K_ANY
	K_APPLY
	K_AS
	K_ASC
	K_ASCII
	K_AUTHORIZE
	K_BATCH
	K_BEGIN
	K_BIGINT
	K_BLOB
	K_BOOLEAN
	K_BY
	K_CLUSTERING
	K_COLUMNFAMILY
	K_COMPACT
	K_CONSISTENCY
	K_COUNT
	K_COUNTER
	K_CREATE
	K_CUSTOM
	K_DECIMAL
	K_DELETE
	K_DESC
	K_DISTINCT
	K_DOUBLE
	K_DROP
	K_EACH_QUORUM
	K_ENTRIES
	K_EXISTS
	K_FILTERING
	K_FLOAT
	K_FROM
	K_FROZEN
	K_FULL
	K_GRANT
	K_IF
	K_IN
	K_INDEX
	K_INET
	K_INFINITY
	K_INSERT
	K_INT
	K_INTO
	K_KEY
	K_KEYSPACE
	K_KEYSPACES
	K_LEVEL
	K_LIMIT
	K_LIST
	K_LOCAL_ONE
	K_LOCAL_QUORUM
	K_MAP
	K_MATERIALIZED
	K_MODIFY
	K_NAN
	K_NORECURSIVE
	K_NOSUPERUSER
	K_NOT
	K_OF
	K_ON
	K_ONE
	K_ORDER
	K_PARTITION
	K_PASSWORD
	K_PER
	K_PERMISSION
	K_PERMISSIONS
	K_PRIMARY
	K_QUORUM
	K_RENAME
	K_REVOKE
	K_SCHEMA
	K_SELECT
	K_SET
	K_STATIC
	K_STORAGE
	K_SUPERUSER
	K_TABLE
	K_TEXT
	K_TIME
	K_TIMESTAMP
	K_TIMEUUID
	K_THREE
	K_TO
	K_TOKEN
	K_TRUNCATE
	K_TTL
	K_TUPLE
	K_TWO
	K_TYPE
	K_UNLOGGED
	K_UPDATE
	K_USE
	K_USER
	K_USERS
	K_USING
	K_UUID
	K_VALUES
	K_VARCHAR
	K_VARINT
	K_VIEW
	K_WHERE
	K_WITH
	K_WRITETIME

	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:   "IDENT",
	QIDENT:  "QUOTED IDENT",
	INT:     "INT",
	FLOAT:   "FLOAT",
	IMAG:    "IMAG",
	STRING:  "STRING",
	DSTRING: "DOLLAR STRING",
	BLOB:    "BLOB",
	UUID:    "UUID",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

	LAND:  "&&",
	LOR:   "||",
	ARROW: "<-",
	INC:   "++",
	DEC:   "--",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ:      "!=",
	LEQ:      "<=",
	GEQ:      ">=",
	DEFINE:   ":=",
	ELLIPSIS: "...",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:     ")",
	RBRACK:     "]",
	RBRACE:     "}",
	SEMICOLON:  ";",
	COLON:      ":",
	BLOBPrefix: "0X",

	// Keywords
	// --> COMBINED

	IfNotExists:       "IF NOT EXISTS",
	PrimaryKey:        "PRIMARY KEY",
	CompactStorage:    "COMPACT STORAGE",
	ClusteringOrderBy: "CLUSTERING ORDER BY",
	// --> ALL

	K_NULL:         "NULL",
	K_TRUE:         "TRUE",
	K_FALSE:        "FALSE",
	K_ADD:          "ADD",
	K_AGGREGATE:    "AGGREGATE",
	K_ALL:          "ALL",
	K_ALLOW:        "ALLOW",
	K_ALTER:        "ALTER",
	K_AND:          "AND",
	K_ANY:          "ANY",
	K_APPLY:        "APPLY",
	K_AS:           "AS",
	K_ASC:          "ASC",
	K_ASCII:        "ASCII",
	K_AUTHORIZE:    "AUTHORIZE",
	K_BATCH:        "BATCH",
	K_BEGIN:        "BEGIN",
	K_BIGINT:       "BIGINT",
	K_BLOB:         "BLOB",
	K_BOOLEAN:      "BOOLEAN",
	K_BY:           "BY",
	K_CLUSTERING:   "CLUSTERING",
	K_COLUMNFAMILY: "COLUMNFAMILY",
	K_COMPACT:      "COMPACT",
	K_CONSISTENCY:  "CONSISTENCY",
	K_COUNT:        "COUNT",
	K_COUNTER:      "COUNTER",
	K_CREATE:       "CREATE",
	K_CUSTOM:       "CUSTOM",
	K_DECIMAL:      "DECIMAL",
	K_DELETE:       "DELETE",
	K_DESC:         "DESC",
	K_DISTINCT:     "DISTINCT",
	K_DOUBLE:       "DOUBLE",
	K_DROP:         "DROP",
	K_EACH_QUORUM:  "EACH_QUORUM",
	K_ENTRIES:      "ENTRIES",
	K_EXISTS:       "EXISTS",
	K_FILTERING:    "FILTERING",
	K_FLOAT:        "FLOAT",
	K_FROM:         "FROM",
	K_FROZEN:       "FROZEN",
	K_FULL:         "FULL",
	K_GRANT:        "GRANT",
	K_IF:           "IF",
	K_IN:           "IN",
	K_INDEX:        "INDEX",
	K_INET:         "INET",
	K_INFINITY:     "INFINITY",
	K_INSERT:       "INSERT",
	K_INT:          "INT",
	K_INTO:         "INTO",
	K_KEY:          "KEY",
	K_KEYSPACE:     "KEYSPACE",
	K_KEYSPACES:    "KEYSPACES",
	K_LEVEL:        "LEVEL",
	K_LIMIT:        "LIMIT",
	K_LIST:         "LIST",
	K_LOCAL_ONE:    "LOCAL_ONE",
	K_LOCAL_QUORUM: "LOCAL_QUORUM",
	K_MAP:          "MAP",
	K_MATERIALIZED: "MATERIALIZED",
	K_MODIFY:       "MODIFY",
	K_NAN:          "NAN",
	K_NORECURSIVE:  "NORECURSIVE",
	K_NOSUPERUSER:  "NOSUPERUSER",
	K_NOT:          "NOT",
	K_OF:           "OF",
	K_ON:           "ON",
	K_ONE:          "ONE",
	K_ORDER:        "ORDER",
	K_PARTITION:    "PARTITION",
	K_PASSWORD:     "PASSWORD",
	K_PER:          "PER",
	K_PERMISSION:   "PERMISSION",
	K_PERMISSIONS:  "PERMISSIONS",
	K_PRIMARY:      "PRIMARY",
	K_QUORUM:       "QUORUM",
	K_RENAME:       "RENAME",
	K_REVOKE:       "REVOKE",
	K_SCHEMA:       "SCHEMA",
	K_SELECT:       "SELECT",
	K_SET:          "SET",
	K_STATIC:       "STATIC",
	K_STORAGE:      "STORAGE",
	K_SUPERUSER:    "SUPERUSER",
	K_TABLE:        "TABLE",
	K_TEXT:         "TEXT",
	K_TIME:         "TIME",
	K_TIMESTAMP:    "TIMESTAMP",
	K_TIMEUUID:     "TIMEUUID",
	K_THREE:        "THREE",
	K_TO:           "TO",
	K_TOKEN:        "TOKEN",
	K_TRUNCATE:     "TRUNCATE",
	K_TTL:          "TTL",
	K_TUPLE:        "TUPLE",
	K_TWO:          "TWO",
	K_TYPE:         "TYPE",
	K_UNLOGGED:     "UNLOGGED",
	K_UPDATE:       "UPDATE",
	K_USE:          "USE",
	K_USER:         "USER",
	K_USERS:        "USERS",
	K_USING:        "USING",
	K_UUID:         "UUID",
	K_VALUES:       "VALUES",
	K_VARCHAR:      "VARCHAR",
	K_VARINT:       "VARINT",
	K_VIEW:         "VIEW",
	K_WHERE:        "WHERE",
	K_WITH:         "WITH",
	K_WRITETIME:    "WRITETIME",
}
