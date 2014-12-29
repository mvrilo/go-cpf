package cpf

import (
	"strings"
	"testing"
)

func Test_Generate(t *testing.T) {
	data := Generate()

	if len(data) != 11 {
		t.Errorf("cpf length is not right")
	}

	if strings.Contains(data, "-") || strings.Contains(data, ".") {
		t.Errorf("Generate should return non-formatted string")
	}
}

func Test_GeneratePretty(t *testing.T) {
	data := GeneratePretty()

	if len(data) != 14 {
		t.Errorf("cpf length is not right")
	}

	if !strings.Contains(data, "-") || !strings.Contains(data, ".") {
		t.Errorf("GeneratePretty should return formatted string")
	}
}
