package main

import (
	"os"
	"strings"

	"github.com/mvrilo/cpfer/lib"
)

const help = `cpfer is a tool to validate and generate CPF numbers.

Usage:
	cpfer [xxxxxxxxxxx]

To generate just run it without any arguments.

Links:
	https://en.wikipedia.org/wiki/Cadastro_de_Pessoas_F%C3%ADsicas
	https://github.com/mvrilo/cpfer
`

func main() {
	if len(os.Args) < 2 {
		generate()
		return
	}

	if strings.Contains(os.Args[1], "-h") {
		println(help)
		return
	}

	validate()
}

func validate() {
	var (
		res = "\u2718"
		ok  = cpf.Valid(os.Args[1])
	)

	if len(os.Args) == 3 {
		println(ok)
		return
	}

	if ok {
		res = "\u2714"
	}
	println(res)
}

func generate() {
	println(cpf.GeneratePretty())
}
