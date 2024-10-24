package jsonParser

type TokenType int

const (
	Undefined TokenType = iota
	LeftBrace
	RightBrace
	LeftBracket
	RightBracket
	Comma
	Colon
	EOF
)

func (t TokenType) String() string {
	switch t {
	case RightBrace:
		return "RIGHT_BRACE"
	case LeftBrace:
		return "LEFT_BRACE"
	case RightBracket:
		return "RIGHT_BRACKET"
	case LeftBracket:
		return "LEFT_BRACKET"
	case Comma:
		return "COMMA"
	case Colon:
		return "COLOM"
	case EOF:
		return "EOF"
	default:
		return "Unknown"
	}
}

func (t TokenType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}
