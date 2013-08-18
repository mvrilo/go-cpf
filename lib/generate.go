package cpf

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func Generate() []int {
	rand.Seed(time.Now().UTC().UnixNano())
	var (
		digits = rand.Perm(9)
		cpf    = digits
	)
	cpf = append(cpf, verify(digits))
	cpf = append(cpf, verify(digits))
	return cpf
}

func GenerateStr() string {
	var (
		digits = Generate()
		cpf    []byte
		i      int
	)

	for ; i < len(digits); i++ {
		cpf = strconv.AppendInt(cpf, int64(digits[i]), 10)
	}

	return string(cpf)
}

func GeneratePretty() string {
	var (
		digits    = GenerateStr()
		re        = regexp.MustCompile(`(?P<1>\d{3})(?P<2>\d{3})(?P<3>\d{3})(?P<4>\d{2})`)
		formatted = fmt.Sprintf("${%s}.${%s}.${%s}-${%s}", re.SubexpNames()[1], re.SubexpNames()[2], re.SubexpNames()[3], re.SubexpNames()[4])
		cpf       = re.ReplaceAllString(digits, formatted)
	)

	return cpf
}

func verify(digits []int) int {
	var (
		total int
		i     int
	)

	for ; i < len(digits); i++ {
		total += digits[i] * (10 - i)
	}

	total = total % 11
	if total < 2 {
		return 0
	}
	return 11 - total
}
