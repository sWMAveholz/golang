package exe04_Parser

type Lexer struct {
	input   string
	tokens  []string
	current int
}

func NewLexer(input string) *Lexer {
	lexer := new(Lexer)
	lexer.input = input
	lexer.current = 0
	lexer.tokens = splitTokens(input)
	return lexer
}

func (l *Lexer) NextToken() string {
	if l.current == len(l.tokens) {
		return ""
	}
	token := l.tokens[l.current]
	l.current++
	return token
}

func splitTokens(input string) []string {
	result := make([]string, 0)
	token := ""
	for i := 0; i < len(input); i++ {
		currentChar := input[i]
		if currentChar == byte(' ') {
			continue // ignore whitespace
		}
		// START OMIT
		switch currentChar {
		// check for terminal
		case byte('&'), byte('|'), byte('!'), byte('('), byte(')'):
			if token != "" {
				result = append(result, token)
				token = ""
			}
			result = append(result, string(currentChar))
			break
		// var assumed
		default:
			token += string(currentChar) // concat var chars
		}
		// END OMIT
	}

	// append last token if exists
	if token != "" {
		result = append(result, token)
	}
	return result
}



