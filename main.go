package jsonParser

func IsJsonValid(jsonString string) bool {
	var isValid bool

	tokenizer := new(Tokenizer)
	tokenizer = tokenizer.NewTokenizer(jsonString)

	tokens := tokenizer.Scan()

	if tokens[len(tokens)-1] != nil {
		isValid = true
	}

	return isValid
}
