package cpf

import (
	"errors"
	"fmt"
	"strconv"
)

func Valid(digits interface{}) bool {
	ret, _ := Check(digits)
	return ret
}

func Check(digits interface{}) (bool, error) {
	cpf, err := fixType(digits)
	if err != nil {
		return false, err
	}

	if len(cpf) != 11 {
		return false, errors.New("Length exceeded")
	}

	if fmt.Sprintf("%v", cpf) == fmt.Sprintf("%v", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}) {
		return false, errors.New("Invalid value")
	}

	return !checkEqual(cpf) && checkEach(cpf, 9) && checkEach(cpf, 10), nil
}

func fixType(digits interface{}) ([]int, error) {
	switch v := digits.(type) {
	case string:
		return fixType([]byte(v))
	case []byte:
		var slice = []int{}
		for i, _ := range v {
			fix, _ := strconv.Atoi(string(v[i]))
			slice = append(slice, fix)
		}
		return fixType(slice)
	case []int:
		return v, nil
	default:
		return nil, errors.New("Type not supported")
	}
}

func checkEqual(cpf []int) (res bool) {
	var (
		i int
		j = 1
	)

	for ; i < len(cpf); i++ {
		if i == len(cpf)-1 {
			j = 0
		}

		if res || i == 0 {
			res = cpf[i] == cpf[i+j]
		}
	}

	return
}

func checkEach(cpf []int, n int) bool {
	var (
		i     int
		sum   int
		final int
	)

	for ; i < n; i++ {
		sum += cpf[i] * (n + 1 - i)
	}

	final = sum % 11
	if final < 2 {
		final = 0
	} else {
		final = 11 - final
	}

	return final == cpf[n]
}
