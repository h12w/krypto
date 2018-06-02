package main

func scan(expr string) []string {
	var tokens []string
	var token []rune
	for _, c := range expr {
		switch c {
		case '+', '-', '*', '/', '(', ')':
			if len(token) > 0 {
				tokens = append(tokens, string(token))
				token = nil
			}
			tokens = append(tokens, string(c))
		default:
			token = append(token, c)
		}
	}
	if len(token) > 0 {
		tokens = append(tokens, string(token))
		token = nil
	}
	return tokens
}
