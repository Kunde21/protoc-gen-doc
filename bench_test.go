package gendoc_test

import (
	"testing"

	"github.com/Kunde21/protoc-gen-doc/parser"
	"github.com/Kunde21/protoc-gen-doc/test"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	codeGenRequest, _ := test.MakeCodeGeneratorRequest()

	for i := 0; i < b.N; i++ {
		parser.ParseCodeRequest(codeGenRequest, nil)
	}
}
