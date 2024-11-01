package main

import (
	"encoding/json"
	"fmt"
	"golang_json_parser/json_parser"
	"log"
)

func main() {
	strToTokenize := `{
  "key-l": ['list value']
}`

	tokenizer := new(json_parser.Tokenizer)

	tokenizer = tokenizer.NewTokenizer(strToTokenize)

	tokens := tokenizer.Scan()

	parser := new(json_parser.Parser)
	parser = parser.NewParser(tokens)

	values, err := parser.Parse()

	if err != nil {
		log.Println(err.Error())
	} else {
		dump(values)
	}
}

func dump(data interface{}) {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Print(string(b))
}
