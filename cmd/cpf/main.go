package main

import (
	"os"

	"github.com/mvrilo/go-cpf"
)

const help = `cpf is a tool to validate and generate CPF numbers.

Usage:
	cpf [xxxxxxxxxxx] [-vh]

To generate just run it without any arguments.

Links:
	https://en.wikipedia.org/wiki/Cadastro_de_Pessoas_F%C3%ADsicas
	https://github.com/mvrilo/go-cpf
`

var verbose = false

func main() {
	if len(os.Args) < 2 {
		generate()
		return
	}

	for _, arg := range os.Args {
		if arg == "-v" {
			verbose = true
		}

		if arg == "-h" {
			println(help)
			return
		}
	}

	validate()
}

func validate() {
	_, err := cpf.Valid(os.Args[1])

	if verbose && err != nil {
		os.Stdout.Write([]byte(err.Error() + " "))
		os.Stderr.Write([]byte("\u2718\n"))
		os.Exit(1)
		return
	}

	os.Stderr.Write([]byte("\u2714\n"))
	os.Exit(0)
}

func generate() {
	println(cpf.GeneratePretty())
	os.Exit(0)
}
