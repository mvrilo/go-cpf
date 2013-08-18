package cpf

import (
	"errors"
	"fmt"
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

	if blackListed(cpf) || !checkEach(cpf, 9) || !checkEach(cpf, 10) {
		return false, errors.New("Invalid value")
	}

	return true, nil
}

func fixType(digits interface{}) ([]int, error) {
	switch v := digits.(type) {
	case string:
		v = strings.Replace(v, ".", "", -1)
		v = strings.Replace(v, "-", "", -1)
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

var blackList = []string{"1 2 3 4 5 6 7 8 9",
	"0 0 0 0", "1 1 1 1",
	"2 2 2 2", "3 3 3 3",
	"4 4 4 4", "5 5 5 5",
	"6 6 6 6", "7 7 7 7",
	"8 8 8 8", "9 9 9 9"}

func blackListed(digits []int) (res bool) {
	var cpf = fmt.Sprintf("%v", digits)
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
