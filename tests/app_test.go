package tests

import (
	. "golang_json_parser/json_parser"
	"os"
	"regexp"
	"testing"
)

func TestJsonParserValido(t *testing.T) {
	testFiles := []string{
		"step1/valid.json",
		"step2/valid.json",
		"step2/valid2.json",
		"step3/valid.json",
		"step4/valid.json",
		"step4/valid2.json",
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

func TestJsonParserInvalido(t *testing.T) {
	testFiles := []string{
		"step1/invalid.json",
		"step2/invalid.json",
		"step2/invalid2.json",
		"step3/invalid.json",
		"step4/invalid.json",
	}

	for _, f := range testFiles {
		strFromJson, _ := os.ReadFile(f)
		filenameRegex, _ := regexp.Compile("(\\w+).json$")

		teste := filenameRegex.FindAllStringSubmatch(f, -1)

		res := IsJsonValid(string(strFromJson))

		t.Run(teste[0][1], func(t *testing.T) {
			if false != res {
				t.Error(`Json Válido`)
			}
		})
	}
}
