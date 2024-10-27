package json_parser

func IsJsonValid(jsonString string) bool {
	var isValid bool

	tokenizer := new(Tokenizer)
	tokenizer = tokenizer.NewTokenizer(jsonString)

	tokens := tokenizer.Scan()

	parser := new(Parser)

	parser = parser.NewParser(tokens)
	_, err := parser.Parse()

	if err == nil {
		isValid = true
	}
	return isValid
}
