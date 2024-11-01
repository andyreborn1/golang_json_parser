package json_parser

type TokenType int

type JsonValue = any

type JsonObject = map[string]JsonValue

type JsonArray = []JsonValue

const (
	Undefined TokenType = iota
	LeftBrace
	RightBrace
	LeftBracket
	RightBracket
	String
	Number
	Boolean
	Null
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
	case String:
		return "STRING"
	case Number:
		return "NUMBER"
	case Boolean:
		return "BOOLEAN"
	case Null:
		return "NULL"
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
