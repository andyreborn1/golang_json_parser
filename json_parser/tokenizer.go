package app

type Tokenizer struct {
	Source string
	cursor int64
	start  int64
	line   int64
	tokens []*Token
}

func (t *Tokenizer) NewTokenizer(str string) *Tokenizer {
	t.line = 1
	return &Tokenizer{Source: str}
}

func (t *Tokenizer) Scan() []*Token {
	if len(t.Source) == 0 {
		return []*Token{nil}
	}

	for !t.isEOF() {
		t.start = t.cursor
		token := t.getNextToken()

		if token != nil {
			t.tokens = append(t.tokens, token)
		} else {
			t.tokens = append(t.tokens, nil)
			break
		}

	}

	if t.tokens[len(t.tokens)-1] != nil {
		token := new(Token)

		t.tokens = append(t.tokens, token.NewToken(EOF, "EOF"))
	}

	return t.tokens
}

func (t *Tokenizer) hasMoreTokens() bool {
	return t.cursor < int64(len(t.Source))
}

func (t *Tokenizer) isEOF() bool {
	return t.cursor == int64(len(t.Source))
}

func (t *Tokenizer) getNextToken() *Token {
	char := *t.advance()

	token := new(Token)

	switch char {
	case "{":
		token = token.NewToken(LeftBrace, char)
	case "}":
		token = token.NewToken(RightBrace, char)
	case "[":
		token = token.NewToken(LeftBracket, char)
	case "]":
		token = token.NewToken(RightBracket, char)
	case ",":
		token = token.NewToken(Comma, char)
	case ":":
		token = token.NewToken(Colon, char)
	case "\n":
		t.line++
	case `"`:
		token = token.NewToken(String, char)
	default:
		token = nil
	}

	return token
}

func (t *Tokenizer) addString() {
	for (!t.isEOF()) && (t.peek() != `"`) {
		t.advance()
	}

	if t.isEOF() {
		t.tokens = append(t.tokens, nil)
		return
	}

	t.advance()

	t.tokens = append(t.tokens, &Token{TokenType: String})
}

func (t *Tokenizer) advance() *string {
	var char string

	if t.hasMoreTokens() {
		char = t.Source[t.cursor : t.cursor+1]
	} else {
		return nil
	}

	t.cursor++

	return &char
}

func (t *Tokenizer) peek() string {
	if t.isEOF() {
		return ""
	}

	return t.Source[t.cursor : t.cursor+1]
}
