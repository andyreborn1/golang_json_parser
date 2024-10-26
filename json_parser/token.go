package app

type Token struct {
	TokenType TokenType
	Value     string
}

func (token Token) NewToken(tokenType TokenType, value string) *Token {
	return &Token{TokenType: tokenType, Value: value}
}
