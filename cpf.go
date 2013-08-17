package cpf

import (
	"errors"
	"strconv"
	"strings"
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
	return check(cpf)
}

func check(cpf []int) (bool, error) {
	if len(cpf) != 11 {
		return false, errors.New("Invalid length")
	}

	var (
		digits []byte
		i      int
	)

	for ; i < len(cpf); i++ {
		digits = strconv.AppendInt(digits, int64(cpf[i]), 10)
	}

	if blackListed(digits) || !checkEach(cpf, 9) || !checkEach(cpf, 10) {
		return false, errors.New("Invalid value")
	}

	return true, nil
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

var blackList = []string{"123456789", "0000",
	"1111", "2222", "3333", "4444", "5555",
	"6666", "7777", "8888", "9999"}

func blackListed(cpf []byte) (res bool) {
	for i := 0; i < len(blackList); i++ {
		if strings.Contains(string(cpf), blackList[i]) {
			res = true
			break
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
