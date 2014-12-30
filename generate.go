package cpf

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func Generate() (cpfString string) {
	rand.Seed(time.Now().UTC().UnixNano())
	digits := rand.Perm(9)
	cpf := digits
	cpf = append(cpf, verify(digits, len(cpf)))
	cpf = append(cpf, verify(cpf, len(cpf)))

	for _, c := range cpf {
		cpfString += strconv.Itoa(c)
	}

	return
}

func GeneratePretty() (cpf string) {
	digits := Generate()

	re := regexp.MustCompile(`(?P<1>\d{3})(?P<2>\d{3})(?P<3>\d{3})(?P<4>\d{2})`)

	names := re.SubexpNames()
	formatted := fmt.Sprintf("${%s}.${%s}.${%s}-${%s}",
		names[1], names[2], names[3], names[4])

	cpf = re.ReplaceAllString(digits, formatted)
	return
}
