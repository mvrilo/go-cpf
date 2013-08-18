# go-cpf

[CPF] validator and generator written in Go.

## Usage

#### Command-line example:
```bash
$ go-cpf
487.160.523-07
$ go-cpf 487.160.523-07
✔
```

To generate, just run it without any arguments.

#### API example:
```go
package main

import "github.com/mvrilo/go-cpf"

func main() {
        println(cpf.GeneratePretty())
}
```

## LICENSE

See [LICENSE](https://github.com/mvrilo/go-cpf/blob/master/LICENSE)

[CPF]: https://en.wikipedia.org/wiki/Cadastro_de_Pessoas_F%C3%ADsicas
