package cpf

import (
	"strings"
	"testing"
)

func Test_sanitize(t *testing.T) {
	in := GeneratePretty()
	out := sanitize(in)
	if strings.Index(out, ".") > -1 || strings.Index(out, "-") > -1 {
		t.Errorf("sanitize didn't work")
	}
}

func Test_stringToIntSlice(t *testing.T) {
	stringToIntSlice_helper(t, "01234", []int{0, 1, 2, 3, 4})
	stringToIntSlice_helper(t, "01-23.4", []int{0, 1, 2, 3, 4})
	stringToIntSlice_helper(t, "724.806.351-71", []int{7, 2, 4, 8, 0, 6, 3, 5, 1, 7, 1})
}

func stringToIntSlice_helper(t *testing.T, in string, out []int) {
	b := stringToIntSlice(in)
	if len(b) != len(out) {
		t.Errorf("length is not right")
	}
	for i, digit := range out {
		if b[i] != digit {
			t.Errorf("comparison failed")
		}
	}
}

func Test_Valid(t *testing.T) {
	for _, d := range strings.Split(blacklist, "\n") {
		if ok, _ := Valid(d); ok {
			t.Errorf("blacklist values shouldn't be valid")
		}
	}
}
