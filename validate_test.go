package cpf

import (
	"strings"
	"testing"
)

func Test_sanitize(t *testing.T) {
	in := "215.830.467-71"
	out := sanitize(in)
	if strings.Index(out, ".") > -1 || strings.Index(out, "-") > -1 {
		t.Errorf("sanitize didn't work")
	}
}
