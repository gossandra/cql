package lexer

// stateFn represents current state, and returns next state according to query
type stateFn func(*lexer) stateFn

var (
	startState stateFn = rootState
)

func lex(input []byte) (*lexer, chan Item) {
	l := lexer{
		input: input,
		items: make(chan Item, 10),
	}

	go l.run()
	return &l, l.items
}
