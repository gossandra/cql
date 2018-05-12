package lexer

func isAlpha(r rune) bool {
	if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') {
		return false
	}
	return true
}

func isNum(r rune) bool {
	if r < '0' || r > '9' {
		return false
	}
	return true
}

func isAlphaNum(r rune) bool {
	if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && (r < '0' || r > '9') {
		return false
	}
	return true
}

func isSpace(r rune) bool {
	return r == ' '
}

func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}
