package json_parser

import "errors"

type Parser struct {
	tokens []*Token
	cursor int
}

func (p *Parser) NewParser(tokens []*Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) parseFromToken(token *Token) (JsonValue, error) {
	if token == nil {
		return nil, errors.New("invalid token")
	}

	switch token.TokenType {
	case String, Number, Boolean, Null:
		return token.Value, nil
	case LeftBrace:
		return p.parseObject()
	default:
		return nil, errors.New("invalid token")
	}
}

func (p *Parser) Parse() (JsonValue, error) {
	token := p.advance()
	return p.parseFromToken(token)
}

func (p *Parser) parseObject() (JsonObject, error) {
	keyToken := p.advance()
	var jsonObject = make(JsonObject)
	var err error

	for keyToken.TokenType != RightBrace {
		if keyToken.TokenType == EOF {
			return nil, errors.New("unexpected end of file")
		}

		if keyToken.TokenType != String {
			return nil, errors.New("unexpected non-key value")
		}

		err = p.consume(Colon, "key must be colon separated")

		if err != nil {
			return nil, err
		}

		valueToken := p.advance()
		jsonObject[keyToken.Value.(string)], err = p.parseFromToken(valueToken)

		if err != nil {
			return nil, err
		}

		err = p.consumeComma(RightBrace)

		if err != nil {
			return nil, err
		}

		keyToken = p.advance()
	}

	return jsonObject, nil

}

func (p *Parser) peek() *Token {
	return p.tokens[p.cursor]
}

func (p *Parser) peekNext() *Token {
	if p.cursor+1 < len(p.tokens) {
		return p.tokens[p.cursor+1]
	}
	return nil
}

func (p *Parser) consume(tokenType TokenType, errMsg string) error {
	if p.peek().TokenType == tokenType {
		p.cursor++
		return nil
	}

	return errors.New(errMsg)
}

func (p *Parser) consumeComma(exception TokenType) error {
	if p.peek().TokenType == Comma {
		p.advance()

		err := p.checkCommaError()

		return err
	}

	if p.peek().TokenType != exception {
		return errors.New("invalid token")
	}

	return nil
}

func (p *Parser) advance() *Token {
	token := p.tokens[p.cursor]

	p.cursor++

	return token
}

func (p *Parser) checkCommaError() error {
	if p.peekNext() != nil && p.peekNext().TokenType != EOF {
		return nil
	}

	return errors.New("invalid token")
}
