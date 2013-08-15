package main

import (
	".." // testing
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: go run ./main.go <XXXXXXXXXXX>")
		return
	}
	println(cpf.Valid(os.Args[1]))
}
