package json_parser

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

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
	token := new(Token)

	if len(t.Source) == 0 {
		t.tokens = append(t.tokens, token.NewToken(EOF, "EOF"))
		return t.tokens
	}

	for !t.isEOF() {
		t.start = t.cursor
		token = t.getNextToken()

		if token != nil && token.TokenType != Undefined {
			t.tokens = append(t.tokens, token)
		} else if token == nil {
			t.tokens = append(t.tokens, nil)
			break
		} else if token.TokenType == Undefined {
			//pass
		}

	}

	if t.tokens[len(t.tokens)-1] != nil {
		token = new(Token)

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
	//char := *t.advance()
	numberRegex := regexp.MustCompile("\\d")
	stringRegex := regexp.MustCompile("\\w")
	spaceRegex := regexp.MustCompile("\\s+")

	token := new(Token)

	switch char := *t.advance(); {
	case char == "{":
		token = token.NewToken(LeftBrace, char)
	case char == "}":
		token = token.NewToken(RightBrace, char)
	case char == "[":
		token = token.NewToken(LeftBracket, char)
	case char == "]":
		token = token.NewToken(RightBracket, char)
	case char == ":":
		token = token.NewToken(Colon, char)
	case char == ",":
		token = token.NewToken(Comma, char)
	case char == "\n":
		t.line++
	case spaceRegex.MatchString(char):
		token = token.NewToken(Undefined, char)
	case char == `"`:
		token = t.getStringToken()
	case numberRegex.MatchString(char):
		token = t.getNumberToken(numberRegex)
	case stringRegex.MatchString(char):
		token = t.getKeywordToken(stringRegex)
	default:
		token = nil
	}

	return token
}

func (t *Tokenizer) getStringToken() *Token {
	for (!t.isEOF()) && (t.peek() != `"`) {
		t.advance()
	}

	if t.isEOF() {
		log.Println(errors.New("unexpected end of file"))
		return nil
	}

	t.advance()

	return &Token{TokenType: String, Value: t.Source[t.start+1 : t.cursor-1]}

}

func (t *Tokenizer) getKeywordToken(stringRegex *regexp.Regexp) *Token {
	for (!t.isEOF()) && (stringRegex.MatchString(t.peek())) {
		t.advance()
	}

	switch kw := t.Source[t.start:t.cursor]; {
	case kw == "true", kw == "false":
		boolKw, _ := strconv.ParseBool(kw)
		return &Token{TokenType: Boolean, Value: boolKw}
	case kw == "null":
		return &Token{TokenType: Null, Value: nil}
	default:
		return nil
	}
}

func (t *Tokenizer) getNumberToken(numberRegex *regexp.Regexp) *Token {
	for (!t.isEOF()) && (numberRegex.MatchString(t.peek())) {
		t.advance()
	}

	if t.isEOF() {
		return nil
	}

	var intNum JsonValue

	if t.peek() == "." {
		if !numberRegex.MatchString(t.peekNext()) {
			log.Println(errors.New("invalid number"))
			return nil
		}

		t.advance()

		for numberRegex.MatchString(t.peek()) {
			t.advance()
		}

		intNum, _ = strconv.ParseFloat(t.Source[t.start:t.cursor], 32)
	} else {
		intNum, _ = strconv.ParseInt(t.Source[t.start:t.cursor], 10, 32)
	}

	return &Token{TokenType: Number, Value: intNum}

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

func (t *Tokenizer) peekNext() string {
	if t.cursor+1 >= int64(len(t.Source)) {
		return ""
	}

	return t.Source[t.cursor+1 : t.cursor+2]
}
