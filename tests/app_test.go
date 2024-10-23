package tests

import (
	. "golang_json_parser"
	"os"
	"regexp"
	"testing"
)

func TestJsonParserValido(t *testing.T) {
	testFiles := []string{
		"tests/step1/valid.json",
	}

	for _, f := range testFiles {
		strFromJson, _ := os.ReadFile(f)
		filenameRegex, _ := regexp.Compile("(\\w+).json$")

		teste := filenameRegex.FindAllStringSubmatch(f, -1)

		res := IsJsonValid(string(strFromJson))

		t.Run(teste[0][1], func(t *testing.T) {
			if true != res {
				t.Error(`Json Inválido`)
			}
		})
	}
}
